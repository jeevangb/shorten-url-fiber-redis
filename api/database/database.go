package database

import (
	"context"
	"os"

	"github.com/go-redis/redis"
)

var Ctx = context.Background()

func CreateClient(dbno int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDDR"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbno,
	})

	return rdb
}
