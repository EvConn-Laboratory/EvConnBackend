package main

import (
	"evconn/internal/core/domain/models"
	"evconn/internal/infrastructure/database"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func createAdmin(db *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &models.User{
		NIM:      "admin",
		Name:     "System Admin",
		Email:    "admin@evconn.com",
		Password: string(hashedPassword),
		Role:     "admin",
		Lab:      "ALL",
	}

	return db.Create(admin).Error
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize DB with explicit configuration
	db, err := database.NewMySQLConnection(&database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Course{},
		&models.Module{},
		&models.Assignment{},
		&models.Submission{},
		&models.ModuleMentor{},
		&models.Feedback{},
		&models.Lab{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create admin user after migration
	if err := createAdmin(db); err != nil {
		log.Printf("Note: Admin user already exists or error: %v", err)
	}

	log.Println("Database migration completed successfully")
}
