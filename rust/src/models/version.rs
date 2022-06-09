use crate::collections::Metadata;
use crate::errors::{Define, Error, Result};

pub enum VersionError {
    Invalid,
}

impl Define for VersionError {
    fn define(&self) -> &str {
        match self {
            VersionError::Invalid => "version.invalid",
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct Version {
    version: i64,
    updated: bool,
}

impl Version {
    pub fn new(version: i64) -> Result<Version> {
        if version < 1 {
            return Err(Error::new(
                VersionError::Invalid,
                "version is smaller than 1",
                Metadata::with("version", version),
            ));
        }

        Ok(Version {
            version,
            updated: false,
        })
    }

    pub fn init_version() -> Version {
        Version {
            version: 1,
            updated: true,
        }
    }

    pub fn value(&self) -> i64 {
        self.version
    }

    pub fn incr(&self) -> Version {
        if self.updated {
            return self.clone();
        }

        Version {
            version: self.version + 1,
            updated: true,
        }
    }
}
