use async_trait::async_trait;
use thiserror::Error;

#[derive(Error, Debug)]
pub enum CacheError {
    #[error("internal cache error")]
    Internal,
}

#[async_trait]
pub trait Cache<K, V> {
    async fn get(&self, k: &K) -> Result<Option<V>, CacheError>;
    async fn set(&self, k: K, v: V) -> Result<(), CacheError>;
    async fn delete(&self, k: &K) -> Result<(), CacheError>;
}
