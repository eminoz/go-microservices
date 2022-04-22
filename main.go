package main

import (
	"github.com/eminoz/go-microservices/pkg/config"
	"github.com/eminoz/go-microservices/pkg/database"
	"github.com/eminoz/go-microservices/pkg/router"
)

func main() {
	config.SetupConfig()
	database.SetDatabase()
	setup := router.Setup()
	setup.Run()
}
