package http

import (
	"net/http"
	"polygames/internal/pkg/config"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/rs/cors"
)

func NewHandler(
	userService UserService,
	authService AuthService,
	gameService GameService,
	teamService TeamService,
) http.Handler {
	uh := newUserHandler(userService)
	ah := newAuthHandler(authService, userService)
	gh := newGameHandler(gameService)
	th := newTeamHandler(teamService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(httprate.LimitByIP(69, time.Minute))
	r.Use(middleware.Recoverer)
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   config.Get().Http.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Session-Id"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)
	r.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	r.Use(middleware.NoCache)

	r.Get("/static/*", getStatic)

	// TODO: move to auth middleware handlers to keep API private
	r.Get(`/api/v1/docs`, getApiDocs)
	r.Get(`/api/v1/docs/swagger.json`, getApiDocsSwagger)

	r.Post(`/api/v1/auth/sign-in`, ah.signIn)
	r.Post(`/api/v1/auth/sign-out`, ah.signOut)

	r.Group(func(r chi.Router) {
		//r.Use(ah.authMiddleware)
		// TODO: role middlewares

		// Users
		r.Get(`/api/v1/users/{user_id}`, uh.getUser)
		r.Get(`/api/v1/users`, uh.getUsers)
		r.Post(`/api/v1/users`, uh.createUser)
		r.Put(`/api/v1/users/{user_id}`, uh.updateUser)
		r.Delete(`/api/v1/users/{user_id}`, uh.removeUser)
		r.Post(`/api/v1/users/{user_id}/image`, uh.setUserImage)
		r.Get(`/api/v1/users/{user_id}/image`, uh.getUserImage)

		// Games
		r.Get(`/api/v1/games/{game_id}`, gh.getGame)
		r.Get(`/api/v1/games`, gh.getGames)
		r.Post(`/api/v1/games`, gh.createGame)
		r.Put(`/api/v1/games/{game_id}`, gh.updateGame)
		r.Delete(`/api/v1/games/{game_id}`, gh.deleteGame)
		r.Post(`/api/v1/games/{game_id}/image`, gh.setGameImage)
		r.Get(`/api/v1/games/{game_id}/image`, gh.getGameImage)
		r.Post(`/api/v1/games/{game_id}/file`, gh.setGameFile)
		r.Get(`/api/v1/games/{game_id}/file`, gh.getGameFile)

		// Teams
		r.Get(`/api/v1/teams/{team_id}`, th.getTeam)
		r.Get(`/api/v1/teams`, th.getTeams)
		r.Post(`/api/v1/teams`, th.createTeam)
		r.Put(`/api/v1/teams/{team_id}`, th.updateTeam)
		r.Post(`/api/v1/teams/{team_id}/image`, th.setTeamImage)
		r.Get(`/api/v1/teams/{team_id}/image`, th.getTeamImage)
		r.Post(`/api/v1/teams/{team_id}/disable`, th.disableTeam)
		r.Post(`/api/v1/teams/{team_id}/enable`, th.enableTeam)
	})

	return r
}
