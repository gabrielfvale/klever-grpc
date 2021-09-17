package pkg

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads a .env file to os.Environ()
func LoadEnv() error {
	return godotenv.Load()
}

// GetEnvVar returns an environment variable, given a key
func GetEnvVar(key string) string {
	return os.Getenv(key)
}
