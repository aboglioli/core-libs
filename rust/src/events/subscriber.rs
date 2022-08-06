use async_trait::async_trait;

use crate::events::{Event, EventError};

#[async_trait]
pub trait Handler: Sync + Send {
    async fn handle(&self, event: &Event) -> Result<(), EventError>;
}

#[async_trait]
pub trait Subscriber {
    async fn subscribe(&self, subject: &str, handler: Box<dyn Handler>) -> Result<(), EventError>;
}
