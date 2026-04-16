package auth

import (
	"net/http"

	"github.com/PPlaner/Backend/internal/dto"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *AuthService
}

func NewHandler(authService *AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: "invalid request body",
		})
		return
	}

	accessToken, refreshToken, err := h.authService.Register(req.Email, req.Password)
	if err != nil {
		if err == ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, dto.MessageResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.MessageResponse{
			Message: "failed to register",
		})
		return
	}

	setRefreshCookie(c, refreshToken)

	c.JSON(http.StatusCreated, dto.AuthResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.MessageResponse{
			Message: "invalid request body",
		})
		return
	}

	accessToken, refreshToken, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		if err == ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, dto.MessageResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dto.MessageResponse{
			Message: "failed to login",
		})
		return
	}

	setRefreshCookie(c, refreshToken)

	c.JSON(http.StatusOK, dto.AuthResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, _ := c.Cookie(refreshCookieName)

	accessToken, newRefreshToken, err := h.authService.Refresh(refreshToken)
	if err != nil {
		clearRefreshCookie(c)

		c.JSON(http.StatusUnauthorized, dto.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	setRefreshCookie(c, newRefreshToken)

	c.JSON(http.StatusOK, dto.AuthResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	refreshToken, _ := c.Cookie(refreshCookieName)

	_ = h.authService.Logout(refreshToken)

	clearRefreshCookie(c)

	c.JSON(http.StatusOK, dto.MessageResponse{
		Message: "logged out successfully",
	})
}
