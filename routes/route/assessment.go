package route

import (
	"customer-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAssessmentRoutes(r *gin.RouterGroup) {
	r.GET("/assessments", handler.GetAssessments)
	r.GET("/assessments/:id", handler.GetAssessment)
	r.POST("/assessments", handler.CreateAssessment)
	r.PUT("/assessments/:id", handler.UpdateAssessment)
	r.DELETE("/assessments/:id", handler.DeleteAssessment)

	// detail routes
	r.GET("/assessments/:id/details", handler.GetAssessmentDetail)
	r.POST("/assessments/:id/details", handler.CreateAssessmentDetail)
	r.PUT("/assessments/:id/details/:detail_id", handler.UpdateAssessmentDetail)
	r.DELETE("/assessments/:id/details/:detail_id", handler.DeleteAssessmentDetail)
	
}
