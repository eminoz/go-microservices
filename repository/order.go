package repository

import (
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
)

func (o *OrderCollection) CreateAOrder(ctx *gin.Context, order *model.Order) (interface{}, error) {

	insertOne, err := o.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return insertOne, nil
}
