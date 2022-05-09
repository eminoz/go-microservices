package api

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserService interface {
	InsertOneUser(ctx *gin.Context) (*mongo.InsertOneResult, error)
	GetOneUser(ctx *gin.Context) (*bson.M, error)
	GetAllUser(ctx *gin.Context) (*[]bson.M, error)
	UpdateOneUser(ctx *gin.Context) (*mongo.UpdateResult, error)
}
type UserController struct {
	UserServices UserService
}

func (c *UserController) UpdateOneUser(ctx *gin.Context) {
	updateOneUser, err := c.UserServices.UpdateOneUser(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, updateOneUser)
}
func (c *UserController) GetAllUser(ctx *gin.Context) {
	users, err := c.UserServices.GetAllUser(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
func (c *UserController) GetOneUser(ctx *gin.Context) {
	oneUser, err := c.UserServices.GetOneUser(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": oneUser})
}

func (c *UserController) InsertOneUser(ctx *gin.Context) {
	insertOneResult, err := c.UserServices.InsertOneUser(ctx)

	if err != nil {
		ctx.JSON(201, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"userId": insertOneResult.InsertedID})

}
