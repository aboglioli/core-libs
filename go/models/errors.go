package models

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidTimestamps = errors.New("invalid timestamps")
)

// ErrInvalidId
type ErrInvalidId struct {
	Id string
}

func (e *ErrInvalidId) Error() string {
	return fmt.Sprintf("invalid id: %s", e.Id)
}

// ErrInvalidVersion
type ErrInvalidVersion struct {
	Version int64
}

func (e *ErrInvalidVersion) Error() string {
	return fmt.Sprintf("invalid version: %d", e.Version)
}
