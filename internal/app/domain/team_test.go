package domain

import (
	"github.com/stretchr/testify/require"
	"polygames/internal/pkg/strhelp"
	"testing"
)

func TestTeam_Validate(t *testing.T) {
	tests := []struct {
		name      string
		wantError bool
		team      *Team
	}{
		{
			name:      "empty structure",
			wantError: true,
			team:      &Team{},
		},
		{
			name:      "empty title",
			wantError: true,
			team:      &Team{Description: "something"},
		},
		{
			name:      "title is too long",
			wantError: true,
			team: &Team{Title: (func() string {
				s, _ := strhelp.GenerateRandomString(129)
				return s
			})()},
		},
		{
			name:      "description is too long",
			wantError: true,
			team: &Team{
				Title: "title",
				Description: (func() string {
					s, _ := strhelp.GenerateRandomString(10001)
					return s
				})(),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			err := tc.team.Validate()
			if tc.wantError {
				require.Error(tt, err)
				return
			}
			require.NoError(tt, err)
		})
	}
}
