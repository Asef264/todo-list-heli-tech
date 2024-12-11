package main

import (
	"database/sql"
	"log"
	"todo-list/api/router"
	"todo-list/config"
	"todo-list/internal/adapters/persistence"

	"github.com/gin-gonic/gin"
)

func main() {
	app, cfg, db := Init()

	// Ensure the database connection is closed when the server stops
	defer db.Close()
	// Start the server
	log.Printf("Server is running on port %d\n", cfg.Server.Port)
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func Init() (*gin.Engine, *config.Config, *sql.DB) {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Set up database connection
	db, err := persistence.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Database connection established successfully!")

	// Run database migrations
	persistence.MigrateUp(db, "migrations", cfg.DB.DBName)

	// Initialize Gin router
	app := gin.Default()

	// Register routes
	router.RegisterRoutes(app, db, cfg)

	return app, cfg, db
}
