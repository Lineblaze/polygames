package domain

import (
	"fmt"
	"time"

	"polygames/internal/app/domain/apperr"
)

type (
	Game struct {
		ID          int32     `json:"id"`
		Name        string    `json:"name"`
		UserID      int32     `json:"userID"`
		TeamID      int32     `json:"teamID"`
		GenreID     int32     `json:"genreID"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`

		ImageID      string `json:"-"`
		ImageContent []byte `json:"-"`

		FileID      string `json:"-"`
		FileContent []byte `json:"-"`
	}
)

func (t *Game) Validate() error {
	var validations []apperr.ValidationError

	if t.Name == "" {
		validations = append(validations, apperr.ValidationError{
			Message: "Name cannot be empty.",
			Field:   "name",
		})
	}

	if t.Description == "" {
		validations = append(validations, apperr.ValidationError{
			Message: "Description cannot be empty.",
			Field:   "description",
		})
	}

	if len(t.Description) > 512 {
		validations = append(validations, apperr.ValidationError{
			Message: fmt.Sprintf("Description length must be less than %d characters.", 512),
			Field:   "description",
		})
	}

	if len(validations) > 0 {
		return apperr.NewValidationError(validations, "")
	}

	return nil
}
