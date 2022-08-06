use async_trait::async_trait;
use serde::Serialize;
use std::cmp::Eq;
use std::collections::HashMap;
use std::hash::Hash;
use std::sync::Arc;
use tokio::sync::RwLock;

use crate::cache::{Cache, Error};

#[derive(Clone)]
pub struct InMemCache<K, V> {
    items: Arc<RwLock<HashMap<K, V>>>,
}

impl<K, V> InMemCache<K, V>
where
    K: Clone,
    V: Clone,
{
    pub fn new() -> InMemCache<K, V> {
        InMemCache {
            items: Arc::new(RwLock::new(HashMap::new())),
        }
    }

    pub async fn all(&self) -> HashMap<K, V> {
        let items = self.items.read().await;
        items.clone()
    }
}

#[async_trait]
impl<K, V> Cache<K, V> for InMemCache<K, V>
where
    K: Eq + Hash + Sync + Send + Serialize,
    V: Clone + Sync + Send,
{
    async fn get(&self, k: &K) -> Result<Option<V>, Error> {
        let items = self.items.read().await;

        Ok(items.get(k).cloned())
    }

    async fn set(&self, k: K, v: V) -> Result<(), Error> {
        let mut items = self.items.write().await;

        items.insert(k, v);

        Ok(())
    }

    async fn delete(&self, k: &K) -> Result<(), Error> {
        let mut items = self.items.write().await;

        items.remove(k);

        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn set_get_and_delete() {
        let cache = InMemCache::new();

        let key1 = "key_1".to_string();
        let value1 = 64;

        let res = cache.get(&key1).await;
        assert!(res.is_ok());
        assert!(res.unwrap().is_none());

        let res = cache.set(key1.clone(), value1).await;
        assert!(!res.is_err());

        let res = cache.get(&key1).await;
        assert!(!res.is_err());
        assert_eq!(res.unwrap().unwrap(), value1);

        let res = cache.delete(&key1).await;
        assert!(!res.is_err());

        let res = cache.get(&key1).await;
        assert!(res.unwrap().is_none());
    }

    #[tokio::test]
    async fn thread_safe() {
        let cache = InMemCache::new();

        let cache1 = cache.clone();
        let t1 = tokio::spawn(async move {
            cache1.set(1, 128).await.unwrap();
        });

        let cache2 = cache.clone();
        let t2 = tokio::spawn(async move {
            cache2.set(2, 256).await.unwrap();
        });

        let (_, _) = tokio::join!(t1, t2);

        let res = cache.get(&1).await;
        assert!(!res.is_err());
        assert_eq!(res.unwrap().unwrap(), 128);

        let res = cache.get(&2).await;
        assert!(!res.is_err());
        assert_eq!(res.unwrap().unwrap(), 256);
    }
}
