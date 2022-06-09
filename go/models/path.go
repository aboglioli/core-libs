package models

import (
	"regexp"
	"strings"

	"github.com/aboglioli/libs/go/collections"
	"github.com/aboglioli/libs/go/errors"
)

var rePart = regexp.MustCompile("^[a-z]+[a-z0-9_]*[a-z0-9]+$")

var (
	ErrInvalidPath = errors.Define("path.invalid")
)

type Path struct {
	parts     []string
	separator string
	wildcards []string
}

func NewPath(path string, separator string, wildcards ...string) (Path, error) {
	parts := strings.Split(path, separator)

	if len(parts) == 0 {
		return Path{}, errors.New(
			ErrInvalidPath,
			"full path is empty",
		)
	}

	loweredParts := make([]string, len(parts))
	for i, part := range parts {
		if len(part) == 0 {
			return Path{}, errors.New(
				ErrInvalidPath,
				"path part is empty",
				collections.WithMetadata("index", i),
				collections.WithMetadata("parts", parts),
			)
		}

		part = strings.ToLower(part)

		if !isPartValid(part, separator, wildcards...) {
			return Path{}, errors.New(
				ErrInvalidPath,
				"path part has invalid characters",
				collections.WithMetadata("path", path),
				collections.WithMetadata("parts", parts),
				collections.WithMetadata("index", i),
				collections.WithMetadata("separator", separator),
				collections.WithMetadata("wildcards", wildcards),
			)
		}

		loweredParts[i] = part
	}

	return Path{
		parts:     loweredParts,
		separator: separator,
		wildcards: wildcards,
	}, nil
}

func (p Path) Parts() []string {
	return p.parts
}

func (p Path) Separator() string {
	return p.separator
}

func (p Path) Wildcards() []string {
	return p.wildcards
}

func (p Path) String() string {
	return strings.Join(p.parts, p.separator)
}

func (p Path) Equals(other Path) bool {
	if len(p.parts) != len(other.parts) {
		return false
	}

	for i, p1 := range p.parts {
		p2 := other.parts[i]
		if p1 != p2 {
			return false
		}
	}

	return true
}

func isPartValid(part string, separator string, wildcards ...string) bool {
	if part == separator {
		return true
	}

	for _, wildcard := range wildcards {
		if part == wildcard {
			return true
		}
	}

	return rePart.MatchString(part)
}
