package service

import (
	"fmt"
	"github.com/eminoz/go-microservices/model"
	"github.com/gin-gonic/gin"
)

type OrderRepository interface {
	CreateAOrder(ctx *gin.Context, order *model.Order) (interface{}, error)
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
