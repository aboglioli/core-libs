use lazy_static::lazy_static;
use regex::Regex;

use crate::collections::Metadata;
use crate::errors::{Define, Error, Result};

lazy_static! {
    static ref RE_PART: Regex = Regex::new("^[a-z]+[a-z0-9_]*[a-z0-9]+$").unwrap();
}

pub enum PathError {
    Invalid,
}

impl Define for PathError {
    fn define(&self) -> &str {
        match self {
            PathError::Invalid => "path.invalid",
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct Path {
    parts: Vec<String>,
    separator: String,
    wildcards: Vec<String>,
}

fn is_part_valid(part: &str, separator: &str, wildcards: &[String]) -> bool {
    if part == separator {
        return true;
    }

    if wildcards.iter().any(|wildcard| part == wildcard) {
        return true;
    }

    RE_PART.is_match(part)
}

impl Path {
    pub fn new<P, S, W>(path: P, separator: S, wildcards: Vec<W>) -> Result<Path>
    where
        P: Into<String>,
        S: Into<String>,
        W: Into<String>,
    {
        let path = path.into();
        let separator = separator.into();
        let wildcards: Vec<String> = wildcards.into_iter().map(Into::into).collect();

        let parts: Vec<String> = path
            .split(&separator)
            .into_iter()
            .map(|part| part.to_lowercase())
            .collect();

        if parts.is_empty() {
            return Err(Error::new(PathError::Invalid, "full path is empty", None));
        }

        Ok(Path {
            parts: parts
                .into_iter()
                .map(|part| {
                    if !is_part_valid(&part, &separator, &wildcards) {
                        return Err(Error::new(
                            PathError::Invalid,
                            "path part has invalid characters",
                            Metadata::with("part", part),
                        ));
                    }

                    Ok(part)
                })
                .collect::<Result<Vec<String>>>()?,
            separator,
            wildcards,
        })
    }

    pub fn parts(&self) -> &[String] {
        &self.parts
    }

    pub fn separator(&self) -> &str {
        &self.separator
    }

    pub fn wildcards(&self) -> &[String] {
        &self.wildcards
    }
}

impl ToString for Path {
    fn to_string(&self) -> String {
        self.parts.join(&self.separator)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn new_path() {
        // Simple
        let path = Path::new::<_, _, &str>("my.path", ".", Vec::new()).unwrap();
        assert_eq!(path.parts(), &["my", "path"]);
        assert_eq!(path.separator(), ".");
        assert_eq!(path.wildcards(), &[] as &[String; 0]);
        assert_eq!(path.to_string(), "my.path");

        // With wildcards
        let path = Path::new("my.path", ".", vec!["*", ">"]).unwrap();
        assert_eq!(path.parts(), &["my", "path"]);
        assert_eq!(path.separator(), ".");
        assert_eq!(path.wildcards(), &["*", ">"]);
        assert_eq!(path.to_string(), "my.path");

        // Single partt
        let path = Path::new("my_path", "#", vec!["*", ">"]).unwrap();
        assert_eq!(path.parts(), &["my_path"]);
        assert_eq!(path.separator(), "#");
        assert_eq!(path.wildcards(), &["*", ">"]);
        assert_eq!(path.to_string(), "my_path");

        // Invalid character
        assert!(Path::new::<_, _, &str>("my.p@th", ".", Vec::new()).is_err())
    }

    #[test]
    fn equals() {
        let path1 = Path::new::<_, _, &str>("one.two.three", ".", Vec::new()).unwrap();
        let path2 = Path::new::<_, _, &str>("one.three.two", ".", Vec::new()).unwrap();
        let path3 = Path::new::<_, _, &str>("two.one.three", ".", Vec::new()).unwrap();

        assert!(path1 == path1);
        assert!(path1 != path2);
        assert!(path2 != path3);
    }
}
