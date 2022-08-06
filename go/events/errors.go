package events

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidEvent   = errors.New("invalid event")
	ErrUnknownPayload = errors.New("unknown event payload")
)

// ErrSerializingEvent
type ErrSerializingEvent struct {
	Event *Event
	err   error
}

func (e *ErrSerializingEvent) Error() string {
	return fmt.Sprintf("could not serialize event: %s", e.err)
}

func (e *ErrSerializingEvent) Unwrap() error {
	return e.err
}

// ErrDeserializingEvent
type ErrDeserializingEvent struct {
	err error
}

func (e *ErrDeserializingEvent) Error() string {
	return fmt.Sprintf("could not deserialize event: %s", e.err)
}

func (e *ErrDeserializingEvent) Unwrap() error {
	return e.err
}

// ErrSerializingPayload
type ErrSerializingPayload struct {
	Event *Event
	err   error
}

func (e *ErrSerializingPayload) Error() string {
	return fmt.Sprintf("could not serialize event payload: %s", e.err)
}

func (e *ErrSerializingPayload) Unwrap() error {
	return e.err
}

// ErrDeserializingPâyload
type ErrDeserializingPayload struct {
	Event *Event
	err   error
}

func (e *ErrDeserializingPayload) Error() string {
	return fmt.Sprintf("could not deserialize event payload: %s", e.err)
}

func (e *ErrDeserializingPayload) Unwrap() error {
	return e.err
}

// ErrPublishingEvent
type ErrPublishingEvent struct {
	Event *Event
	err   error
}

func (e *ErrPublishingEvent) Error() string {
	return fmt.Sprintf("could not publish event: %s", e.err)
}

func (e *ErrPublishingEvent) Unwrap() error {
	return e.err
}

// ErrSubscribingToEvent
type ErrSubscribingToSubject struct {
	Subject string
	err     error
}

func (e *ErrSubscribingToSubject) Error() string {
	return fmt.Sprintf("could not subscribe to %s subject: %s", e.Subject, e.err)
}

func (e *ErrSubscribingToSubject) Unwrap() error {
	return e.err
}
