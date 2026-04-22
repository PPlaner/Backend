package service

import (
	"time"

	"github.com/PPlaner/Backend/internal/auth/repository"
)

type AuthService struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
	jwtSecret        string
	accessTokenTTL   time.Duration
	refreshTokenTTL  time.Duration
}

func NewAuthService(
	userRepo repository.UserRepository,
	refreshTokenRepo repository.RefreshTokenRepository,
	jwtSecret string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        jwtSecret,
		accessTokenTTL:   accessTokenTTL,
		refreshTokenTTL:  refreshTokenTTL,
	}
}
