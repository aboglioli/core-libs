package errors

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aboglioli/libs/go/collections"
)

// ErrorCode
type ErrorCode struct {
	code string
}

func Define(code string) *ErrorCode {
	if len(code) == 0 {
		panic("empty error code")
	}

	return &ErrorCode{code}
}

func (c ErrorCode) String() string {
	return c.code
}

// Error
type Error struct {
	code     *ErrorCode
	message  string
	cause    error
	metadata collections.Metadata
}

func New(code *ErrorCode, message string, metadata ...collections.Metadata) *Error {
	m := collections.NewMetadata()
	for _, metadata := range metadata {
		m = m.Merge(metadata)
	}

	return &Error{
		code:     code,
		message:  message,
		metadata: m,
	}
}

func Wrap(code *ErrorCode, cause error, message string, metadata ...collections.Metadata) *Error {
	err := New(code, message, metadata...)
	err.cause = cause

	return err
}

func (err *Error) Code() *ErrorCode {
	return err.code
}

func (err *Error) Message() string {
	return err.message
}

func (err *Error) Cause() error {
	return err.cause
}

func (err *Error) Metadata() collections.Metadata {
	return err.metadata
}

func (err *Error) Error() string {
	str := fmt.Sprintf("%s: %s", err.code.code, err.message)

	if err.cause != nil {
		str += fmt.Sprintf(" (%s)", err.cause.Error())
	}

	if len(err.metadata.Values()) > 0 {
		metadataStr := make([]string, 0, len(err.metadata.Values()))
		for k, v := range err.metadata.Values() {
			metadataStr = append(metadataStr, fmt.Sprintf("[%s = %v]", k, v))
		}

		str += fmt.Sprintf(" %s", strings.Join(metadataStr, ", "))
	}

	return str
}

func (err *Error) MarshalJSON() ([]byte, error) {
	var cause json.Marshaler
	if err.cause != nil {
		switch err := err.cause.(type) {
		case *Error:
			cause = err
		case error:
			cause = &Error{
				message: err.Error(),
			}
		}
	}

	return json.Marshal(map[string]interface{}{
		"code":     err.code.code,
		"message":  err.message,
		"cause":    cause,
		"metadata": err.metadata,
	})
}
