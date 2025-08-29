package handler

import (
	"net/http"
	"strconv"

	"customer-api/internal/config"
	"customer-api/internal/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"customer-api/internal/dto"
)

// @Summary Get all Projects
// @Description Get list of all projects
// @Tags Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Success 200 {array} entity.Project
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/projects [get]
func ReadProjects(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	if limit == 0 {
		limit = 10 // Default limit
	}
	if page == 0 {
		page = 1 // Default page
	}

	var projects []entity.Project
	offset := (page - 1) * limit

	if result := config.DB.Limit(limit).Offset(offset).Find(&projects); result.Error != nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Status:  http.StatusNotFound,
			Message: "No projects found",
			Data:    []entity.Project{},
		})
		return
	}

	if len(projects) == 0 {
		c.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "No projects found",
			Data:    projects,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Projects retrieved successfully",
		Data:    projects,
	})
}

// @Summary Create a new Project
// @Description Create a new project
// @Tags Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param project body entity.Project true "Project"
// @Success 201 {object} entity.Project
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/projects [post]
func CreateProject(c *gin.Context) {
	var project entity.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

// @Summary Get a Project by ID
// @Description Get a project by ID
// @Tags Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Project ID"
// @Success 200 {object} entity.Project
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/projects/{id} [get]
func ReadOneProject(c *gin.Context) {
	var project entity.Project
	id := c.Param("id")
	if result := config.DB.Where("id = ?", id).First(&project); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data project"})
		}
		return
	}
	c.JSON(http.StatusOK, project)
}

// @Summary Update a Project by ID
// @Description Update a project by ID
// @Tags Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Project ID"
// @Param project body entity.Project true "Project"
// @Success 200 {object} entity.Project
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/projects/{id} [put]
func UpdateProject(c *gin.Context) {
	var project entity.Project
	id := c.Param("id")

	// First, find the existing project
	if result := config.DB.Where("id = ?", id).First(&project); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data project"})
		}
		return
	}

	// Bind the new data
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated project
	if err := config.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// @Summary Delete a Project by ID
// @Description Delete a project by ID
// @Tags Projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Project ID"
// @Success 204
// @Failure 401 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/projects/{id} [delete]
func DeleteProject(c *gin.Context) {
	var project entity.Project
	id := c.Param("id")
	if result := config.DB.Where("id = ?", id).First(&project); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data project"})
		}
		return
	}

	if err := config.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data project"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
