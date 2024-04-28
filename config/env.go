package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host                   string
	Port                   string
	DbUser                 string
	DbPassword             string
	DbAddress              string
	DbName                 string
	JwtExpirationInSeconds int64
	JwtSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "8080")

	return Config{
		Host:                   getEnv("HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DbUser:                 getEnv("DB_USER", ""),
		DbPassword:             getEnv("DB_PASSWORD", ""),
		DbAddress:              fmt.Sprintf("%s:%s", dbHost, dbPort),
		DbName:                 getEnv("DB_NAME", ""),
		JwtExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7), // 7 days
		JwtSecret:              getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	value, ok := os.LookupEnv(key)
	if ok {
		intNum, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return intNum
		}
	}

	return fallback
}
