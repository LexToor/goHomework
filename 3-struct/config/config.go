package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	key := os.Getenv("KEY")
	if key == "" {
		fmt.Println("Error KEY is nil")
		return nil
	}

	return &Config{
		key: key,
	}
}
