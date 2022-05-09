package router

import (
	"github.com/eminoz/go-microservices/api"
	"github.com/eminoz/go-microservices/repository"
	"github.com/eminoz/go-microservices/service"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	userCollectionSetting := repository.UserCollectionSetting()
	userService := service.UserService{UserRepo: userCollectionSetting}
	controller := api.UserController{UserServices: &userService}

	router.POST("/insertoneuser", controller.InsertOneUser)
	router.GET("/getoneuser/:id", controller.GetOneUser)
	router.GET("/getallusers", controller.GetAllUser)
	router.PUT("/updateuser/:id", controller.UpdateOneUser)
	router.DELETE("/deleteuser/:id")
	return router
}
