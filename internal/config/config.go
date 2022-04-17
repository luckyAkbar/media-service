package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func PostgresDSN() string {
	host := os.Getenv("PGHOST")
	db := os.Getenv("PGDATABASE")
	user := os.Getenv("PGUSER")
	pw := os.Getenv("PGPASSWORD")
	port := os.Getenv("PGPORT")

	if os.Getenv("ENV") == "production" {
		host = "host.docker.internal" // in the production, will be using docker.
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pw, db, port)

	return dsn

}

func LogLevel() string {
	cfg := os.Getenv("LOG_LEVEL")

	if cfg == "" {
		return "debug"
	}

	return cfg
}

func ServerPort() string {
	cfg := os.Getenv("SERVER_PORT")

	if cfg == "" {
		return "5000" // default port
	}

	return cfg
}
