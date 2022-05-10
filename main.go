package main

import (
	"github.com/eminoz/go-microservices/pkg/config"
	"github.com/eminoz/go-microservices/pkg/database"
	"github.com/eminoz/go-microservices/pkg/redis"
	"github.com/eminoz/go-microservices/pkg/router"
)

func main() {
	config.SetupConfig()
	database.SetDatabase()
	redisconnection.NewRedis()
	setup := router.Setup()
	setup.Run()
}
