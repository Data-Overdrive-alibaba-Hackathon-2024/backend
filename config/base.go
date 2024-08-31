package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func DBName() string {
	return os.Getenv("DB_NAME")
}

func DBUser() string {
	return os.Getenv("DB_USER")
}

func DBPass() string {
	return os.Getenv("DB_PASS")
}

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func ModelUrl() string {
	return os.Getenv("MODEL_URL")
}

func ModelKey() string {
	return os.Getenv("MODEL_KEY")
}
