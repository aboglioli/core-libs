use chrono::{DateTime, Utc};

use crate::collections::Metadata;
use crate::errors::{Define, Error, Result};

pub enum TimestampsError {
    Invalid,
}

impl Define for TimestampsError {
    fn define(&self) -> &str {
        match self {
            TimestampsError::Invalid => "timestamps.invalid",
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct Timestamps {
    created_at: DateTime<Utc>,
    updated_at: DateTime<Utc>,
    deleted_at: Option<DateTime<Utc>>,
}

impl Timestamps {
    pub fn new(
        created_at: DateTime<Utc>,
        updated_at: DateTime<Utc>,
        deleted_at: Option<DateTime<Utc>>,
    ) -> Result<Timestamps> {
        let metadata = Metadata::with("created_at", created_at)
            .and("updated_at", updated_at)
            .and("deleted_at", deleted_at);

        if updated_at < created_at {
            return Err(Error::new(
                TimestampsError::Invalid,
                "update date is before create date",
                metadata,
            ));
        }

        if let Some(deleted_at) = deleted_at {
            if deleted_at < created_at {
                return Err(Error::new(
                    TimestampsError::Invalid,
                    "delete date is before create date",
                    metadata,
                ));
            }

            if deleted_at < updated_at {
                return Err(Error::new(
                    TimestampsError::Invalid,
                    "delete date is before update date",
                    metadata,
                ));
            }
        }

        Ok(Timestamps {
            created_at,
            updated_at,
            deleted_at,
        })
    }

    pub fn create() -> Timestamps {
        let now = Utc::now();

        Timestamps {
            created_at: now,
            updated_at: now,
            deleted_at: None,
        }
    }

    pub fn created_at(&self) -> &DateTime<Utc> {
        &self.created_at
    }

    pub fn updated_at(&self) -> &DateTime<Utc> {
        &self.updated_at
    }

    pub fn update(&self) -> Timestamps {
        let mut timestamps = self.clone();
        timestamps.updated_at = Utc::now();
        timestamps
    }

    pub fn deleted_at(&self) -> Option<&DateTime<Utc>> {
        self.deleted_at.as_ref()
    }

    pub fn delete(&self) -> Timestamps {
        let mut timestamps = self.clone();
        timestamps.deleted_at = Some(Utc::now());
        timestamps
    }
}
