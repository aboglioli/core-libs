package errors

import (
	"errors"
	"testing"

	"github.com/aboglioli/libs/go/collections"
	"github.com/stretchr/testify/assert"
)

func TestCreateError(t *testing.T) {
	type test struct {
		name     string
		code     *ErrorCode
		message  string
		cause    error
		metadata map[string]interface{}
	}

	tests := []test{
		{
			name:     "error with message only",
			code:     Define("basic_error"),
			message:  "custom message",
			metadata: make(map[string]interface{}),
		},
		{
			name:    "error with message and metadata",
			code:    Define("error_with_metadata"),
			message: "custom message",
			metadata: map[string]interface{}{
				"str": "value",
				"num": 123,
			},
		},
		{
			name:     "wrap error with cause",
			code:     Define("error_with_metadata"),
			message:  "custom message",
			cause:    errors.New("raw error"),
			metadata: make(map[string]interface{}),
		},
		{
			name:    "error with message, metadata and cause",
			code:    Define("error_with_metadata"),
			message: "custom message",
			cause:   errors.New("raw error"),
			metadata: map[string]interface{}{
				"str": "value",
				"num": 123,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := collections.NewMetadata()
			for k, v := range test.metadata {
				m = m.And(k, v)
			}

			var err *Error
			if test.cause != nil {
				err = Wrap(test.code, test.cause, test.message, m)
			} else {
				err = New(test.code, test.message, m)
			}

			assert.Equal(t, test.code, err.code)
			assert.Equal(t, test.message, err.message)
			assert.Equal(t, test.cause, err.cause)
			assert.Equal(t, test.metadata, err.metadata.Values())
		})
	}
}
