package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load(`.env`)
}

// UseEnv returns the value of the environment variable envName or fallback if it is not set
func UseEnv(envName string, fallback string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		os.Setenv(envName, fallback)
		return fallback
	}
	return os.Getenv(envName)
}
