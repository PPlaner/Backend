package main

import (
	"database/sql"
	"fmt"
	"net/http"

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
		fmt.Printf("Warning: Database not connected: %v\n", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

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

	// Група API v1 згідно зі специфікацією
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
			})
			auth.POST("/login", func(c *gin.Context) {
				c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
			})
		}
	}

	// 4. Запуск сервера
	port := ":8080"
	fmt.Printf("Starting server on %s\n", port)
	if err := r.Run(port); err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}
