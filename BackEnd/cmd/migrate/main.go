package main

import (
	"log"
	"os"
	"web_doscom/internal/env"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// load env
	env.LoadEnv()

	if len(os.Args) < 2 {
		log.Fatal("Please provide a command: 'up' or 'down'")
	}

	dirrection := os.Args[1]

	dbURL := os.Getenv("DBURL")

	m, err := migrate.New("file://./migrations", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	switch dirrection {

	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	case "down":
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	default:
		log.Fatal("Unknown command: use 'up' or 'down'")
	}

}
