package domain

import (
	"fmt"
	"net/url"
	"time"

	"polygames/internal/app/domain/apperr"
)

type (
	Game struct {
		ID           int32     `json:"id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		UserID       int32     `json:"userID"`
		TeamID       int32     `json:"teamID"`
		GenreID      int32     `json:"genreID"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
		ImageID      string    `json:"-"`
		ImageContent []byte    `json:"-"`
		FileID       string    `json:"-"`
		FileContent  []byte    `json:"-"`
		Link         string    `json:"link,omitempty"`
	}
)

func (g *Game) Validate() error {
	var validations []apperr.ValidationError

	if g.Title == "" {
		validations = append(validations, apperr.ValidationError{
			Message: "Title cannot be empty.",
			Field:   "title",
		})
	}

	if g.Description == "" {
		validations = append(validations, apperr.ValidationError{
			Message: "Description cannot be empty.",
			Field:   "description",
		})
	}

	if len(g.Description) > 512 {
		validations = append(validations, apperr.ValidationError{
			Message: fmt.Sprintf("Description length must be less than %d characters.", 512),
			Field:   "description",
		})
	}

	if g.Link != "" {
		_, err := url.ParseRequestURI(g.Link)
		if err != nil {
			validations = append(validations, apperr.ValidationError{
				Message: "Link has invalid format.",
				Field:   "link",
			})
		}
	}

	if len(validations) > 0 {
		return apperr.NewValidationError(validations, "")
	}

	return nil
}
