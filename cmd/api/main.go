package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	authHandler "github.com/PPlaner/Backend/internal/auth/handler"
	authMiddleware "github.com/PPlaner/Backend/internal/auth/middleware"
	authRepository "github.com/PPlaner/Backend/internal/auth/repository"
	authService "github.com/PPlaner/Backend/internal/auth/service"

	syncHandler "github.com/PPlaner/Backend/internal/sync/handler"
	syncRepository "github.com/PPlaner/Backend/internal/sync/repository"
	syncService "github.com/PPlaner/Backend/internal/sync/service"

	keyHandler "github.com/PPlaner/Backend/internal/user/handler"
	keyRepository "github.com/PPlaner/Backend/internal/user/repository"
	keyService "github.com/PPlaner/Backend/internal/user/service"

	"github.com/PPlaner/Backend/internal/config"
	"github.com/PPlaner/Backend/internal/database"
	"github.com/PPlaner/Backend/internal/email"
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
		panic(fmt.Sprintf("failed to connect database: %v", err))
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

	userRepo := authRepository.NewUserRepo(db)
	refreshTokenRepo := authRepository.NewRefreshTokenRepo(db)

	emailVerificationRepo := authRepository.NewEmailVerificationRepository(db)
	emailVerificationSvc := authService.NewEmailVerificationService(emailVerificationRepo)
	emailSender := email.NewSender(cfg.SMTP)

	authSvc := authService.NewAuthService(
		userRepo,
		refreshTokenRepo,
		"secret-key",
		15*time.Minute,
		7*24*time.Hour,
	)

	authH := authHandler.NewHandler(authSvc)

	emailVerificationH := authHandler.NewEmailVerificationHandler(
		emailVerificationSvc,
		authSvc,
		emailSender,
	)

	keyRepo := keyRepository.NewKeyRepository(db)
	keySvc := keyService.NewKeyService(keyRepo)
	keyH := keyHandler.NewKeyHandler(keySvc)

	syncRepo := syncRepository.NewSyncRepository(db)
	syncSvc := syncService.NewSyncService(syncRepo)
	syncH := syncHandler.NewSyncHandler(syncSvc)

	// Група API v1 згідно зі специфікацією
	v1 := r.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		authHandler.RegisterRoutes(authGroup, authH)
		authGroup.POST("/verify-email", emailVerificationH.VerifyEmail)
		authGroup.POST("/confirm-register", emailVerificationH.ConfirmRegister)
	}

	mw := authMiddleware.AuthMiddleware("secret-key")

	protected := v1.Group("/protected")
	protected.Use(mw)

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("user_id")

		c.JSON(http.StatusOK, gin.H{
			"user_id": userID,
		})
	})

	protected.GET("/me/keys", keyH.GetKeys)
	protected.POST("/me/keys", keyH.SaveKeys)
	protected.PATCH("/me/keys", keyH.SaveKeys)

	protected.POST("/sync", syncH.Sync)

	// 4. Запуск сервера
	port := ":8080"
	fmt.Printf("Starting server on %s\n", port)
	if err := r.Run(port); err != nil {
		panic(fmt.Sprintf("failed to start server: %v", err))
	}
}
