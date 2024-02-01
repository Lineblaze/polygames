package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUser_ComparePassword(t *testing.T) {
	tests := []struct {
		name     string
		result   bool
		encode   bool
		password string
		user     *User
	}{
		{
			name:     "empty structure",
			result:   false,
			password: "123",
			user:     &User{},
		},
		{
			name:     "empty password",
			result:   false,
			password: "",
			user:     &User{EncodedPassword: "123", Salt: "321"},
		},
		{
			name:     "empty encoded password",
			result:   false,
			password: "123",
			user:     &User{Salt: "321"},
		},
		{
			name:     "empty salt",
			result:   false,
			password: "123",
			user:     &User{EncodedPassword: "123"},
		},
		{
			name:     "passwords match",
			result:   true,
			password: "123",
			encode:   true,
			user:     &User{EncodedPassword: "123"},
		},
		{
			name:     "passwords dont match",
			result:   false,
			password: "1234",
			encode:   true,
			user:     &User{EncodedPassword: "123"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			if tc.encode {
				err := tc.user.EncodePassword()
				require.NoError(tt, err)
			}

			res := tc.user.ComparePassword(tc.password)
			require.Equal(tt, tc.result, res)
		})
	}
}
