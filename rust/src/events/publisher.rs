use async_trait::async_trait;

use crate::events::{Event, EventError};

#[async_trait]
pub trait Publisher {
    async fn publish(&self, events: &[Event]) -> Result<(), EventError>;
}
