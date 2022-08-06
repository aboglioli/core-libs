use thiserror::Error;

#[derive(Error, Debug)]
pub enum VersionError {
    #[error("invalid version")]
    Invalid,
}

#[derive(Debug, Clone, PartialEq)]
pub struct Version {
    version: i64,
    updated: bool,
}

impl Version {
    pub fn new(version: i64) -> Result<Version, VersionError> {
        if version < 1 {
            return Err(VersionError::Invalid);
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
