package models

import (
	"testing"

	"github.com/Fairuzzzzz/clipvault/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}
	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-existent ID",
			userID: 2,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cal the newTestDB() helper function to get a connection pool to test database
			db := newTestDB(t)

			// New instance of the UserModel
			m := UserModel{db}

			// Call the UserModel.Exists() method and check that the return value and error mathc
			// the expected values for the sub-test
			exists, err := m.Exists(tt.userID)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}
