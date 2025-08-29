package handler

import (
	"customer-api/internal/config"
	"customer-api/internal/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OtherConfig struct {
	ID        string         `json:"id" gorm:"primaryKey;size:26"`
	Name      string         `json:"name" gorm:"not null;unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
}

type OtherConfigDetail struct {
	ID        string         `json:"id" gorm:"primaryKey;size:26"`
	ConfigID  string         `json:"config_id" gorm:"index"`
	Icon      string         `json:"icon"`
	IsActive  bool           `json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// relationships
	Config OtherConfig `json:"config" gorm:"foreignKey:ConfigID;references:ID"`
}

func CreateConfigOther(c *gin.Context) {
	var input entity.GroupConfig
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	if result := config.DB.Create(&input); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Gagal membuat config group",
			"data":    nil,
		})
		return
	}

	// insert group config

	groupConfig := entity.GroupConfig{
		Name: input.Name,
	}
	if result := config.DB.Create(&groupConfig); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Gagal membuat group config",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Config group berhasil dibuat",
		"data":    groupConfig,
	})
}

func GetConfigOther(c *gin.Context) {
	var groups []entity.GroupConfig
	if result := config.DB.Find(&groups); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data config groups"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data config groups berhasil diambil",
		"data":    groups,
	})
}

func GetConfigOtherDetail(c *gin.Context) {
	id := c.Param("id")
	var group entity.GroupConfig
	if result := config.DB.Where("id = ?", id).First(&group); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Config group not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Config group fetched successfully",
		"data":    group,
	})
}

func UpdateConfigOther(c *gin.Context) {
	id := c.Param("id")
	var input entity.GroupConfig
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	var group entity.GroupConfig
	if result := config.DB.Where("id = ?", id).First(&group); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Config group not found"})
		return
	}

	if input.Name != "" {
		group.Name = input.Name
	}

	if result := config.DB.Save(&group); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Gagal memperbarui config group",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Config group updated successfully",
		"data":    group,
	})
}

func DeleteConfigOther(c *gin.Context) {
	id := c.Param("id")
	var group entity.GroupConfig
	if result := config.DB.Where("id = ?", id).First(&group); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Config group not found"})
		return
	}

	if result := config.DB.Delete(&group); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Gagal menghapus config group",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Config group deleted successfully",
		"data":    nil,
	})
}

// CreateConfigOtherDetail function has been removed as it's handled by CreateConfigGroupDetail in groupConfig.go

func GetConfigOtherDetails(c *gin.Context) {
	var details []entity.GroupConfigDetail
	if result := config.DB.Find(&details); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data config other details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data config other details berhasil diambil",
		"data":    details,
	})
}

func GetConfigOthersDetail(c *gin.Context) {
	id := c.Param("id")
	var detail entity.OthersConfigDetail
	if result := config.DB.Where("id = ?", id).First(&detail); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Config other detail not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Config other detail fetched successfully",
		"data":    detail,
	})
}

func UpdateConfigOtherDetail(c *gin.Context) {
	id := c.Param("id")
	var input entity.OthersConfigDetail
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	var detail entity.GroupConfigDetail
	if result := config.DB.Where("id = ?", id).First(&detail); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Config other detail not found"})
		return
	}
	if input.ConfigID != "" {
		detail.Name = input.ConfigID
	}
	if input.IsActive != detail.IsActive {
		detail.IsActive = input.IsActive
	}
	if result := config.DB.Save(&detail); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Gagal memperbarui config other detail",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Config other detail updated successfully",
		"data":    detail,
	})
}

func DeleteConfigOtherDetail(c *gin.Context) {
	id := c.Param("id")
	var detail entity.OthersConfigDetail
	if result := config.DB.Where("id = ?", id).First(&detail); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Config other detail not found"})
		return
	}
	if result := config.DB.Delete(&detail); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "Gagal menghapus config other detail",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Config other detail deleted successfully",
		"data":    nil,
	})
}
