package handler

import (
	"net/http"

	"github.com/PPlaner/Backend/internal/auth/service"
	"github.com/PPlaner/Backend/internal/dto"
	"github.com/PPlaner/Backend/internal/email"
	"github.com/PPlaner/Backend/internal/response"
	"github.com/gin-gonic/gin"
)

type EmailVerificationHandler struct {
	service     *service.EmailVerificationService
	authService *service.AuthService
	emailSender *email.Sender
}

func NewEmailVerificationHandler(
	service *service.EmailVerificationService,
	authService *service.AuthService,
	emailSender *email.Sender,
) *EmailVerificationHandler {
	return &EmailVerificationHandler{
		service:     service,
		authService: authService,
		emailSender: emailSender,
	}
}

func (h *EmailVerificationHandler) VerifyEmail(c *gin.Context) {
	var req dto.VerifyEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request body")
		return
	}

	code, err := h.service.CreateVerification(c.Request.Context(), req.Email)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to create verification code")
		return
	}

	if err := h.emailSender.SendVerificationCode(req.Email, code); err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to send verification email")
		return
	}

	response.Success(c, http.StatusOK, gin.H{
		"message": "verification code sent",
	})
}

func (h *EmailVerificationHandler) ConfirmRegister(c *gin.Context) {
	var req dto.ConfirmRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.ConfirmCode(c.Request.Context(), req.Email, req.Code); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.authService.Register(req.Email, req.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			response.Error(c, http.StatusConflict, "user already exists")
			return
		}

		response.Error(c, http.StatusInternalServerError, "internal server error")
		return
	}

	setRefreshCookie(c, refreshToken)

	response.Success(c, http.StatusOK, gin.H{
		"message":     "registration confirmed",
		"accessToken": accessToken,
	})
}
