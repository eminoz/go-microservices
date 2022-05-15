package service

import (
	"fmt"
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	CreateAOrder(ctx *gin.Context, order *model.Order) (interface{}, error)
	GetUsersOrders(ctx *gin.Context, filter bson.D) (model.Order, error)
	FindOrderAndUpdate(ctx *gin.Context, filter bson.D, update bson.D) interface{}
	UpdateOneOrder(ctx *gin.Context, filter bson.D, update bson.D) (interface{}, error)
	Find(ctx *gin.Context, filter bson.D) (*mongo.Cursor, error)
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

	userId := ctx.Param("id")
	orderModel := new(model.Order)
	ctx.ShouldBindJSON(&orderModel)
	filter := bson.D{{"customerid", userId}}

	usersOrders, _ := o.OrderRepo.GetUsersOrders(ctx, filter)
	if usersOrders.CustomerId == "" {
		addNewOrder, err := o.OrderRepo.CreateAOrder(ctx, orderModel)

		if err != nil {
			return nil, err
		}
		return addNewOrder, nil
	} else {

		for _, incomingOrder := range orderModel.Product {
			fmt.Println(incomingOrder.ProductName)
			for _, currentOrder := range usersOrders.Product {

				if incomingOrder.ProductName == currentOrder.ProductName {
					filter := bson.D{{Key: "customerid", Value: userId}, {Key: "product.productname", Value: currentOrder.ProductName}}
					update := bson.D{{Key: "$set", Value: bson.D{{"product.$.quantity", currentOrder.Quantity + incomingOrder.Quantity}}}}
					o.OrderRepo.FindOrderAndUpdate(ctx, filter, update)
				}

			}
		}
		/*
			filter := bson.D{{Key: "customerid", Value: userId}}
			update := bson.D{{Key: "$push", Value: bson.D{{"product", orderModel.Product}}}}
			_, err := o.OrderRepo.AddNewOrder(ctx, filter, update)
			if err != nil {
				return nil, err
			}
		*/
		return nil, nil

	}

}
func (o *OrderService) RemoveOneOrder(ctx *gin.Context) (interface{}, error) {
	userID := ctx.Param("id")
	removeOrder := new(model.RemoveOneOrder)
	err := ctx.ShouldBindJSON(removeOrder)
	if err != nil {
		return err, nil
	}
	filter := bson.D{{"customerid", userID}}
	usersOrders, _ := o.OrderRepo.GetUsersOrders(ctx, filter)
	for _, currentOrders := range usersOrders.Product {
		//remove edilecek ürün veri tabanında bir tane ise ürünü siler
		if currentOrders.Quantity == 1 && currentOrders.ProductName == removeOrder.ProductName {
			filter := bson.D{{"customerid", userID}, {Key: "product.productname", Value: removeOrder.ProductName}}
			update := bson.D{{Key: "$pull", Value: bson.D{{Key: "product", Value: bson.D{{Key: "productname", Value: removeOrder.ProductName}}}}}}
			_, err := o.OrderRepo.UpdateOneOrder(ctx, filter, update)

			if err != nil {
				return nil, err
			}

		}
		//remove edilecek ürün veri tabanında birden fazla ise ürününün quantitysini bir azaltır
		if currentOrders.Quantity > 1 && currentOrders.ProductName == removeOrder.ProductName {
			filter := bson.D{{"customerid", userID}, {Key: "product.productname", Value: currentOrders.ProductName}}
			update := bson.D{{Key: "$set", Value: bson.D{{Key: "product.$.quantity", Value: currentOrders.Quantity - removeOrder.Quantity}}}}
			o.OrderRepo.FindOrderAndUpdate(ctx, filter, update)

		} else {
			return "nothing to remove", nil
		}
	}
	return nil, nil

}

// TODO: create remove one order function
// TODO: create delete all order function
