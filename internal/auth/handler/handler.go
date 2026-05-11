package handler

import (
	"net/http"

	"github.com/PPlaner/Backend/internal/auth/service"
	"github.com/PPlaner/Backend/internal/dto"
	"github.com/PPlaner/Backend/internal/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService *service.AuthService
}

func NewHandler(authService *service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	accessToken, refreshToken, err := h.authService.Register(req.Email, req.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			response.Error(c, http.StatusConflict, "user already exists")
			return
		}

		response.Error(c, http.StatusInternalServerError, "failed to register")
		return
	}

	setRefreshCookie(c, refreshToken)

	response.Success(c, http.StatusCreated, dto.AuthResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	accessToken, refreshToken, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			response.Error(c, http.StatusUnauthorized, err.Error())
			return
		}

		response.Error(c, http.StatusInternalServerError, "failed to login")
		return
	}

	setRefreshCookie(c, refreshToken)

	response.Success(c, http.StatusOK, dto.AuthResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, _ := c.Cookie(refreshCookieName)

	accessToken, newRefreshToken, err := h.authService.Refresh(refreshToken)
	if err != nil {
		clearRefreshCookie(c)

		response.Error(c, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	setRefreshCookie(c, newRefreshToken)

	response.Success(c, http.StatusOK, dto.AuthResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	refreshToken, _ := c.Cookie(refreshCookieName)

	_ = h.authService.Logout(refreshToken)

	clearRefreshCookie(c)

	response.Success(c, http.StatusOK, dto.MessageResponse{
		Message: "logged out successfully",
	})
}
