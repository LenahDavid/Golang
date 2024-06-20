package routers

import (
	"awesomeProject1/controllers"
	_ "awesomeProject1/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/signin", controllers.SignIn)
	router.GET("/users", controllers.GetUsers)
	router.GET("/user/:id", controllers.GetUser)
	router.PUT("/user/:id", controllers.UpdateUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
