package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var DB *gorm.DB

var RDB *redis.Client

func InitDB() {
	envPath := filepath.Join("internal", "config", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error loading .env file from %s: %v", envPath, err)
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

func InitRedis() {

	envPath := filepath.Join("internal", "config", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error loading .env file from %s: %v", envPath, err)
		return
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})
}
