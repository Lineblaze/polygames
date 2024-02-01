package domain

import (
	"github.com/stretchr/testify/require"
	"polygames/internal/pkg/strhelp"
	"testing"
)

func TestGame_Validate(t *testing.T) {
	tests := []struct {
		name      string
		wantError bool
		game      *Game
	}{
		{
			name:      "empty structure",
			wantError: true,
			game:      &Game{},
		},
		{
			name:      "empty title",
			wantError: true,
			game:      &Game{Description: "something"},
		},
		{
			name:      "description is too long",
			wantError: true,
			game: &Game{
				Description: (func() string {
					s, _ := strhelp.GenerateRandomString(10001)
					return s
				})(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			err := tc.game.Validate()
			if tc.wantError {
				require.Error(tt, err)
				return
			}
			require.NoError(tt, err)
		})
	}
}
