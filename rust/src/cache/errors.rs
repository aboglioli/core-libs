use thiserror::Error;

#[derive(Error, Debug)]
pub enum Error {
    #[error("internal cache error")]
    Internal,
}
