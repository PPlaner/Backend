package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/PPlaner/Backend/internal/auth/handler"
	"github.com/PPlaner/Backend/internal/auth/middleware"
	"github.com/PPlaner/Backend/internal/auth/repository"
	"github.com/PPlaner/Backend/internal/auth/service"
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

	userRepo := repository.NewUserRepo(db)
	refreshTokenRepo := repository.NewRefreshTokenRepo(db)

	authService := service.NewAuthService(
		userRepo,
		refreshTokenRepo,
		"secret-key",
		15*time.Minute,
		7*24*time.Hour,
	)

	authHandler := handler.NewHandler(authService)
	// Група API v1 згідно зі специфікацією
	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		handler.RegisterRoutes(authGroup, authHandler)
	}

	authMiddleware := middleware.AuthMiddleware("secret-key")
	protected := v1.Group("/protected")
	protected.Use(authMiddleware)

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")

		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
		})
	})
	// 4. Запуск сервера
	port := ":8080"
	fmt.Printf("Starting server on %s\n", port)
	if err := r.Run(port); err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}
