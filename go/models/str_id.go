package models

import (
	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"github.com/aboglioli/core-libs/go/collections"
	"github.com/aboglioli/core-libs/go/errors"
)

var (
	ErrInvalidId = errors.Define("id.invalid")
)

type StrId struct {
	id string
}

func NewStrId(id string) (StrId, error) {
	if id == "" {
		return StrId{}, errors.New(
			ErrInvalidId,
			"empty string id",
			collections.WithMetadata("id", id),
		)
	}

	return StrId{
		id: id,
	}, nil
}

func GenerateUuid() StrId {
	id, _ := NewStrId(uuid.New().String())
	return id
}

func GenerateSlug(str string) (StrId, error) {
	s := slug.Make(str)
	return NewStrId(s)
}

func (id StrId) Value() string {
	return id.id
}

func (id StrId) Equals(other StrId) bool {
	return id.id == other.id
}
