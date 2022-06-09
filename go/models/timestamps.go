package models

import (
	"time"

	"github.com/aboglioli/core-libs/go/collections"
	"github.com/aboglioli/core-libs/go/errors"
)

var (
	ErrInvalidTimestamps = errors.Define("timestamps.invalid")
)

type Timestamps struct {
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewTimestamps(
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (Timestamps, error) {
	m := collections.WithMetadata("created_at", createdAt).
		And("updated_at", updatedAt).
		And("deleted_at", deletedAt)

	if updatedAt.Before(createdAt) {
		return Timestamps{}, errors.New(
			ErrInvalidTimestamps,
			"update date is before create date",
			m,
		)
	}

	if deletedAt != nil {
		if deletedAt.Before(createdAt) {
			return Timestamps{}, errors.New(
				ErrInvalidTimestamps,
				"delete date is before create date",
				m,
			)
		}

		if deletedAt.Before(updatedAt) {
			return Timestamps{}, errors.New(
				ErrInvalidTimestamps,
				"delete date is before update date",
				m,
			)
		}
	}

	return Timestamps{
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}, nil
}

func CreateTimestamps() Timestamps {
	now := time.Now()

	return Timestamps{
		createdAt: now,
		updatedAt: now,
		deletedAt: nil,
	}
}

func (t Timestamps) CreatedAt() time.Time {
	return t.createdAt
}

func (t Timestamps) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t Timestamps) Update() Timestamps {
	t.updatedAt = time.Now()
	return t
}

func (t Timestamps) DeletedAt() *time.Time {
	return t.deletedAt
}

func (t Timestamps) Delete() Timestamps {
	now := time.Now()
	t.deletedAt = &now
	return t
}
