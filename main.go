package main

import (
	"bookingAPI/database"
	"bookingAPI/models"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment file
	loadEnv()
	// load database configuration and connection
	loadDatabase()
	// start the server
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
	database.InitDB()
	err := database.DB.AutoMigrate(&models.Role{})
	if err != nil {
		return
	}
	err1 := database.DB.AutoMigrate(&models.User{})
	if err1 != nil {
		return
	}
	seedData()
}

func serveApplication() {
	router := gin.Default()

	// Set SecureJSON middleware to handle proxy headers securely
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Security-Policy", "default-src 'self'")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		c.Writer.Header().Set("X-Frame-Options", "DENY")
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")

		// Check X-Forwarded-Proto header for HTTPS requests
		if c.Request.Header.Get("X-Forwarded-Proto") == "https" {
			c.Request.URL.Scheme = "https"
			c.Request.URL.Host = c.Request.Header.Get("Host")
		}

		c.Next()
	})

	err := router.Run(":8000")
	if err != nil {
		return
	}
	fmt.Println("Server running on port 8000")
}

// load seed data into the database
func seedData() {
	var roles = []models.Role{
		{
			Name:        "admin",
			Description: "Administrator role",
		},
		{
			Name:        "customer",
			Description: "Authenticated customer role",
		},
		{
			Name:        "anonymous",
			Description: "Unauthenticated customer role",
		},
	}
	var user = []models.User{
		{
			Username: os.Getenv("ADMIN_USERNAME"),
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: os.Getenv("ADMIN_PASSWORD"),
			RoleID:   1,
		},
	}
	database.DB.Save(&roles)
	database.DB.Save(&user)
}
