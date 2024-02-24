package postgresql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"

	"polygames/internal/app/domain"
	"polygames/internal/app/infrastructure/repository"
)

type GameRepository struct {
	pool Driver
}

func NewGameRepository(pool Driver) *GameRepository {
	return &GameRepository{pool: pool}
}

func (r *GameRepository) GetGame(ctx context.Context, id int32) (*domain.Game, error) {
	var game domain.Game

	err := r.pool.QueryRow(ctx, `
		SELECT 
		    id, title, description, user_id, team_id, genre_id, created_at, updated_at, image_id, file_id, link
		FROM games
		WHERE id=$1`, id).Scan(
		&game.ID,
		&game.Title,
		&game.Description,
		&game.UserID,
		&game.TeamID,
		&game.GenreID,
		&game.CreatedAt,
		&game.UpdatedAt,
		&game.ImageID,
		&game.FileID,
		&game.Link,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrObjectNotFound
		}
		return nil, fmt.Errorf("scanning game: %w", err)
	}
	return &game, nil
}

func (r *GameRepository) GetGames(ctx context.Context) ([]domain.Game, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT 
		    team_id, genre_id, created_at, updated_at
		FROM games
		ORDER BY created_at`)
	if err != nil {
		return nil, fmt.Errorf("selecting games: %w", err)
	}
	defer rows.Close()

	var games []domain.Game
	for rows.Next() {
		var game domain.Game

		err = rows.Scan(
			&game.TeamID,
			&game.GenreID,
			&game.CreatedAt,
			&game.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning game: %w", err)
		}

		games = append(games, game)
	}

	return games, nil
}

func (r *GameRepository) CreateGame(ctx context.Context, game *domain.Game) (int32, error) {
	var id int32

	err := r.pool.QueryRow(ctx, `
		INSERT INTO games(title, description, user_id, team_id, genre_id, created_at, updated_at, image_id, file_id, link)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`,
		game.Title,
		game.Description,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("inserting game: %w", err)
	}

	return id, nil
}

func (r *GameRepository) UpdateGame(ctx context.Context, game *domain.Game) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE games
		SET title=$2, description=$3, updated_at=now(), link=$4
		WHERE id=$1`,
		game.ID,
		game.Title,
		game.Description,
		game.Link,
	)
	if err != nil {
		return fmt.Errorf("updating game: %w", err)
	}

	return nil
}

func (r *GameRepository) SetGameImage(ctx context.Context, gameID int32, imageID string) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE games
		SET image_id=$2, updated_at=now()
		WHERE id=$1`,
		gameID,
		imageID,
	)
	if err != nil {
		return fmt.Errorf("updating game: %w", err)
	}

	return nil
}

func (r *GameRepository) SetGameFile(ctx context.Context, gameID int32, fileID string) error {
	_, err := r.pool.Exec(ctx, `
		UPDATE games
		SET file_id=$2, updated_at=now()
		WHERE id=$1`,
		gameID,
		fileID,
	)
	if err != nil {
		return fmt.Errorf("updating game: %w", err)
	}

	return nil
}

func (r *GameRepository) DeleteGame(ctx context.Context, gameID int32) error {
	_, err := r.pool.Exec(ctx, `
		DELETE FROM games
		WHERE id=$1 `, gameID)
	if err != nil {
		return fmt.Errorf("deleting game: %w", err)
	}

	return nil
}

func (r *GameRepository) CheckGameUniqueness(ctx context.Context, title string) (*domain.Game, error) {
	var game domain.Game

	err := r.pool.QueryRow(ctx, `
		SELECT 
		    id, title, description, created_at, updated_at, image_id, file_id
		FROM games
		WHERE lower(title)=lower($1)`, title).Scan(
		&game.ID,
		&game.Title,
		&game.Description,
		&game.CreatedAt,
		&game.UpdatedAt,
		&game.ImageID,
		&game.FileID,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrObjectNotFound
		}
		return nil, fmt.Errorf("scanning game: %w", err)
	}

	return &game, nil
}
