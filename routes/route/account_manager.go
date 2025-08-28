package route

import (
	"customer-api/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAccountManagerRoutes(r *gin.RouterGroup) {
	r.POST("/account-managers", handler.CreateAccountManager)
	r.GET("/account-managers", handler.GetAccountManagers)
	r.GET("/account-managers/:id", handler.GetAccountManager)
	r.PUT("/account-managers/:id", handler.UpdateAccountManager)
	r.DELETE("/account-managers/:id", handler.DeleteAccountManager)
}