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

	// @Summary Create a new user
	// @Description Create a new user
	// @Tags users
	// @Accept  json
	// @Produce  json
	// @Success 201 {object} User
	// @Router /signup [post]
	router.POST("/signup", controllers.SignUp)

	// @Summary Sign in a user
	// @Description Sign in a user
	// @Tags users
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} User
	// @Router /signin [post]
	router.POST("/signin", controllers.SignIn)

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
