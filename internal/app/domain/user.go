package domain

import (
	"crypto/sha512"
	"fmt"
	"time"

	"polygames/internal/app/domain/apperr"
	"polygames/internal/pkg/strhelp"
)

type (
	Gender   int8
	UserRole int16
)

const (
	Male Gender = iota + 1
	Female
)
const (
	UserRoleUser UserRole = iota + 1
	UserRoleModerator
	UserRoleAdmin
	UserRoleGlobalAdmin
)

func (gr Gender) String() string {
	switch gr {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return ""
	}
}

func (ur UserRole) String() string {
	switch ur {
	case UserRoleUser:
		return "User"
	case UserRoleModerator:
		return "Moderator"
	case UserRoleAdmin:
		return "Admin"
	case UserRoleGlobalAdmin:
		return "Global admin"
	default:
		return ""
	}
}

type User struct {
	ID          int32      `json:"id"`
	Name        string     `json:"name"`
	Surname     string     `json:"surname"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Gender      Gender     `json:"gender"`
	Role        UserRole   `json:"role"`
	RoleName    string     `json:"roleName"`
	DateOfBirth time.Time  `json:"dateOfBirth"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DisabledAt  *time.Time `json:"disabledAt,omitempty"`

	Salt            string `json:"-"`
	EncodedPassword string `json:"-"`

	ImageID      string `json:"-"`
	ImageContent []byte `json:"-"`
}

func (u *User) Validate() error {
	var validations []apperr.ValidationError

	if u.Name == "" || len(u.Name) > 30 {
		validations = append(validations, apperr.ValidationError{
			Message: fmt.Sprintf("Name cannot be empty and must not exceed %d characters.", 30),
			Field:   "name",
		})
	}

	if u.Surname == "" || len(u.Surname) > 50 {
		validations = append(validations, apperr.ValidationError{
			Message: fmt.Sprintf("Surname cannot be empty and must not exceed %d characters.", 50),
			Field:   "surname",
		})
	}

	if u.Role.String() == "" {
		validations = append(validations, apperr.ValidationError{
			Message: fmt.Sprintf("Unknown role %d.", u.Role),
			Field:   "role",
		})
	}

	if len(validations) > 0 {
		return apperr.NewValidationError(validations, "")
	}

	return nil
}

func (u *User) EncodePassword() error {
	if u.EncodedPassword == "" {
		return fmt.Errorf("empty password")
	}

	salt, err := strhelp.GenerateRandomString(32)
	if err != nil {
		return fmt.Errorf("generating random string: %w", err)
	}
	u.Salt = salt

	u.EncodedPassword = fmt.Sprintf("%x", sha512.Sum512([]byte(u.EncodedPassword+u.Salt)))

	return nil
}

func (u *User) ComparePassword(password string) bool {
	if password == "" || u.Salt == "" || u.EncodedPassword == "" {
		return false
	}

	passwordHashBytes := sha512.Sum512(append([]byte(password), u.Salt...))
	passwordHash := fmt.Sprintf("%x", passwordHashBytes)

	return passwordHash == u.EncodedPassword
}
