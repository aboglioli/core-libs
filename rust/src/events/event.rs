use async_trait::async_trait;
use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use std::borrow::Cow;
use uuid::Uuid;

use crate::collections::Metadata;
use crate::errors::{Define, Error, Result};

pub enum EventError {
    Invalid,
    PayloadSerialization,
    Internal,
}

impl Define for EventError {
    fn define(&self) -> &str {
        match self {
            EventError::Invalid => "event.invalid",
            EventError::PayloadSerialization => "event.payload_serialization",
            EventError::Internal => "event.internal",
        }
    }
}

// Publisher and subscriber
#[async_trait]
pub trait Publisher {
    async fn publish(&self, events: &[Event]) -> Result<()>;
}

#[async_trait]
pub trait Handler: Sync + Send {
    async fn handle(&self, event: &Event) -> Result<()>;
}

#[async_trait]
pub trait Subscriber {
    async fn subscribe(&self, subject: Cow<'_, str>, handler: Box<dyn Handler>) -> Result<()>;
}

// Event
#[derive(Debug, Clone, PartialEq)]
pub struct Event {
    id: String,
    entity_id: String,
    topic: String,
    payload: Vec<u8>,
    timestamp: DateTime<Utc>,
}

impl Event {
    pub fn new(
        id: String,
        entity_id: String,
        topic: String,
        payload: Vec<u8>,
        timestamp: DateTime<Utc>,
    ) -> Result<Event> {
        let metadata = Metadata::with("id", &id)
            .and("entity_id", &entity_id)
            .and("topic", &topic)
            .and("payload", &payload);

        if id.is_empty() {
            return Err(Error::new(
                EventError::Invalid,
                "event id is empty",
                metadata,
            ));
        }

        if entity_id.is_empty() {
            return Err(Error::new(
                EventError::Invalid,
                "event entity_id is empty",
                metadata,
            ));
        }

        if topic.is_empty() {
            return Err(Error::new(
                EventError::Invalid,
                "event topic is empty",
                metadata,
            ));
        }

        if payload.is_empty() {
            return Err(Error::new(
                EventError::Invalid,
                "cannot marshal event payload",
                metadata,
            ));
        }

        Ok(Event {
            id,
            entity_id,
            topic,
            payload,
            timestamp,
        })
    }

    pub fn create<I, T, P>(entity_id: I, topic: T, payload: &P) -> Result<Event>
    where
        I: Into<String>,
        T: Into<String>,
        P: Serialize,
    {
        let entity_id = entity_id.into();
        let topic = topic.into();

        let payload = serde_json::to_vec(payload).map_err(|err| {
            Error::wrap_raw(
                EventError::PayloadSerialization,
                &err,
                "could not serialize event payload",
                Metadata::with("entity_id", &entity_id)
                    .and("topic", &topic)
                    .and("payload", payload),
            )
        })?;

        Event::new(
            Uuid::new_v4().to_string(),
            entity_id,
            topic,
            payload,
            Utc::now(),
        )
    }

    pub fn id(&self) -> &str {
        &self.id
    }

    pub fn entity_id(&self) -> &str {
        &self.entity_id
    }

    pub fn topic(&self) -> &str {
        &self.topic
    }

    pub fn payload(&self) -> &[u8] {
        &self.payload
    }

    pub fn deserialize_payload<'a, T>(&'a self) -> Result<T>
    where
        T: Deserialize<'a>,
    {
        serde_json::from_slice(&self.payload).map_err(|err| {
            Error::wrap_raw(
                EventError::PayloadSerialization,
                &err,
                "could not deserialize event payload",
                Metadata::with("id", &self.id)
                    .and("entity_id", &self.entity_id)
                    .and("topic", &self.topic)
                    .and("payload", &self.payload),
            )
        })
    }

    pub fn timestamp(&self) -> &DateTime<Utc> {
        &self.timestamp
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[derive(Debug, Serialize, Deserialize)]
    struct Data {
        msg: String,
    }

    #[test]
    fn create() {
        let res = Event::create(
            "entity#01",
            "topic.code",
            &Data {
                msg: "Hello World".to_string(),
            },
        );
        assert!(res.is_ok());

        let event = res.unwrap();
        assert_eq!(event.entity_id(), "entity#01");
        assert_eq!(event.topic(), "topic.code");
        assert!(event.payload().len() > 0);
    }

    #[test]
    fn payload_serialization_and_deserialization() {
        let event = Event::create(
            "entity#01",
            "topic.code",
            &Data {
                msg: "Hello World".to_string(),
            },
        )
        .unwrap();

        let res = event.deserialize_payload();
        assert!(res.is_ok());

        let data: Data = res.unwrap();
        assert_eq!(data.msg, "Hello World");
    }
}
