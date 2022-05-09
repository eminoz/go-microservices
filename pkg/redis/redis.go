package redisconnection

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

var Redis *redis.Client

func NewRedis() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	Redis = client
	return &Client{
		client: client,
	}, nil
}
func GetRedisClient() *redis.Client {
	return Redis
}
