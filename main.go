package main

import (
	"ewallet/config"
	"ewallet/internal/routes"
	"ewallet/pkg/connection"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load the configuration (from environment variables or a config file)
	cfg := config.LoadConfig()

	app := fiber.New()

	// Call NewDB to establish a database connection
	conn, err := connection.NewDB(
		cfg.DBDriver,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Test the connection (optional, but useful for troubleshooting)
	sqlDB, err := conn.DB().DB()
	if err != nil {
		log.Fatalf("Error getting raw database connection: %v", err)
	}
	defer sqlDB.Close()

	// You can now interact with the database using `db` (the GORM instance)
	fmt.Println("Database connected successfully!")

	// routes
	routes.RegisterRoutes(app)
	routes.TopupRoutes(app)

	// Start server
	app.Listen(":" + cfg.AppPort)
}
