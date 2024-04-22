package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envAWSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func envAWSAccess() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func EnvAWSRegion() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("AWS_REGION")
}

func EnvPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	val := ":" + os.Getenv("PORT")
	if val == ":" {
		return ":6000"
	}
	return val
}