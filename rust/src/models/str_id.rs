use uuid::Uuid;

use crate::collections::Metadata;
use crate::errors::{Define, Error, Result};

pub enum IdError {
    Invalid,
}

impl Define for IdError {
    fn define(&self) -> &str {
        match self {
            IdError::Invalid => "id.invalid",
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct StrId {
    id: String,
}

impl StrId {
    pub fn new<S: Into<String>>(id: S) -> Result<StrId> {
        let id = id.into();

        if id.is_empty() {
            return Err(Error::new(
                IdError::Invalid,
                "empty string id",
                Metadata::with("id", id),
            ));
        }

        Ok(StrId { id })
    }

    pub fn generate_uuid() -> Result<StrId> {
        let uuid = Uuid::new_v4();
        StrId::new(uuid.to_string())
    }

    pub fn generate_slug<S: Into<String>>(str: S) -> Result<StrId> {
        StrId::new(slug::slugify(str.into()))
    }

    pub fn value(&self) -> &str {
        &self.id
    }
}

impl ToString for StrId {
    fn to_string(&self) -> String {
        self.id.to_owned()
    }
}
