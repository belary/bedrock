package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		AWSRegion:          os.Getenv("AWS_REGION"),
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}, nil
}
