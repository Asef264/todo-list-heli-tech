package main

import (
	"database/sql"
	"log"
	"os"

	"todo-list/config"
	"todo-list/internal/adapters/api/router"
	dbPkg "todo-list/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	app, cfg, db := Init()
	defer db.Close()
	log.Printf("Server is running on port %d\n", cfg.Server.Port)
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func Init() (*gin.Engine, *config.Config, *sql.DB) {
	if err := os.Setenv("STORAGE_TYPE", "s3"); err != nil {
		log.Fatalf("error on set storage type,%e", err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	config.AppConfig = cfg

	db, err := dbPkg.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Database connection established successfully!")

	dbPkg.MigrateUp(db, "migrations", cfg.DB.DBName)

	app := gin.Default()

	router.RegisterRoutes(app, db, cfg)

	return app, cfg, db
}
