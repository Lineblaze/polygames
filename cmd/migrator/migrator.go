package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"polygames/pkg/postgres"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"polygames/internal/pkg/config"
	"polygames/internal/pkg/wcrypto"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", "config.yml", "Path to application config file.")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("use up or down syntax to make migration: ./migrate up")
		return
	}

	config.Read(configPath)
	cfg := config.Get()

	user, password, err := wcrypto.DecodeUserPass(cfg.Database.User, cfg.Database.Password, config.Block)
	if err != nil {
		log.Fatalf("decondig db credentials: %v", err)
	}

	dbConnString := postgres.ConnectionString(user, password, cfg.Database.Host, cfg.Database.Database)

	m, err := migrate.New("file://migrations", dbConnString)
	if err != nil {
		log.Fatalf("creating migration: %v", err)
	}

	action := os.Args[1]

	switch action {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "drop":
		err = m.Drop()
	case "drop-up":
		err = m.Drop()
		err = m.Up()
	default:
		fmt.Println("unknown action:", action)
	}

	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("Migrate: No changes")
		}
		log.Fatalf("applying migration: %v", err)
	}

	fmt.Printf("Migrate %s done.\n", action)
}
