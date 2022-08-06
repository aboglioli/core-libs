use serde::Serialize;
use serde_json::Value;
use std::collections::HashMap;

#[derive(Debug, Clone, PartialEq)]
pub struct Metadata {
    metadata: HashMap<String, Value>,
}

impl Metadata {
    pub fn new() -> Self {
        Metadata {
            metadata: HashMap::new(),
        }
    }

    pub fn with<K, V>(key: K, value: V) -> Self
    where
        K: Into<String>,
        V: Serialize,
    {
        Metadata::new().and(key, value)
    }
    pub fn and<K, V>(mut self, key: K, value: V) -> Self
    where
        K: Into<String>,
        V: Serialize,
    {
        if let Ok(value) = serde_json::to_value(value) {
            self.metadata.insert(key.into(), value);
        }

        self
    }

    pub fn merge(mut self, other: Metadata) -> Self {
        self.metadata.extend(other.metadata);
        self
    }

    pub fn values(&self) -> &HashMap<String, Value> {
        &self.metadata
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    use serde_json::json;

    #[derive(Serialize)]
    struct Data {
        msg: String,
    }

    #[test]
    fn metadata() {
        let m = Metadata::with("prop1", "hello world")
            .and("prop2", 123)
            .and(
                "prop3",
                Data {
                    msg: "Hello World".to_string(),
                },
            );

        assert_eq!(
            m.values().get("prop1").unwrap(),
            &Value::String("hello world".to_string())
        );
        assert_eq!(m.values().get("prop2").unwrap(), &json!(123));
        assert_eq!(
            m.values().get("prop3").unwrap(),
            &json!(Data {
                msg: "Hello World".to_string()
            })
        );
    }
}
