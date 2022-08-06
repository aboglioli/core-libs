use std::error::Error as StdError;
use std::fmt::{Display, Formatter, Result as FmtResult};
use std::result::Result as StdResult;

use crate::errors::Metadata;

// Result
pub type Result<T> = StdResult<T, Error>;

// ErrorCode
pub trait Define {
    fn define(&self) -> &str;
}

// Error
#[derive(Debug, Clone, PartialEq)]
pub struct Error {
    code: String,
    message: String,
    cause: Option<Box<Error>>,
    metadata: Metadata,
}

impl Error {
    pub fn new<C, S, M>(code: C, message: S, metadata: M) -> Error
    where
        C: Define,
        S: Into<String>,
        M: Into<Option<Metadata>>,
    {
        let metadata = metadata.into().unwrap_or_else(Metadata::new);

        Error {
            code: code.define().to_owned(),
            message: message.into(),
            cause: None,
            metadata,
        }
    }

    pub fn wrap<C, S, M>(code: C, err: Error, message: S, metadata: M) -> Error
    where
        C: Define,
        S: Into<String>,
        M: Into<Option<Metadata>>,
    {
        let mut error = Error::new(code, message, metadata);
        error.cause = Some(Box::new(err));
        error
    }

    pub fn wrap_raw<C, E, S, M>(code: C, err: &E, message: S, metadata: M) -> Error
    where
        C: Define,
        E: StdError + ?Sized,
        S: Into<String>,
        M: Into<Option<Metadata>>,
    {
        Error::new(
            code,
            format!("{} ({})", message.into(), err.to_string()),
            metadata,
        )
    }

    pub fn code(&self) -> &str {
        &self.code
    }

    pub fn message(&self) -> &str {
        &self.message
    }

    pub fn cause(&self) -> Option<&Box<Error>> {
        self.cause.as_ref()
    }

    pub fn metadata(&self) -> &Metadata {
        &self.metadata
    }
}

impl Display for Error {
    fn fmt(&self, f: &mut Formatter<'_>) -> FmtResult {
        write!(f, "{}: {}", self.code, self.message)
    }
}

impl StdError for Error {}

#[cfg(test)]
mod tests {
    use super::*;

    use serde_json::Value;
    use std::io::{Error as IoError, ErrorKind as IoErrorKind};

    #[derive(PartialEq)]
    enum CustomError {
        Internal,
        Validation,
        Application,
    }

    impl Define for CustomError {
        fn define(&self) -> &str {
            match self {
                CustomError::Internal => "internal",
                CustomError::Validation => "validation",
                CustomError::Application => "application",
            }
        }
    }

    #[test]
    fn create_new_error() {
        let err = Error::new(CustomError::Internal, "internal error", None);
        assert_eq!(err.code().to_string(), "internal");
        assert_eq!(err.message(), "internal error");
        assert!(err.cause().is_none());
        assert_eq!(err.metadata().values().len(), 0);
        assert_eq!(err.to_string(), "internal: internal error");
    }

    #[test]
    fn with_metadata() {
        let err = Error::new(
            CustomError::Validation,
            "some error",
            Metadata::with("key1", "value1")
                .and("key2", "value2")
                .and("key3", "value3"),
        );
        assert_eq!(err.metadata().values().len(), 3);

        assert_eq!(
            err.metadata().values()["key1"],
            Value::String("value1".to_string())
        );
        assert_eq!(
            err.metadata().values()["key2"],
            Value::String("value2".to_string())
        );
        assert_eq!(
            err.metadata().values()["key3"],
            Value::String("value3".to_string())
        );
    }

    #[test]
    fn wrap_error() {
        let internal_err = Error::new(
            CustomError::Internal,
            "internal",
            Metadata::with("internal", true),
        );
        let application_err = Error::wrap(
            CustomError::Application,
            internal_err.clone(),
            "application",
            Metadata::with("application", true),
        );

        assert!(application_err.cause().is_some());
        assert_eq!(application_err.cause(), Some(&Box::new(internal_err)));
    }

    #[test]
    fn wrap_raw_error() {
        let raw_err = IoError::new(IoErrorKind::NotFound, "raw_err");
        let err = Error::wrap_raw(CustomError::Internal, &raw_err, "internal error", None);

        assert!(err.cause().is_none());
        assert_eq!(err.message(), "internal error (raw_err)");
    }
}
