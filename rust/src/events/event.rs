use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use thiserror::Error;
use uuid::Uuid;

#[derive(Error, Debug)]
pub enum EventError {
    #[error("invalid event")]
    Invalid,
    #[error("payload serialization")]
    PayloadSerialization,
    #[error("internal event error")]
    Internal,
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
    ) -> Result<Event, EventError> {
        if id.is_empty() {
            return Err(EventError::Invalid);
        }

        if entity_id.is_empty() {
            return Err(EventError::Invalid);
        }

        if topic.is_empty() {
            return Err(EventError::Invalid);
        }

        if payload.is_empty() {
            return Err(EventError::Invalid);
        }

        Ok(Event {
            id,
            entity_id,
            topic,
            payload,
            timestamp,
        })
    }

    pub fn create<I, T, P>(entity_id: I, topic: T, payload: &P) -> Result<Event, EventError>
    where
        I: Into<String>,
        T: Into<String>,
        P: Serialize,
    {
        let entity_id = entity_id.into();
        let topic = topic.into();

        let payload = serde_json::to_vec(payload).map_err(|_| EventError::PayloadSerialization)?;

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

    pub fn deserialize_payload<'a, T>(&'a self) -> Result<T, EventError>
    where
        T: Deserialize<'a>,
    {
        serde_json::from_slice(&self.payload).map_err(|_| EventError::PayloadSerialization)
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
