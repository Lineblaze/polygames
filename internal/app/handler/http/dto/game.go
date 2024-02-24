package dto

import (
	"polygames/internal/app/domain"
)

type (
	CreateGameRequest struct {
		Title       string `json:"name"`
		Description string `json:"description"`
		UserID      int32  `json:"userID"`
		TeamID      int32  `json:"teamID"`
		GenreID     int32  `json:"genreID"`
		Link        string `json:"link,omitempty"`
	}
	UpdateGameRequest struct {
		Title       string `json:"name"`
		Description string `json:"description"`
		UserID      int32  `json:"userID"`
		TeamID      int32  `json:"teamID"`
		GenreID     int32  `json:"genreID"`
		Link        string `json:"link,omitempty"`
	}
)

func (r *CreateGameRequest) ToDomain() *domain.Game {
	if r == nil {
		return nil
	}

	return &domain.Game{
		Title:       r.Title,
		Description: r.Description,
	}
}

func (r *UpdateGameRequest) ToDomain(gameID int32) *domain.Game {
	if r == nil {
		return nil
	}

	return &domain.Game{
		ID:          gameID,
		Title:       r.Title,
		Description: r.Description,
	}
}
