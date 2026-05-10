package handler

import (
	"net/http"

	"github.com/PPlaner/Backend/internal/auth/service"
	"github.com/PPlaner/Backend/internal/dto"
	"github.com/PPlaner/Backend/internal/email"
	"github.com/gin-gonic/gin"
)

type EmailVerificationHandler struct {
	service     *service.EmailVerificationService
	authService *service.AuthService
	emailSender *email.Sender
}

func NewEmailVerificationHandler(service *service.EmailVerificationService, authService *service.AuthService, emailSender *email.Sender) *EmailVerificationHandler {
	return &EmailVerificationHandler{
		service:     service,
		authService: authService,
		emailSender: emailSender,
	}

}

func (h *EmailVerificationHandler) VerifyEmail(c *gin.Context) {
	var req dto.VerifyEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	code, err := h.service.CreateVerification(c.Request.Context(), req.Email)
	err = h.emailSender.SendVerificationCode(req.Email, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to send verification email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "verification code sent",
	})
}

func (h *EmailVerificationHandler) ConfirmRegister(c *gin.Context) {
	var req dto.ConfirmRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.service.ConfirmCode(c.Request.Context(), req.Email, req.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accessToken, refreshToken, err := h.authService.Register(
		req.Email,
		req.Password,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	setRefreshCookie(c, refreshToken)

	c.JSON(http.StatusOK, gin.H{
		"message":      "registration confirmed",
		"access_token": accessToken,
	})
}
