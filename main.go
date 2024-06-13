package main

import (
	"awesomeProject1/database"
	"awesomeProject1/routers"
)

// @title Awesome API
// @version 1.0
// @description API for awesome project
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	database.Connect()
	database.Migrate()

	router := routers.SetupRouter()
	err := router.Run()
	if err != nil {
		return
	}
}
