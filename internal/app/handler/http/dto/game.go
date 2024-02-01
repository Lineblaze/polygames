package dto

import (
	"polygames/internal/app/domain"
)

type (
	CreateGameRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	UpdateGameRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

func (r *CreateGameRequest) ToDomain() *domain.Game {
	if r == nil {
		return nil
	}

	return &domain.Game{
		Name:        r.Name,
		Description: r.Description,
	}
}

func (r *UpdateGameRequest) ToDomain(gameID int32) *domain.Game {
	if r == nil {
		return nil
	}

	return &domain.Game{
		ID:          gameID,
		Name:        r.Name,
		Description: r.Description,
	}
}
