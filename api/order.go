package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderServices interface {
	CreateOneOrder(ctx *gin.Context) (interface{}, error)
	GetUsersOrders(ctx *gin.Context) (interface{}, error)
	AddNewOrder(ctx *gin.Context) (interface{}, error)
	RemoveOneOrder(ctx *gin.Context) (interface{}, error)
	RemoveAllOrderByOrderID(ctx *gin.Context) (interface{}, error)
}
type OrderController struct {
	OrderService OrderServices
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	order, err := o.OrderService.CreateOneOrder(ctx)
	if err != nil {
		ctx.JSON(201, gin.H{"message": "order not created"})
	}
	ctx.JSON(http.StatusOK, gin.H{"order id": order})

}

func (o *OrderController) GetUserOrders(ctx *gin.Context) {
	orders, err := o.OrderService.GetUsersOrders(ctx)
	if err != nil {
		ctx.JSON(201, gin.H{"message": "order not created"})
	}
	ctx.JSON(http.StatusOK, gin.H{"orders": orders})

}
func (o *OrderController) AddNewOrder(ctx *gin.Context) {
	addNewOrder, err := o.OrderService.AddNewOrder(ctx)
	if err != nil {
		ctx.JSON(201, gin.H{"message": "order not created"})
	}
	ctx.JSON(http.StatusOK, gin.H{"orders": addNewOrder})
}
func (o *OrderController) RemoveOneOrder(ctx *gin.Context) {
	addNewOrder, err := o.OrderService.RemoveOneOrder(ctx)
	if err != nil {
		ctx.JSON(201, gin.H{"message": "order not updated"})
	}
	ctx.JSON(http.StatusOK, gin.H{"orders": addNewOrder})
}
func (o *OrderController) RemoveAllOrderByOrderID(ctx *gin.Context) {
	removeAllOrderByID, err := o.OrderService.RemoveAllOrderByOrderID(ctx)
	if err != nil {
		ctx.JSON(201, gin.H{"message": "order not remove"})
	}
	ctx.JSON(http.StatusOK, gin.H{"deleted order": removeAllOrderByID})
}
