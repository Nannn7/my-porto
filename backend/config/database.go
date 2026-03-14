package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	if err := LoadDotEnvIfExists(".env", filepath.Join("backend", ".env")); err != nil {
		return fmt.Errorf("failed loading .env file: %w", err)
	}

	if databaseURL := strings.TrimSpace(os.Getenv("DATABASE_URL")); databaseURL != "" {
		database, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect to database via DATABASE_URL: %w", err)
		}

		log.Println("Connected to database")
		DB = database
		return nil
	}

	envKeys := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASS", "DB_NAME", "DB_SSLMODE"}
	values := make(map[string]string, len(envKeys))
	missing := make([]string, 0)

	for _, key := range envKeys {
		value := strings.TrimSpace(os.Getenv(key))
		if value == "" {
			missing = append(missing, key)
			continue
		}
		values[key] = value
	}

	if len(missing) > 0 {
		err := fmt.Errorf("missing required database environment variables: %s", strings.Join(missing, ", "))
		log.Printf("database configuration error: %v", err)
		return err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		values["DB_HOST"],
		values["DB_USER"],
		values["DB_PASS"],
		values["DB_NAME"],
		values["DB_PORT"],
		values["DB_SSLMODE"],
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connected to database")
	DB = database
	return nil
}
