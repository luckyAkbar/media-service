package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
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

func AllowedImageExt() []string {
	var allowedExt []string
	cfg := os.Getenv("ALLOWED_IMAGE_EXT")

	if cfg == "" { // if none, default to this
		return []string{"jpg", "jpeg"}
	}

	tmp := strings.Split(cfg, ",")
	for _, ext := range tmp { // prevent accidental space
		allowedExt = append(allowedExt, strings.ToLower(strings.Trim(ext, " ")))
	}

	return allowedExt
}

func AllowedMimeType() []string {
	var allowedMimeType []string
	cfg := os.Getenv("ALLOWED_IMAGE_MIME_TYPE")

	if cfg == "" { // if none, default to this
		return []string{"image/jpeg"}
	}

	tmp := strings.Split(cfg, ",")
	for _, ext := range tmp { // prevent accidental space
		allowedMimeType = append(allowedMimeType, strings.ToLower(strings.Trim(ext, " ")))
	}

	return allowedMimeType
}

func MaxImageSizeBytes() int64 {
	var MB int64 = 1048576 // MB as bytes
	maxSizeMB, err := strconv.Atoi(os.Getenv("MAX_IMAGE_SIZE_MB"))

	if err != nil {
		logrus.Warn("MAX_IMAGE_SIZE_MB configuration is invalid. Using default value.")
		return MB * 3
	}

	return int64(maxSizeMB) * MB // will return max size in bytes
}

func ImageNameLength() int {
	cfg, err := strconv.Atoi(os.Getenv("IMAGE_NAME_LENGTH"))
	if err != nil {
		logrus.Warn("IMAGE_NAME_LENGTH configuration is error. Using default value")

		return 35
	}

	return cfg
}

func ImageStoragePath() string {
	cfg := os.Getenv("IMAGE_STORAGE_PATH")

	if cfg == "" {
		logrus.Warn("IMAGE_STORAGE_PATH value is unset. Using default value.")
		return "/var/local/image_service_server/image_storage"
	}

	return cfg
}
