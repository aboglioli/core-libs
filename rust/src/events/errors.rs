use thiserror::Error;

#[derive(Error, Debug)]
pub enum Error {
    #[error("invalid event")]
    InvalidEvent,
    #[error("could not serialize event: {0}")]
    SerializingEvent(#[source] serde_json::Error),
    #[error("could not deserialize event: {0}")]
    DeserializingEvent(#[source] serde_json::Error),
    #[error("could not serialize event payload: {0}")]
    SerializingPayload(#[source] serde_json::Error),
    #[error("could not deserialize event payload: {0}")]
    DeserializingPayload(#[source] serde_json::Error),
    #[error("could not publish event: {0}")]
    PublishingEvent(#[source] Box<dyn std::error::Error + Sync + Send>),
    #[error("could not subscribe to {subject} subject: {err}")]
    SubscribingToSubject {
        subject: String,
        #[source]
        err: Box<dyn std::error::Error + Sync + Send>,
    },
}
