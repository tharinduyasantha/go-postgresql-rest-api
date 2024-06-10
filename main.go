package main

import (
	"github.com/tharinduyasantha/go-postgresql-rest-api/config"
	"github.com/tharinduyasantha/go-postgresql-rest-api/models"
	"github.com/tharinduyasantha/go-postgresql-rest-api/routers"
)

// @title Go REST API
// @version 1.0
// @description This is a sample server for a REST API in Go.
// @host localhost:8080
// @BasePath /
func main() {
	config.Connect()
	config.DB.AutoMigrate(&models.User{})
	r := routers.SetupRouter()
	// Swagger documentation
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
