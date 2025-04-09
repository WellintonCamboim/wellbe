package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort    int
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	JWTSecret     string
}

func LoadConfig(path string) (*Config, error) {
	err := godotenv.Load(path + "/app.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		port = 8080
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}

	return &Config{
		ServerPort:    port,
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        dbPort,
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
	}, nil
}