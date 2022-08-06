package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

var (
	ErrInvalidId = errors.New("invalid id")
)

type StrId struct {
	id string
}

func NewStrId(id string) (StrId, error) {
	if len(id) == 0 {
		return StrId{}, ErrInvalidId
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
