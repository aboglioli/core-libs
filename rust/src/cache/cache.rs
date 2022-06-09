use async_trait::async_trait;

use crate::errors::{Define, Result};

pub enum CacheError {
    Internal,
}

impl Define for CacheError {
    fn define(&self) -> &str {
        match self {
            CacheError::Internal => "cache.internal",
        }
    }
}

#[async_trait]
pub trait Cache<K, V> {
    async fn get(&self, k: &K) -> Result<Option<V>>;
    async fn set(&self, k: K, v: V) -> Result<()>;
    async fn delete(&self, k: &K) -> Result<()>;
}
