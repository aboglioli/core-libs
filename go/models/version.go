package models

import (
	"errors"
)

var (
	ErrInvalidVersion = errors.New("invalid version")
)

type Version struct {
	version int64

	updated bool
}

func NewVersion(version int64) (Version, error) {
	if version < 1 {
		return Version{}, ErrInvalidVersion
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
