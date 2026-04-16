package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRoutes, handler *Handler) {
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.POST("/refresh", handler.Refresh)
	router.POST("/logout", handler.Logout)
}
