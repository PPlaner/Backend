package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/PPlaner/Backend/internal/auth"
	"github.com/PPlaner/Backend/internal/config"
	"github.com/PPlaner/Backend/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Backend is running")

	// 1. Завантаження конфігурації
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// 2. Підключення до БД
	db, err := database.Connect(cfg.DB)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	defer func(db *sql.DB) {
		if db != nil {
			_ = db.Close()
		}
	}(db)

	fmt.Println("Connected to database")

	// 3. Ініціалізація HTTP-сервера
	r := gin.Default()

	// Базовий ендпоінт для перевірки статусу
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "active",
			"service": "PPlaner API",
		})
	})

	userRepo := auth.NewUserRepo(db)
	refreshTokenRepo := auth.NewRefreshTokenRepo(db)

	authService := auth.NewAuthService(
		userRepo,
		refreshTokenRepo,
		"secret-key",
		15*time.Minute,
		7*24*time.Hour,
	)

	authHandler := auth.NewHandler(authService)
	// Група API v1 згідно зі специфікацією
	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		auth.RegisterRoutes(authGroup, authHandler)
	}

	// 4. Запуск сервера
	port := ":8080"
	fmt.Printf("Starting server on %s\n", port)
	if err := r.Run(port); err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}
