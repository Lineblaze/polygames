package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"

	"polygames/internal/app/domain"
	"polygames/internal/app/domain/apperr"
	"polygames/internal/app/infrastructure/repository"
)

//go:generate mockgen -source=game.go -destination=./mocks/game.go -package=mocks
type GameRepository interface {
	GetGame(ctx context.Context, id int32) (*domain.Game, error)
	GetGames(ctx context.Context) ([]domain.Game, error)
	CreateGame(ctx context.Context, game *domain.Game) (int32, error)
	UpdateGame(ctx context.Context, game *domain.Game) error
	SetGameImage(ctx context.Context, gameID int32, imageID string) error
	SetGameFile(ctx context.Context, gameID int32, fileID string) error
	DeleteGame(ctx context.Context, gameID int32) error
	CheckGameUniqueness(ctx context.Context, title string) (*domain.Game, error)
}

type GameService struct {
	filesDir string
	repo     GameRepository
	fileRepo FileRepository
}

func NewGameService(repo GameRepository, fileRepo FileRepository) *GameService {
	return &GameService{"games", repo, fileRepo}
}

func (s *GameService) GetGame(ctx context.Context, id int32) (*domain.Game, error) {
	game, err := s.repo.GetGame(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperr.NewNotFound("game_id")
		}
		return nil, fmt.Errorf("getting game %d: %w", id, err)
	}

	return game, nil
}

func (s *GameService) GetGames(ctx context.Context) ([]domain.Game, error) {
	games, err := s.repo.GetGames(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting games: %w", err)
	}

	return games, nil
}

func (s *GameService) CreateGame(ctx context.Context, game *domain.Game) (*domain.Game, error) {
	if err := game.Validate(); err != nil {
		return nil, fmt.Errorf("validating game: %w", err)
	}

	foundGame, err := s.repo.CheckGameUniqueness(ctx, game.Title)
	if err != nil && !errors.Is(err, repository.ErrObjectNotFound) {
		return nil, fmt.Errorf("checking game uniqueness: %w", err)
	}
	if foundGame != nil {
		return nil, apperr.NewDuplicate("Title already taken.", "title")
	}

	id, err := s.repo.CreateGame(ctx, game)
	if err != nil {
		return nil, fmt.Errorf("creating game: %w", err)
	}

	createdGame, err := s.repo.GetGame(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting game %d: %w", id, err)
	}

	return createdGame, nil
}

func (s *GameService) UpdateGame(ctx context.Context, game *domain.Game) (*domain.Game, error) {
	if err := game.Validate(); err != nil {
		return nil, fmt.Errorf("validating game: %w", err)
	}

	_, err := s.repo.GetGame(ctx, game.ID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperr.NewNotFound("game_id")
		}
		return nil, fmt.Errorf("getting game %d: %w", game.ID, err)
	}

	if game.Title != "" {
		foundGame, err := s.repo.CheckGameUniqueness(ctx, game.Title)
		if err != nil && !errors.Is(err, repository.ErrObjectNotFound) {
			return nil, fmt.Errorf("checking game uniqueness: %w", err)
		}
		if foundGame != nil {
			return nil, apperr.NewDuplicate("Title already taken.", "title")
		}
	}

	err = s.repo.UpdateGame(ctx, game)
	if err != nil {
		return nil, fmt.Errorf("creating game: %w", err)
	}

	updatedGame, err := s.repo.GetGame(ctx, game.ID)
	if err != nil {
		return nil, fmt.Errorf("getting game %d: %w", game.ID, err)
	}

	return updatedGame, nil
}

func (s *GameService) SetGameImage(ctx context.Context, gameID int32, img []byte) error {
	if len(img) > 5<<20 {
		return apperr.NewInvalidRequest("Image is too big.", "file")
	}

	mt := mimetype.Detect(img)
	if !mt.Is("image/jpeg") &&
		!mt.Is("image/png") &&
		!mt.Is("image/webp") {
		return apperr.NewInvalidRequest("Invalid image mime type.", "file")
	}

	fileID := uuid.New().String()
	fileName := fileID + mt.Extension()

	game, err := s.repo.GetGame(ctx, gameID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return apperr.NewNotFound("game_id")
		}
		return fmt.Errorf("getting game %d: %w", gameID, err)
	}

	err = s.repo.SetGameImage(ctx, gameID, fileName)
	if err != nil {
		return fmt.Errorf("setting game %d image: %w", gameID, err)
	}

	err = s.fileRepo.Save(ctx, img, filepath.Join(s.filesDir, fileName))
	if err != nil {
		return fmt.Errorf("saving game image: %w", err)
	}

	if game.ImageID != "" {
		err = s.fileRepo.Delete(ctx, filepath.Join(s.filesDir, game.ImageID))
		if err != nil {
			slog.Error("Deleting old game image", slog.String("error", err.Error()))
		}
	}

	return nil
}

func (s *GameService) GetGameImage(ctx context.Context, gameID int32) (*domain.Game, error) {
	game, err := s.repo.GetGame(ctx, gameID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperr.NewNotFound("game_id")
		}
		return nil, fmt.Errorf("getting game %d: %w", gameID, err)
	}

	if game.ImageID == "" {
		return nil, apperr.NewNotFound("image_id")
	}

	data, err := s.fileRepo.Read(ctx, filepath.Join(s.filesDir, game.ImageID))
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperr.NewNotFound("image_id")
		}
		return nil, fmt.Errorf("reading game image: %w", err)
	}

	game.ImageContent = data

	return game, nil
}

func (s *GameService) SetGameFile(ctx context.Context, gameID int32, file []byte) error {
	if len(file) > 5<<20 {
		return apperr.NewInvalidRequest("File is too big.", "file")
	}

	mt := mimetype.Detect(file)
	if !mt.Is("file/exe") {
		return apperr.NewInvalidRequest("Invalid file mime type.", "file")
	}

	fileID := uuid.New().String()
	fileName := fileID + mt.Extension()

	game, err := s.repo.GetGame(ctx, gameID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return apperr.NewNotFound("game_id")
		}
		return fmt.Errorf("getting game %d: %w", gameID, err)
	}

	err = s.repo.SetGameFile(ctx, gameID, fileName)
	if err != nil {
		return fmt.Errorf("setting game %d file: %w", gameID, err)
	}

	err = s.fileRepo.Save(ctx, file, filepath.Join(s.filesDir, fileName))
	if err != nil {
		return fmt.Errorf("saving game file: %w", err)
	}

	if game.FileID != "" {
		err = s.fileRepo.Delete(ctx, filepath.Join(s.filesDir, game.FileID))
		if err != nil {
			slog.Error("Deleting old game file", slog.String("error", err.Error()))
		}
	}

	return nil
}

func (s *GameService) GetGameFile(ctx context.Context, gameID int32) (*domain.Game, error) {
	game, err := s.repo.GetGame(ctx, gameID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperr.NewNotFound("game_id")
		}
		return nil, fmt.Errorf("getting game %d: %w", gameID, err)
	}

	if game.FileID == "" {
		return nil, apperr.NewNotFound("file_id")
	}

	data, err := s.fileRepo.Read(ctx, filepath.Join(s.filesDir, game.FileID))
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, apperr.NewNotFound("file_id")
		}
		return nil, fmt.Errorf("reading game file: %w", err)
	}

	game.FileContent = data

	return game, nil
}

func (s *GameService) DeleteGame(ctx context.Context, gameID int32) error {
	game, err := s.repo.GetGame(ctx, gameID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return apperr.NewNotFound("game_id")
		}
		return fmt.Errorf("getting game %d: %w", gameID, err)
	}
	err = s.repo.DeleteGame(ctx, gameID)
	if err != nil {
		return fmt.Errorf("deleting game %d: %w", gameID, err)
	}

	err = s.fileRepo.Delete(ctx, filepath.Join(s.filesDir, game.FileID))
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return apperr.NewNotFound("game_id")
		}
		return fmt.Errorf("deleting game %d fs: %w", gameID, err)
	}

	return nil
}
