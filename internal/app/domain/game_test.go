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
			name:      "title is too long",
			wantError: true,
			game: &Game{Title: (func() string {
				s, _ := strhelp.GenerateRandomString(129)
				return s
			})()},
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
		{
			name:      "invalid link",
			wantError: true,
			game: &Game{
				Title: "title",
				Link:  "qwefqwef",
			},
		},
		{
			name:      "invalid link #2",
			wantError: true,
			game: &Game{
				Title: "title",
				Link:  "www.mysite.com",
			},
		},
		{
			name:      "valid link #1",
			wantError: false,
			game: &Game{
				Title:       "title",
				Description: "some description",
				Link:        "/relative/path/1",
			},
		},
		{
			name:      "valid link #2",
			wantError: false,
			game: &Game{
				Title:       "title",
				Description: "some description",
				Link:        "https://something.com/relative/path/2",
			},
		},
		{
			name:      "valid link #3",
			wantError: false,
			game: &Game{
				Title:       "title",
				Description: "some description",
				Link:        "http://10.0.0.0:8443",
			},
		},
		{
			name:      "valid link #4",
			wantError: false,
			game: &Game{
				Title:       "title",
				Description: "some description",
				Link:        "https://google.com",
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
