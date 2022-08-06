use async_nats::Client;
use async_trait::async_trait;
use chrono::{DateTime, Utc};
use futures::StreamExt;
use serde::{Deserialize, Serialize};

use crate::events::{Error, Event, Handler, Publisher, Subscriber};

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
    async fn publish(&self, events: &[Event]) -> Result<(), Error> {
        for event in events {
            let nats_event = NatsEvent {
                id: event.id().to_string(),
                entity_id: event.entity_id().to_string(),
                topic: event.topic().to_string(),
                payload: event.payload().to_vec(),
                timestamp: event.timestamp().clone(),
            };

            let msg = serde_json::to_vec(&nats_event).map_err(Error::SerializingEvent)?;

            self.client
                .publish(event.topic().to_string(), msg.into())
                .await
                .map_err(|err| Error::PublishingEvent(Box::new(err)))?;
        }

        Ok(())
    }
}

#[async_trait]
impl Subscriber for NatsEventBus {
    async fn subscribe(&self, subject: &str, handler: Box<dyn Handler>) -> Result<(), Error> {
        let mut sub = self
            .client
            .queue_subscribe(subject.to_string(), self.consumer_group.to_string())
            .await
            .map_err(|err| Error::SubscribingToSubject {
                subject: subject.to_string(),
                err,
            })?;

        tokio::spawn(async move {
            while let Some(msg) = sub.next().await {
                let nats_event: NatsEvent = serde_json::from_slice(&msg.payload)
                    .map_err(Error::DeserializingEvent)
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
