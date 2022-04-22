package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserService interface {
	InsertOneUser(ctx *gin.Context) (*mongo.InsertOneResult, error)
}
type UserController struct {
	UserServices UserService
}

func (c UserController) InsertOneUser(ctx *gin.Context) {
	insertOneResult, err := c.UserServices.InsertOneUser(ctx)

	if err != nil {
		ctx.JSON(201, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"userId": insertOneResult.InsertedID})

}
