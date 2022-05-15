package router

import (
	"github.com/eminoz/go-microservices/api"
	"github.com/eminoz/go-microservices/pkg/middleware"
	"github.com/eminoz/go-microservices/redisContoller"
	"github.com/eminoz/go-microservices/repository"
	"github.com/eminoz/go-microservices/service"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()
	redisClient := redisContoller.RedisClient()
	userCollectionSetting := repository.UserCollectionSetting()
	userService := service.UserService{UserRepo: userCollectionSetting, UserRedisRepo: redisClient}
	controller := api.UserController{UserServices: &userService}

	corsMiddleware, isAuth := middleware.CORSMiddleware(), middleware.IsAuth()
	router.Use(corsMiddleware)
	router.POST("/insertoneuser", controller.InsertOneUser)
	router.POST("/login", controller.Login)
	router.GET("/getoneuser/:id", isAuth, controller.GetOneUser)
	router.GET("/getallusers", isAuth, controller.GetAllUser)
	router.PUT("/updateuser/:id", isAuth, controller.UpdateOneUser)
	router.DELETE("/deleteuser/:id", isAuth, controller.DeleteOneUser)

	orderCollectionSetting := repository.OrderCollectionSetting()
	orderService := service.OrderService{OrderRepo: orderCollectionSetting}
	orderController := api.OrderController{OrderService: &orderService}

	router.POST("/createorder", isAuth, orderController.CreateOrder)
	router.GET("/getorders/:id", orderController.GetUserOrders)
	router.POST("/addneworder/:id", orderController.AddNewOrder)
	router.POST("/removeorder/:id", orderController.RemoveOneOrder)
	return router
}
