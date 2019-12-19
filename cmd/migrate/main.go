package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	up   = "up"
	down = "down"
)

func main() {
	steps := 0
	flag.Parse()
	args := flag.Args()

	for _, arg := range args {
		if arg == up {
			steps = 1
			break
		}

		if arg == down {
			steps = -1
			break
		}
	}

	m, err := migrate.New(
		"file://cmd/migrate/sql",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_SSL"),
		),
	)

	if err != nil {
		log.Fatalf("Error preparing database migrations: %v", err.Error())
	}

	err = m.Steps(steps)

	if err != nil {
		log.Fatalf("Error running database migrations: %v", err.Error())
	}

	log.Printf("Successfully ran migration steps: %d", steps)
}
