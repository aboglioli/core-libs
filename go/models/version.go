package models

import (
	"github.com/aboglioli/libs/go/collections"
	"github.com/aboglioli/libs/go/errors"
)

var (
	ErrInvalidVersion = errors.Define("version.invalid")
)

type Version struct {
	version int64

	updated bool
}

func NewVersion(version int64) (Version, error) {
	if version < 1 {
		return Version{}, errors.New(
			ErrInvalidVersion,
			"version is smaller than 1",
			collections.WithMetadata("version", version),
		)
	}

	return Version{
		version: version,
		updated: false,
	}, nil
}

func InitVersion() Version {
	return Version{
		version: 1,
		updated: true,
	}
}

func (v Version) Value() int64 {
	return v.version
}

func (v Version) Equals(other Version) bool {
	return v.version == other.version
}

func (v Version) Incr() Version {
	if v.updated {
		return v
	}

	return Version{
		version: v.version + 1,
		updated: true,
	}
}
