package main

import (
	"customer-api/internal/config"
	"customer-api/routes"
	"customer-api/middleware"

	_ "customer-api/cmd/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	// Add CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Root
	r.GET("/", func(c *gin.Context) {
		c.String(200, "server sukses berjalan")
	})

	// DB
	config.ConnectDatabase()

	// Register all routes
	routes.RegisterRoutes(r)

	// Static files
	r.Static("/uploads", "./uploads")

	r.Run(":9000")
}
