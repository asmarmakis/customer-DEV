package route

import (
	"customer-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterOtherConfig(r *gin.RouterGroup) {
	r.POST("/other-configs", handler.CreateConfigGroup)
	r.GET("/other-configs", handler.GetConfigGroups)
	r.GET("/other-configs/:id", handler.GetConfigGroup)
	r.PUT("/other-configs/:id", handler.UpdateConfigGroup)
	r.DELETE("/other-configs/:id", handler.DeleteConfigGroup)
	// detail group-configs
	r.GET("/other-configs/:id/details", handler.GetConfigGroupDetails)    // Amb
	r.POST("/other-configs/:id/details", handler.CreateConfigGroupDetail) // Buat detail group-config baru
	r.GET("/other-configs/:id/details/:detail_id", handler.GetConfigGroupDetail)
	r.PUT("/other-configs/:id/details/:detail_id", handler.UpdateConfigGroupDetail) // Update detail group-config
	r.DELETE("/other-configs/:id/details/:detail_id", handler.DeleteConfigGroupDetail)

}
