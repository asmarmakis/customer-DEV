package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"customer-api/internal/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database connection parameters
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Create connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	// Connect to database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Assign database connection to global DB variable
	DB = database

	fmt.Println("Starting auto migration...")

	isProd, _ := strconv.ParseBool(os.Getenv("IS_PRODUCTION"))
	// Drop semua tabel yang bermasalah untuk memastikan skema bersih
	// Drop tabel dengan foreign key terlebih dahulu
	if !isProd {
		tables := []interface{}{
			&entity.Activity{},
			&entity.ActivityAttendee{},
			&entity.ActivityCheckin{},
			&entity.ActivityType{},
			&entity.AccountManager{},
			&entity.Address{},
			&entity.Assessment{},
			&entity.AssessmentDetail{},
			&entity.Contact{},
			&entity.Customer{},
			&entity.Document{},
			&entity.HistoryCustomer{},
			&entity.Event{},
			&entity.EventAttendee{},
			&entity.Group{},
			&entity.Invoice{},
			&entity.Other{},
			&entity.Payment{},
			&entity.Project{},
			&entity.Role{},
			&entity.Sosmed{},
			&entity.Status{},
			&entity.StatusReasons{},
			&entity.Structure{},
			&entity.User{},
			&entity.Stages{},
			&entity.StagesDetail{},
			&entity.Workflows{},
			&entity.WorkflowsDetail{},
			&entity.GroupConfig{},
			&entity.GroupConfigDetail{},
		}
		for _, t := range tables {
			_ = DB.Migrator().DropTable(t)
		}
	}

	// Handle production migration with data cleanup
	if isProd {
		fmt.Println("Production mode: Cleaning up invalid foreign key references...")

		// First, create AccountManager table if it doesn't exist
		if !DB.Migrator().HasTable(&entity.AccountManager{}) {
			err = DB.AutoMigrate(&entity.AccountManager{})
			if err != nil {
				log.Fatal("Failed to create AccountManager table:", err)
			}
			fmt.Println("Created AccountManager table")
		}

		// Check if account_manager_id column exists in customers table
		if DB.Migrator().HasColumn(&entity.Customer{}, "account_manager_id") {
			// Clean up invalid account_manager_id references
			DB.Exec("UPDATE customers SET account_manager_id = NULL WHERE account_manager_id IS NOT NULL AND account_manager_id NOT IN (SELECT id FROM account_managers)")
			fmt.Println("Cleaned up invalid account_manager_id references")
		}
	}

	// Auto migrate the schema - akan membuat tabel sesuai model Go
	err = DB.AutoMigrate(
		&entity.Activity{},
		&entity.ActivityAttendee{},
		&entity.ActivityCheckin{},
		&entity.ActivityType{},
		&entity.AccountManager{},
		&entity.Address{},
		&entity.Assessment{},
		&entity.AssessmentDetail{},
		&entity.Contact{},
		&entity.Customer{},
		&entity.Document{},
		&entity.HistoryCustomer{},
		&entity.Event{},
		&entity.EventAttendee{},
		&entity.Group{},
		&entity.Invoice{},
		&entity.Other{},
		&entity.Payment{},
		&entity.Project{},
		&entity.Role{},
		&entity.Sosmed{},
		&entity.Status{},
		&entity.StatusReasons{},
		&entity.Structure{},
		&entity.User{},
		&entity.Stages{},
		&entity.StagesDetail{},
		&entity.Workflows{},
		&entity.WorkflowsDetail{},
		&entity.GroupConfig{},
		&entity.GroupConfigDetail{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Insert default roles if they don't exist
	var adminRole entity.Role
	result := DB.Where("id = ?", "1").First(&adminRole)
	if result.RowsAffected == 0 {
		DB.Create(&entity.Role{ID: "1", RoleName: "Admin"})
		fmt.Println("Created Admin role with ID 1")
	}

	var userRole entity.Role
	result = DB.Where("id = ?", "2").First(&userRole)
	if result.RowsAffected == 0 {
		DB.Create(&entity.Role{ID: "2", RoleName: "User"})
		fmt.Println("Created User role with ID 2")
	}

	// Insert default account managers if they don't exist
	var defaultManager entity.AccountManager
	result = DB.Where("manager_name = ?", "Default Manager").First(&defaultManager)
	if result.RowsAffected == 0 {
		DB.Create(&entity.AccountManager{ManagerName: "Default Manager"})
		fmt.Println("Created Default Manager")
	}

	var systemManager entity.AccountManager
	result = DB.Where("manager_name = ?", "System Manager").First(&systemManager)
	if result.RowsAffected == 0 {
		DB.Create(&entity.AccountManager{ManagerName: "System Manager"})
		fmt.Println("Created System Manager")
	}

	fmt.Println("Database connected and migrated successfully!")
}
