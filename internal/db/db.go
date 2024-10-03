package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"
	"log"
)

var (
	PostgresConn *pgx.Conn
	RedisClient  *redis.Client
)

func InitPostgres(connString string) {
	var err error
	PostgresConn, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected to PostgreSQL!")
}

func InitRedis(opts *redis.Options) {
	RedisClient = redis.NewClient(opts)
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v\n", err)
	}
	log.Println("Connected to Redis!")
}
