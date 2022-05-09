package redisContoller

import (
	"github.com/eminoz/go-microservices/pkg/redis"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	client *redis.Client
}

func RedisClient() *Client {
	client := redisconnection.GetRedisClient()
	return &Client{
		client: client,
	}
}
