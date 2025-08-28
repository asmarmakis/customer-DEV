package handler

import (
	"net/http"

	"customer-api/internal/config"
	"customer-api/internal/dto"
	"customer-api/internal/entity"

	"github.com/gin-gonic/gin"
)

// CreateAccountManager creates a new account manager
func CreateAccountManager(c *gin.Context) {
	var input dto.CreateAccountManagerRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if manager name already exists
	var existingManager entity.AccountManager
	if result := config.DB.Where("manager_name = ?", input.ManagerName).First(&existingManager); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Manager name sudah digunakan"})
		return
	}

	accountManager := entity.AccountManager{
		ManagerName: input.ManagerName,
	}

	if result := config.DB.Create(&accountManager); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat account manager"})
		return
	}

	// Convert to response DTO
	response := dto.AccountManagerResponse{
		ID:          accountManager.ID,
		ManagerName: accountManager.ManagerName,
		CreatedAt:   accountManager.CreatedAt,
		UpdatedAt:   accountManager.UpdatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account manager berhasil dibuat",
		"data":    response,
	})
}

// GetAccountManagers gets all account managers
func GetAccountManagers(c *gin.Context) {
	var accountManagers []entity.AccountManager
	if result := config.DB.Find(&accountManagers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data account managers"})
		return
	}

	// Convert to response DTOs
	var responses []dto.AccountManagerResponse
	for _, am := range accountManagers {
		responses = append(responses, dto.AccountManagerResponse{
			ID:          am.ID,
			ManagerName: am.ManagerName,
			CreatedAt:   am.CreatedAt,
			UpdatedAt:   am.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, responses)
}

// GetAccountManager gets a specific account manager by UUID
func GetAccountManager(c *gin.Context) {
	uuid := c.Param("id")
	var accountManager entity.AccountManager

	if result := config.DB.Where("uuid = ?", uuid).First(&accountManager); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account manager tidak ditemukan"})
		return
	}

	// Convert to response DTO
	response := dto.AccountManagerResponse{
		ID:          accountManager.ID,
		ManagerName: accountManager.ManagerName,
		CreatedAt:   accountManager.CreatedAt,
		UpdatedAt:   accountManager.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateAccountManager updates an existing account manager
func UpdateAccountManager(c *gin.Context) {
	uuid := c.Param("id")
	var accountManager entity.AccountManager

	if result := config.DB.Where("uuid = ?", uuid).First(&accountManager); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account manager tidak ditemukan"})
		return
	}

	var input dto.UpdateAccountManagerRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if manager name already exists (exclude current manager)
	if input.ManagerName != nil {
		var existingManager entity.AccountManager
		if result := config.DB.Where("manager_name = ? AND uuid != ?", *input.ManagerName, uuid).First(&existingManager); result.Error == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Manager name sudah digunakan"})
			return
		}
		accountManager.ManagerName = *input.ManagerName
	}

	if result := config.DB.Save(&accountManager); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate account manager"})
		return
	}

	// Convert to response DTO
	response := dto.AccountManagerResponse{
		ID:          accountManager.ID,
		ManagerName: accountManager.ManagerName,
		CreatedAt:   accountManager.CreatedAt,
		UpdatedAt:   accountManager.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account manager berhasil diupdate",
		"data":    response,
	})
}

// DeleteAccountManager deletes an account manager
func DeleteAccountManager(c *gin.Context) {
	uuid := c.Param("id")
	var accountManager entity.AccountManager

	if result := config.DB.Where("uuid = ?", uuid).First(&accountManager); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account manager tidak ditemukan"})
		return
	}

	// Check if account manager is being used by customers
	var customerCount int64
	config.DB.Model(&entity.Customer{}).Where("account_manager_id = ?", accountManager.ID).Count(&customerCount)
	if customerCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account manager tidak dapat dihapus karena masih digunakan oleh customer"})
		return
	}

	if result := config.DB.Delete(&accountManager); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus account manager"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account manager berhasil dihapus"})
}
