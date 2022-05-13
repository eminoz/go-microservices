package repository

import (
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

func (o *OrderCollection) AddNewOrder(ctx *gin.Context, filter bson.D, update bson.D) (interface{}, error) {
	updateOne, err := o.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}
