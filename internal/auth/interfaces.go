package auth

import "github.com/PPlaner/Backend/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
}

type RefreshTokenRepository interface {
	Create(token *models.RefreshToken) error
	GetByTokenHash(hash string) (*models.RefreshToken, error)
	RevokeByTokenHash(hash string) error
	Update(token *models.RefreshToken) error
}
