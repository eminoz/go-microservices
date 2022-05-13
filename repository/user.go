package repository

import (
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (c *UserCollection) InsertUser(cnt *gin.Context, user *model.User) (interface{}, error) {
	insertOneResult, err := c.collection.InsertOne(cnt, &user)
	if err != nil {
		return nil, err
	}
	return insertOneResult.InsertedID, nil
}
func (c *UserCollection) GetOneUser(ctx *gin.Context, filter bson.D) (model.Login, error) {

	var user model.Login
	err := c.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return model.Login{}, err
	}
	return user, nil

	/*find, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return find, nil*/
}
func (c *UserCollection) GetOneUserByEmail(ctx *gin.Context, filter bson.D) (bool, model.Email, error) {

	var email model.Email
	err := c.collection.FindOne(ctx, filter).Decode(&email)
	documents, err := c.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, email, nil
	}
	if documents != 1 {
		return false, model.Email{}, nil
	}
	if err != nil {
		return false, model.Email{}, err
	}
	return true, email, nil
}
func (c *UserCollection) GetAllUser(ctx *gin.Context) (*[]bson.M, error) {
	resultUser, err := c.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	resultUser.All(ctx, &results)
	return &results, nil
}
func (c *UserCollection) UpdateUser(ctx *gin.Context, filter *primitive.D, update *primitive.D) (*mongo.UpdateResult, error) {
	updateOne, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil

}
func (c *UserCollection) DeleteOneUser(ctx *gin.Context, filter *primitive.D) *bson.M {
	findOneAndDelete := c.collection.FindOneAndDelete(ctx, filter)
	var deletedDocument bson.M
	findOneAndDelete.Decode(&deletedDocument)
	return &deletedDocument
}
