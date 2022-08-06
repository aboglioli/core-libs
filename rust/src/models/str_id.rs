use uuid::Uuid;

use crate::models::Error;

#[derive(Debug, Clone, PartialEq)]
pub struct StrId {
    id: String,
}

impl StrId {
    pub fn new<S: Into<String>>(id: S) -> Result<StrId, Error> {
        let id = id.into();

        if id.is_empty() {
            return Err(Error::InvalidId(id));
        }

        Ok(StrId { id })
    }

    pub fn generate_uuid() -> Result<StrId, Error> {
        let uuid = Uuid::new_v4();
        StrId::new(uuid.to_string())
    }

    pub fn generate_slug<S: Into<String>>(str: S) -> Result<StrId, Error> {
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
