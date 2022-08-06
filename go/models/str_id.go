package models

import (
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type StrId struct {
	id string
}

func NewStrId(id string) (StrId, error) {
	if len(id) == 0 {
		return StrId{}, &ErrInvalidId{id}
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
