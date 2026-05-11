package dto

type VerifyEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ConfirmRegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}
