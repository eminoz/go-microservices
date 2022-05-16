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
	DeleteOneOrder(ctx *gin.Context, filter bson.D) (interface{}, error)
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
	if usersOrders.CustomerId == "" || len(usersOrders.Product) == 0 {
		addNewOrder, err := o.OrderRepo.CreateAOrder(ctx, orderModel)

		if err != nil {
			return nil, err
		}
		return addNewOrder, nil
	} else {

		for _, incomingOrder := range orderModel.Product {
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
	//in this concurrency frontend will not wait that the procces is done
	go func(data model.Order, removeOrders *model.RemoveOneOrder) (interface{}, error) {

		for _, currentOrders := range data.Product {

			//remove edilecek ürün veri tabanında bir tane ise ürünü siler
			if currentOrders.Quantity == 1 && currentOrders.ProductName == removeOrders.ProductName || removeOrders.Quantity == currentOrders.Quantity {
				filter := bson.D{{"customerid", userID}, {Key: "product.productname", Value: removeOrders.ProductName}}
				update := bson.D{{Key: "$pull", Value: bson.D{{Key: "product", Value: bson.D{{Key: "productname", Value: removeOrders.ProductName}}}}}}
				_, err := o.OrderRepo.UpdateOneOrder(ctx, filter, update)

				if err != nil {
					return nil, err
				}

			}
			//remove edilecek ürün veri tabanında birden fazla ise ürününün quantitysini bir azaltır
			if currentOrders.Quantity > 1 && currentOrders.ProductName == removeOrders.ProductName {
				filter := bson.D{{"customerid", userID}, {Key: "product.productname", Value: currentOrders.ProductName}}
				update := bson.D{{Key: "$set", Value: bson.D{{Key: "product.$.quantity", Value: currentOrders.Quantity - removeOrders.Quantity}}}}
				o.OrderRepo.FindOrderAndUpdate(ctx, filter, update)

			}

		}
		//eğer order boş ise sipariş listesinden kullanıcı çıkarılır
		defer func(ctx *gin.Context, userID string) {
			usersOrders, _ := o.OrderRepo.GetUsersOrders(ctx, filter)
			if len(usersOrders.Product) == 0 {
				filter := bson.D{{"customerid", userID}}
				o.OrderRepo.DeleteOneOrder(ctx, filter)
			}

		}(ctx, userID)
		return nil, nil
	}(usersOrders, removeOrder)

	return "ürün çıkarıldı", nil

}

// TODO: create remove one order function
// TODO: create delete all order function
