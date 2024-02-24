package main

import (
	"context"
	"log"
	"polygames/internal/app/domain"
	"polygames/internal/app/infrastructure/repository/filesystem"
	"polygames/internal/app/infrastructure/repository/postgresql"
	"polygames/internal/app/service"
	"polygames/internal/pkg/config"
	"polygames/internal/pkg/wcrypto"
	"polygames/pkg/postgres"
)

func main() {
	config.Read("config.yml")

	cfg := config.Get()

	user, password, err := wcrypto.DecodeUserPass(cfg.Database.User, cfg.Database.Password, config.Block)
	if err != nil {
		log.Fatalf("decoding database username: %v", err)
	}

	dbConnString := postgres.ConnectionString(user, password, "45.141.78.221", cfg.Database.Database)

	pg, err := postgres.New(context.Background(), dbConnString)
	if err != nil {
		log.Fatalf("creating postgres: %v", err)
	}

	log.Println("Connected to database")

	userRepo := postgresql.NewUserRepository(pg.Pool)
	fs, _ := filesystem.New("")
	userService := service.NewUserService(userRepo, fs)

	_, err = userService.CreateUser(context.Background(), &domain.User{
		Name:            "test_name",
		Surname:         "test_surname",
		Username:        "test_username",
		Email:           "test@mail.com",
		EncodedPassword: "password123",
		Role:            domain.UserRoleGlobalAdmin,
	})
	if err != nil {
		log.Fatalf("Creating user: %v", err)
	}

	log.Println("Done!")
}
