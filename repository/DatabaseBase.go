package repository

import (
	"github.com/eminoz/go-microservices/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func UserCollectionSetting() *UserCollection {
	getDatabase := database.GetDatabase()
	return &UserCollection{
		db:         getDatabase,
		collection: getDatabase.Collection("go-microservice"),
	}
}
