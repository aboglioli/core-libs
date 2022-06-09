use async_nats::Client;
use async_trait::async_trait;
use chrono::{DateTime, Utc};
use futures::StreamExt;
use serde::{Deserialize, Serialize};
use std::borrow::Cow;

use crate::collections::Metadata;
use crate::errors::{Error, Result};
use crate::events::{Event, EventError, Handler, Publisher, Subscriber};

#[derive(Serialize, Deserialize)]
struct NatsEvent {
    id: String,
    entity_id: String,
    topic: String,
    payload: Vec<u8>,
    timestamp: DateTime<Utc>,
}

pub struct NatsEventBus {
    consumer_group: String,
    client: Client,
}

impl NatsEventBus {
    pub fn new(consumer_group: String, client: Client) -> NatsEventBus {
        NatsEventBus {
            consumer_group,
            client,
        }
    }
}

#[async_trait]
impl Publisher for NatsEventBus {
    async fn publish(&self, events: &[Event]) -> Result<()> {
        for event in events {
            let nats_event = NatsEvent {
                id: event.id().to_string(),
                entity_id: event.entity_id().to_string(),
                topic: event.topic().to_string(),
                payload: event.payload().to_vec(),
                timestamp: event.timestamp().clone(),
            };

            let msg = serde_json::to_vec(&nats_event).map_err(|err| {
                Error::wrap_raw(
                    EventError::Internal,
                    &err,
                    "could not marshal message",
                    Metadata::with("message", &nats_event),
                )
            })?;

            self.client
                .publish(event.topic().to_string(), msg.into())
                .await
                .map_err(|err| {
                    Error::wrap_raw(
                        EventError::Internal,
                        &*err,
                        "could not publish message",
                        Metadata::with("message", nats_event),
                    )
                })?;
        }

        Ok(())
    }
}

#[async_trait]
impl Subscriber for NatsEventBus {
    async fn subscribe(&self, subject: Cow<'_, str>, handler: Box<dyn Handler>) -> Result<()> {
        let subject = subject.into_owned();

        let mut sub = self
            .client
            .queue_subscribe(subject.to_string(), self.consumer_group.to_string())
            .await
            .map_err(|err| {
                Error::wrap_raw(
                    EventError::Internal,
                    &err,
                    "could not subscriber to subject",
                    Metadata::with("subject", subject),
                )
            })?;

        tokio::spawn(async move {
            while let Some(msg) = sub.next().await {
                let nats_event: NatsEvent = serde_json::from_slice(&msg.payload)
                    .map_err(|err| {
                        Error::wrap_raw(
                            EventError::Internal,
                            &err,
                            "could not unmarshal message",
                            Metadata::with("message_payload", msg.payload.as_ref()),
                        )
                    })
                    .unwrap();

                let event = Event::new(
                    nats_event.id,
                    nats_event.entity_id,
                    nats_event.topic,
                    nats_event.payload,
                    nats_event.timestamp,
                )
                .unwrap();

                handler.handle(&event).await.unwrap();
            }
        });

        Ok(())
    }
}
