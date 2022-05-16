package repository

import (
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (o *OrderCollection) CreateAOrder(ctx *gin.Context, order *model.Order) (interface{}, error) {

	insertOne, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return insertOne, nil
}

func (o *OrderCollection) GetUsersOrders(ctx *gin.Context, filter bson.D) (model.Order, error) {
	var order model.Order
	err := o.collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		return order, nil
	}
	return order, nil
}

func (o *OrderCollection) UpdateOneOrder(ctx *gin.Context, filter bson.D, update bson.D) (interface{}, error) {
	updateOne, err := o.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}

func (o *OrderCollection) FindOrderAndUpdate(ctx *gin.Context, filter bson.D, update bson.D) interface{} {
	andUpdate := o.collection.FindOneAndUpdate(ctx, filter, update)
	return andUpdate
}
func (o *OrderCollection) Find(ctx *gin.Context, filter bson.D) (*mongo.Cursor, error) {

	find, err := o.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	return find, nil
}
func (o *OrderCollection) DeleteOneOrder(ctx *gin.Context, filter bson.D) (interface{}, error) {
	deleteOne, err := o.collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return deleteOne, nil
}
