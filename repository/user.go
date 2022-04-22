package repository

import (
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c UserCollection) InsertUser(cnt *gin.Context, user *model.User) (*mongo.InsertOneResult, error) {
	insertOneResult, err := c.collection.InsertOne(cnt, user)
	if err != nil {
		return nil, err
	}
	return insertOneResult, nil
}
