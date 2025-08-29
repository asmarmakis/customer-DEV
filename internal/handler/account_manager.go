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

	// Convert to simplified response DTO (excluding unnecessary fields)
	response := dto.AccountManagerListResponse{
		ID:          accountManager.ID,
		ManagerName: accountManager.ManagerName,
		CreatedAt:   accountManager.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   accountManager.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
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

	// Convert to simplified response DTOs (excluding unnecessary fields)
	var responses []dto.AccountManagerListResponse
	for _, am := range accountManagers {
		responses = append(responses, dto.AccountManagerListResponse{
			ID:          am.ID,
			ManagerName: am.ManagerName,
			CreatedAt:   am.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   am.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	c.JSON(http.StatusOK, responses)
}

// GetAccountManager gets a specific account manager by ID
func GetAccountManager(c *gin.Context) {
	id := c.Param("id")
	var accountManager entity.AccountManager

	if result := config.DB.Where("id = ?", id).First(&accountManager); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account manager tidak ditemukan"})
		return
	}

	// Convert to response DTO
	response := dto.AccountManagerResponse{
		ID:          accountManager.ID,
		ManagerName: accountManager.ManagerName,

		CreatedAt: accountManager.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: accountManager.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	c.JSON(http.StatusOK, response)
}

// UpdateAccountManager updates an existing account manager
func UpdateAccountManager(c *gin.Context) {
	id := c.Param("id")
	var accountManager entity.AccountManager

	if result := config.DB.Where("id = ?", id).First(&accountManager); result.Error != nil {
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
		if result := config.DB.Where("manager_name = ? AND id != ?", *input.ManagerName, id).First(&existingManager); result.Error == nil {
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

		CreatedAt: accountManager.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: accountManager.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account manager berhasil diupdate",
		"data":    response,
	})
}

// DeleteAccountManager deletes an account manager
func DeleteAccountManager(c *gin.Context) {
	id := c.Param("id")
	var accountManager entity.AccountManager

	if result := config.DB.Where("id = ?", id).First(&accountManager); result.Error != nil {
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

// @Summary Get account managers for dropdown
// @Description Get simplified list of account managers for dropdown selection (ID and manager_name only)
// @Tags AccountManager
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.AccountManagerDetail
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/account-managers/dropdown [get]
func GetAccountManagersDropdown(c *gin.Context) {
	var accountManagers []entity.AccountManager
	if result := config.DB.Select("id, manager_name").Find(&accountManagers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data account managers"})
		return
	}

	// Convert to simplified response for dropdown
	var dropdownOptions []dto.AccountManagerDetail
	for _, am := range accountManagers {
		dropdownOptions = append(dropdownOptions, dto.AccountManagerDetail{
			ID:          am.ID,
			ManagerName: am.ManagerName,
		})
	}

	c.JSON(http.StatusOK, dropdownOptions)
}
