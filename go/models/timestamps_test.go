package models

import (
	"testing"
	"time"

	"github.com/aboglioli/libs/go/types"
	"github.com/stretchr/testify/assert"
)

func TestTimestamps(t *testing.T) {
	t.Run("new timestamps", func(t *testing.T) {
		timestamps := CreateTimestamps()

		assert.Equal(t, timestamps.CreatedAt(), timestamps.UpdatedAt())
		assert.Nil(t, timestamps.DeletedAt())
	})

	t.Run("buid timestamps", func(t *testing.T) {
		type test struct {
			name        string
			createdAt   time.Time
			updatedAt   time.Time
			deletedAt   *time.Time
			expectedErr bool
		}

		tests := []test{
			{
				name:      "valid",
				createdAt: types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T15:30:00Z")),
				updatedAt: types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T15:30:00Z")),
			},
			{
				name:        "updated_at before created_at",
				createdAt:   types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T15:30:00Z")),
				updatedAt:   types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T15:29:00Z")),
				expectedErr: true,
			},
			{
				name:        "deleted_at before created_at",
				createdAt:   types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T15:30:00Z")),
				updatedAt:   types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T16:15:00Z")),
				deletedAt:   types.Ref(types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T14:45:00Z"))),
				expectedErr: true,
			},
			{
				name:        "deleted_at before updated_at",
				createdAt:   types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T15:30:00Z")),
				updatedAt:   types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T16:15:00Z")),
				deletedAt:   types.Ref(types.Unwrap(time.Parse(time.RFC3339, "2022-04-04T16:14:59Z"))),
				expectedErr: true,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				timestamps, err := NewTimestamps(test.createdAt, test.updatedAt, test.deletedAt)

				if test.expectedErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)

					assert.Equal(t, test.createdAt, timestamps.CreatedAt())
					assert.Equal(t, test.updatedAt, timestamps.UpdatedAt())
					assert.Equal(t, test.deletedAt, timestamps.DeletedAt())
				}
			})
		}
	})
}
