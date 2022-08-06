package models

import (
	"time"
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
	if updatedAt.Before(createdAt) {
		return Timestamps{}, ErrInvalidTimestamps
	}

	if deletedAt != nil {
		if deletedAt.Before(createdAt) {
			return Timestamps{}, ErrInvalidTimestamps
		}

		if deletedAt.Before(updatedAt) {
			return Timestamps{}, ErrInvalidTimestamps
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
