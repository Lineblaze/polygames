package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"polygames/internal/app/domain"
	"polygames/internal/app/domain/apperr"
	"polygames/internal/app/infrastructure/repository"
	"polygames/internal/pkg/auth/session"
)

//go:generate mockgen -source=auth.go -destination=./mocks/auth.go -package=mocks
type AuthRepository interface {
	GetUserByLogin(ctx context.Context, login string) (*domain.User, error)
	GetActiveUser(ctx context.Context, id int32) (*domain.User, error)
}

type AuthService struct {
	repo AuthRepository
}

func NewAuthService(repo AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) SignIn(ctx context.Context, req *domain.SignInRequest) (*domain.SignInResponse, error) {
	user, err := s.repo.GetUserByLogin(ctx, req.Login)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperr.NewInvalidRequest("Invalid credentials.", "")
		}
		return nil, fmt.Errorf("getting user by login: %w", err)
	}

	if !user.ComparePassword(req.Password) {
		slog.Debug("Invalid password")
		return nil, apperr.NewInvalidRequest("Invalid credentials.", "")
	}

	sessionID, csrfToken, err := session.New(user.ID)
	if err != nil {
		return nil, fmt.Errorf("creating new session: %w", err)
	}

	return &domain.SignInResponse{
		SessionID: sessionID,
		CSRFToken: csrfToken,
		UserID:    user.ID,
	}, nil
}

func (s *AuthService) SignOut(_ context.Context, sessionID string) {
	session.Delete(sessionID)
}

func (s *AuthService) CheckUserExists(ctx context.Context, id int32) error {
	_, err := s.repo.GetActiveUser(ctx, id)
	if err != nil {
		return fmt.Errorf("getting active user: %w", err)
	}
	return nil
}
