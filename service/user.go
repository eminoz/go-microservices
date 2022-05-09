package service

import (
	"fmt"
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUser(cnt *gin.Context, user *model.User) (interface{}, error)
	GetOneUser(ctx *gin.Context, filter bson.D) (bson.M, error)
	GetAllUser(ctx *gin.Context) (*[]bson.M, error)
	UpdateUser(ctx *gin.Context, filter *primitive.D, update *primitive.D) (*mongo.UpdateResult, error)
	DeleteOneUser(ctx *gin.Context, filter *primitive.D) *bson.M
}
type UserRedisRepository interface {
	SetUser(ctx *gin.Context, id string, user *model.User) *redis.StatusCmd
	GetUser(ctx *gin.Context, id string) *redis.StringCmd
}
type UserService struct {
	UserRepo      UserRepository
	UserRedisRepo UserRedisRepository
}

func (u *UserService) InsertOneUser(ctx *gin.Context) (interface{}, error) {
	user := new(model.User)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	insertUser, err := u.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}
	var userId = insertUser.(primitive.ObjectID).Hex()
	setUser := u.UserRedisRepo.SetUser(ctx, userId, user)
	fmt.Println(setUser, " user added on redis")
	return insertUser, nil

}

func (u *UserService) GetOneUser(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Param("id")
	id, err2 := primitive.ObjectIDFromHex(userId)
	if err2 != nil {
		return nil, err2
	}
	result, err := u.UserRedisRepo.GetUser(ctx, userId).Result()
	if err == redis.Nil {
		filter := bson.D{{"_id", id}}
		oneUser, err := u.UserRepo.GetOneUser(ctx, filter)
		if err != nil {
			return nil, err
		}
		return &oneUser, nil
	}
	return result, nil

}
func (u *UserService) GetAllUser(ctx *gin.Context) (*[]bson.M, error) {
	user, err := u.UserRepo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (u *UserService) UpdateOneUser(ctx *gin.Context) (*mongo.UpdateResult, error) {
	userId := ctx.Param("id")
	user := new(model.User)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", id}}
	//	update := bson.D{{"$set", bson.D{{"name", user.Name}}}} if you want you can update specific field like this
	update := bson.D{{"$set", user}}
	updateUser, err := u.UserRepo.UpdateUser(ctx, &filter, &update)

	return updateUser, nil
}
func (u *UserService) DeleteOneUser(ctx *gin.Context) (*bson.M, error) {
	userId := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", id}}
	deletedUser := u.UserRepo.DeleteOneUser(ctx, &filter)

	return deletedUser, nil
}
