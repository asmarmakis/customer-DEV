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

	// Set trusted proxies (tambahkan ini untuk mengatasi warning)
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"}) // localhost only
	// Atau untuk production:
	// r.SetTrustedProxies([]string{"192.168.1.0/24"}) // sesuaikan dengan network Anda

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

	// Ganti port 9000 dengan port lain yang tersedia
	r.Run(":8080") // atau :3001, :5000, dll
}
