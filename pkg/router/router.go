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
	corsMiddleware := middleware.CORSMiddleware()
	router.POST("/insertoneuser", corsMiddleware, controller.InsertOneUser)
	router.GET("/getoneuser/:id", controller.GetOneUser)
	router.GET("/getallusers", controller.GetAllUser)
	router.PUT("/updateuser/:id", controller.UpdateOneUser)
	router.DELETE("/deleteuser/:id", controller.DeleteOneUser)
	return router
}
