package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tharinduyasantha/go-postgresql-rest-api/config"
	_ "github.com/tharinduyasantha/go-postgresql-rest-api/docs"
	"github.com/tharinduyasantha/go-postgresql-rest-api/models"
	"github.com/tharinduyasantha/go-postgresql-rest-api/routers"
)

// @title Go REST API
// @version 1.0
// @description This is a sample server for a REST API in Go.
// @host localhost:8081
// @BasePath /
func main() {
	config.Connect()
	config.DB.AutoMigrate(&models.User{})
	r := routers.SetupRouter()
	// CORS middleware

	// Swagger documentation
	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// r.Use(cors.Default())
	r.Run(":8081")
}
