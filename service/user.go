package service

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/eminoz/go-microservices/model"
	"github.com/eminoz/go-microservices/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserRepository interface {
	InsertUser(cnt *gin.Context, user *model.User) (interface{}, error)
	GetOneUser(ctx *gin.Context, filter bson.D) (model.Login, error)
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

func (u *UserService) Login(ctx *gin.Context) (interface{}, error) {
	user := new(model.Login)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	filter := bson.D{{"email", &user.Email}}
	oneUser, err := u.UserRepo.GetOneUser(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(oneUser.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}
	//generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = oneUser.Email
	claims["_id"] = oneUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	getConfig := config.GetConfig()

	signedString, err := token.SignedString([]byte(getConfig.AppSecret))
	if err != nil {
		return nil, err
	}

	l := &model.LoginDal{
		ID:    oneUser.ID,
		Email: oneUser.Email,
		Token: signedString,
	}

	return l, nil

}
func (u *UserService) InsertOneUser(ctx *gin.Context) (interface{}, error) {
	user := new(model.User)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, err
	}
	fmt.Println(password)
	user.Password = string(password)

	insertUser, err := u.UserRepo.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}
	var userId = insertUser.(primitive.ObjectID).Hex()
	u.UserRedisRepo.SetUser(ctx, userId, user)
	return insertUser, nil

}

func (u *UserService) GetOneUser(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Param("id")
	id, err2 := primitive.ObjectIDFromHex(userId)
	if err2 != nil {
		return nil, err2
	}
	//fetch data from redis
	//if data does not exist in redis we fetch data from db
	result, err := u.UserRedisRepo.GetUser(ctx, userId).Result()
	if err == redis.Nil {
		filter := bson.D{{"_id", id}}
		//fetch data from database
		oneUser, err := u.UserRepo.GetOneUser(ctx, filter)
		if err != nil {
			return nil, err
		}

		return &oneUser, nil
	}
	//User dal model
	var user model.UserDal
	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil

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
