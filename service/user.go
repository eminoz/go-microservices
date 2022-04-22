package service

import (
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUser(cnt *gin.Context, user *model.User) (*mongo.InsertOneResult, error)
}
type UserService struct {
	UserRepo UserRepository
}

func (u *UserService) InsertOneUser(ctx *gin.Context) (*mongo.InsertOneResult, error) {
	user := new(model.User)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	insertUser, err := u.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return insertUser, nil

}
