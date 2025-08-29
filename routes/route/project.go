package route

import (
	"customer-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProjectRoutes(r *gin.RouterGroup) {
	r.POST("/projects", handler.CreateProject)
	r.GET("/projects", handler.ReadProjects)
	r.GET("/projects/:id", handler.ReadOneProject)
	r.PUT("/projects/:id", handler.UpdateProject)
	r.DELETE("/projects/:id", handler.DeleteProject)
}
