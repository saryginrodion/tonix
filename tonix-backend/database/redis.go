package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func RedisConnect(dsn string) (*redis.Client, error) {
	connOptions, err := redis.ParseURL(dsn)
	if err != nil {
		log.Panicln(err)
	}

	ctx := context.Background()
	conn := redis.NewClient(connOptions)

	if err := conn.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return conn, nil
}
