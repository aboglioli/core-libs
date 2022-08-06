use thiserror::Error;

#[derive(Error, Debug)]
pub enum Error {
    #[error("invalid id: {0}")]
    InvalidId(String),
    #[error("invalid timestamps")]
    InvalidTimestamps,
    #[error("invalid version: {0}")]
    InvalidVersion(i64),
}
