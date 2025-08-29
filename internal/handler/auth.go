package handler

import (
	"net/http"
	"os"
	"time"

	"customer-api/internal/config"
	"customer-api/internal/dto"
	"customer-api/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi bahwa salah satu dari username atau email harus diisi
	if input.Username == "" && input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username atau email harus diisi"})
		return
	}

	var user entity.User
	var usernameOrEmail string

	if input.Username != "" {
		usernameOrEmail = input.Username
	} else {
		usernameOrEmail = input.Email
	}

	// Cek apakah username/email terdaftar
	result := config.DB.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Username atau email tidak ditemukan
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau email tidak terdaftar"})
			return
		}
		// Error database lainnya
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat login"})
		return
	}

	// Cek password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}

// @Summary Register new user
// @Description Register a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration data"
// @Success 201 {object} dto.User
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		// Enhanced error handling with detailed validation information
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
			"hint":    "Please ensure all required fields are provided with correct data types. role_id can be a string or number.",
		})
		return
	}

	// Check if username already exists
	var existingUser entity.User
	if result := config.DB.Where("username = ?", input.Username).First(&existingUser); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username already exists",
			"hint":  "Please choose a different username",
		})
		return
	}

	// Check if email already exists
	if result := config.DB.Where("email = ?", input.Email).First(&existingUser); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists",
			"hint":  "Please use a different email address",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
			"hint":  "Please try again or contact support",
		})
		return
	}

	// Set default role_id if not provided or validate existing role_id
	roleID := input.RoleID

	if roleID == "" {
		// Find "User" role from database
		var userRole entity.Role
		if err := config.DB.Where("role_name = ?", "User").First(&userRole).Error; err != nil {
			// If User role not found, use the first available role
			if err := config.DB.First(&userRole).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "No roles available in the system",
					"hint":  "Please contact administrator to set up user roles",
				})
				return
			}
		}
		roleID = userRole.ID
	} else {
		// Validate that the provided role_id exists in database
		var role entity.Role
		if err := config.DB.Where("id = ?", roleID).First(&role).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid role_id",
				"hint":  "Please provide a valid role ID. Use GET /api/roles to see available roles",
			})
			return
		}
	}

	// Create new user
	user := entity.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		RoleID:   roleID,
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to register user",
			"details": result.Error.Error(),
			"hint":    "Please try again or contact support",
		})
		return
	}

	// Return success response without sensitive data
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role_id":  user.RoleID,
		},
	})
}
