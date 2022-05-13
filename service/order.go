package service

import (
	"fmt"
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderRepository interface {
	CreateAOrder(ctx *gin.Context, order *model.Order) (interface{}, error)
	GetUsersOrders(ctx *gin.Context, filter bson.D) (model.Order, error)
	AddNewOrder(ctx *gin.Context, filter bson.D, update bson.D) (interface{}, error)
}
type OrderService struct {
	OrderRepo OrderRepository
}

func (o *OrderService) CreateOneOrder(ctx *gin.Context) (interface{}, error) {
	order := new(model.Order)
	err := ctx.ShouldBindJSON(order)
	if err != nil {
		return nil, err
	}
	fmt.Println(order)
	createAOrder, err := o.OrderRepo.CreateAOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return createAOrder, nil

}
func (o *OrderService) GetUsersOrders(ctx *gin.Context) (interface{}, error) {
	userId := ctx.Param("id")
	filter := bson.D{{"customerid", userId}}
	orders, err := o.OrderRepo.GetUsersOrders(ctx, filter)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderService) AddNewOrder(ctx *gin.Context) (interface{}, error) {
	/*
		get userid from param
		get model
		parse model to json
		add filter
		add update bson

	*/
	userId := ctx.Param("id")
	orderModel := new(model.Product)
	ctx.ShouldBindJSON(&orderModel)
	filter := bson.D{{"customerid", userId}}
	update := bson.D{{Key: "$push", Value: bson.D{{"product", orderModel}}}}
	addNewOrder, err := o.OrderRepo.AddNewOrder(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return addNewOrder, nil
}
