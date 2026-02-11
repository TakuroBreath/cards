package main

import (
	"fmt"
	"log"

	"github.com/TakuroBreath/cards/internal/handlers"
	"github.com/TakuroBreath/cards/internal/models"
	"github.com/TakuroBreath/cards/internal/repository"
	"github.com/TakuroBreath/cards/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Database init failed: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	cardRepo := repository.NewCardRepository(db)

	userService := service.NewUserService(userRepo)
	cardService := service.NewCardService(cardRepo, userRepo)

	userHandler := handlers.NewUserHandler(userService)
	cardHandlers := handlers.NewCardHandler(cardService)

	r := gin.Default()
	userHandler.RegisterRoutes(r)
	cardHandlers.RegisterRoutes(r)

	r.StaticFile("/", "./web/index.html")

	log.Printf("Server starting on: :3000")
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func InitDB() (*gorm.DB, error) {
	// docker run -d --name cards -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=cards -p 5432:5432 postgres:16-alpine
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=cards sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Card{}); err != nil {
		return nil, fmt.Errorf("failed to migrate: %w", err)
	}

	return db, nil
}
