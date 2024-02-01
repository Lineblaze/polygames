package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"polygames/internal/app/domain"
	"polygames/internal/app/handler/http/dto"
	"polygames/internal/app/handler/http/httphelp"
)

//go:generate mockgen -source=game.go -destination=./mocks/game.go -package=mocks
type GameService interface {
	GetGame(ctx context.Context, id int32) (*domain.Game, error)
	GetGames(ctx context.Context) ([]domain.Game, error)
	CreateGame(ctx context.Context, game *domain.Game) (*domain.Game, error)
	UpdateGame(ctx context.Context, game *domain.Game) (*domain.Game, error)
	SetGameImage(ctx context.Context, gameID int32, img []byte) error
	GetGameImage(ctx context.Context, gameID int32) (*domain.Game, error)
	DisableGame(ctx context.Context, gameID int32) error
	EnableGame(ctx context.Context, gameID int32) error
}

type GameHandler struct {
	GameService GameService
}

func newGameHandler(service GameService) *GameHandler {
	return &GameHandler{service}
}

// getGame godoc
// @Summary      Get game by identifier
// @Description  Returns information about single game.
// @Tags         Games
// @Produce      json
// @Param        Game_id path int true "Game identifier."
// @Success      200  {object}  domain.Game
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games/{game_id} [get]
func (h *GameHandler) getGame(w http.ResponseWriter, r *http.Request) {
	tid := httphelp.ParseParamInt32("game_id", r)

	response, err := h.GameService.GetGame(r.Context(), tid)
	if err != nil {
		httphelp.SendError(err, w)
		return
	}

	httphelp.SendJSON(http.StatusOK, response, w)
}

// getGames godoc
// @Summary      Get games
// @Description  Returns list of games.
// @Tags         Games
// @Produce      json
// @Success      200  {array}  domain.Game
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games [get]
func (h *GameHandler) getGames(w http.ResponseWriter, r *http.Request) {
	response, err := h.GameService.GetGames(r.Context())
	if err != nil {
		httphelp.SendError(err, w)
		return
	}

	httphelp.SendJSON(http.StatusOK, response, w)
}

// createGame godoc
// @Summary      Create game
// @Description  Creates a new game. Returns an object with information about created game.
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateGameRequest true "Request body."
// @Success      200  {object}	domain.Game
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games [post]
func (h *GameHandler) createGame(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateGameRequest
	if err := httphelp.ReadJSON(&req, r); err != nil {
		httphelp.SendError(err, w)
		return
	}

	response, err := h.GameService.CreateGame(r.Context(), req.ToDomain())
	if err != nil {
		httphelp.SendError(err, w)
		return
	}

	httphelp.SendJSON(http.StatusOK, response, w)
}

// updateGame godoc
// @Summary      Update game
// @Description  Updates a game.
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param        game_id path int true "game identifier."
// @Param        request body dto.UpdateGameRequest true "Request body."
// @Success      200  {object}	domain.Game
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games/{game_id} [put]
func (h *GameHandler) updateGame(w http.ResponseWriter, r *http.Request) {
	tid := httphelp.ParseParamInt32("game_id", r)

	var req dto.UpdateGameRequest
	if err := httphelp.ReadJSON(&req, r); err != nil {
		httphelp.SendError(err, w)
		return
	}

	response, err := h.GameService.UpdateGame(r.Context(), req.ToDomain(tid))
	if err != nil {
		httphelp.SendError(err, w)
		return
	}

	httphelp.SendJSON(http.StatusOK, response, w)
}

// setGameImage godoc
// @Summary      Set game image
// @Description  Updated game image. Accepts `multipart/form-data`.
// @Description
// @Description  Note: if a game already has an image, it will be deleted automatically on success.
// @Tags         Games
// @Accept       mpfd
// @Param        game_id path int true "Game identifier."
// @Param        file formData file true "Image file. MUST have one of the following mime types: [`image/jpeg`, `image/png`, `image/webp`]"
// @Success      200
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games/{game_id}/image [post]
func (h *GameHandler) setGameImage(w http.ResponseWriter, r *http.Request) {
	tid := httphelp.ParseParamInt32("game_id", r)

	file, _, err := r.FormFile("file")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			// TODO: custom http error
			httphelp.SendError(fmt.Errorf("file is not presented"), w)
			return
		}
		httphelp.SendError(fmt.Errorf("parsing form file: %w", err), w)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		httphelp.SendError(fmt.Errorf("reading file: %w", err), w)
		return
	}

	err = h.GameService.SetGameImage(r.Context(), tid, content)
	if err != nil {
		httphelp.SendError(fmt.Errorf("setting game image: %w", err), w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// getGameImage godoc
// @Summary      Get game image content
// @Description  Returns game image.
// @Tags         Games
// @Produce      octet-stream
// @Param        game_id path int true "game identifier."
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games/{game_id}/image [get]
func (h *GameHandler) getGameImage(w http.ResponseWriter, r *http.Request) {
	tid := httphelp.ParseParamInt32("game_id", r)

	response, err := h.GameService.GetGameImage(r.Context(), tid)
	if err != nil {
		httphelp.SendError(fmt.Errorf("getting game image: %w", err), w)
		return
	}

	fileName := fmt.Sprintf("%s.%s", response.Name, filepath.Ext(response.ImageID))

	http.ServeContent(w, r, fileName, response.UpdatedAt, bytes.NewReader(response.ImageContent))
}

// disableGame godoc
// @Summary      Disable game
// @Description  Disables a game.
// @Tags         Games
// @Param        game_id path int true "game identifier."
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games/{game_id}/disable [post]
func (h *GameHandler) disableGame(w http.ResponseWriter, r *http.Request) {
	tid := httphelp.ParseParamInt32("game_id", r)

	err := h.GameService.DisableGame(r.Context(), tid)
	if err != nil {
		httphelp.SendError(fmt.Errorf("disabling game: %w", err), w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// enableGame godoc
// @Summary      Enable game
// @Description  Enables a game.
// @Tags         Games
// @Param        Game_id path int true "game identifier."
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /api/v1/games/{game_id}/enable [post]
func (h *GameHandler) enableGame(w http.ResponseWriter, r *http.Request) {
	tid := httphelp.ParseParamInt32("game_id", r)

	err := h.GameService.EnableGame(r.Context(), tid)
	if err != nil {
		httphelp.SendError(fmt.Errorf("enabling game: %w", err), w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
