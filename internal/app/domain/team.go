package domain

import (
	"fmt"
	"time"

	"polygames/internal/app/domain/apperr"
)

type (
	Team struct {
		ID          int32      `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		HasImage    bool       `json:"hasImage"`
		CreatedAt   time.Time  `json:"createdAt"`
		UpdatedAt   time.Time  `json:"updatedAt"`
		DisabledAt  *time.Time `json:"disabledAt,omitempty"`

		ImageID      string `json:"-"`
		ImageContent []byte `json:"-"`
	}

	TeamMember struct {
		UserID    int32     `json:"userID"`
		TeamID    int32     `json:"teamID"`
		Role      UserRole  `json:"role"`
		CreatedAt time.Time `json:"createdAt"`
	}
)

func (t *Team) Validate() error {
	var validations []apperr.ValidationError

	if t.Title == "" {
		validations = append(validations, apperr.ValidationError{
			Message: "Title cannot be empty.",
			Field:   "title",
		})
	}
	if len(t.Title) > 128 {
		validations = append(validations, apperr.ValidationError{
			Message: fmt.Sprintf("Title must be less than %d characters.", 128),
			Field:   "title",
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
