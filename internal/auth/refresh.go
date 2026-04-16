package auth

import (
	"errors"
	"time"
)

var ErrInvalidRefreshToken = errors.New("invalid refresh token")
var ErrExpiredRefreshToken = errors.New("expired refresh token")
var ErrRevokedRefreshToken = errors.New("revoked refresh token")

func (s *AuthService) Refresh(refreshToken string) (string, string, error) {
	if refreshToken == "" {
		return "", "", ErrInvalidRefreshToken
	}

	tokenHash := HashToken(refreshToken)

	storedToken, err := s.refreshTokenRepo.GetByTokenHash(tokenHash)
	if err != nil {
		return "", "", err
	}

	if storedToken == nil {
		return "", "", ErrRevokedRefreshToken
	}

	if storedToken.RevokedAt != nil {
		return "", "", ErrRevokedRefreshToken
	}

	if time.Now().After(storedToken.ExpiresAt) {
		return "", "", ErrExpiredRefreshToken
	}

	err = s.refreshTokenRepo.RevokeByTokenHash(tokenHash)
	if err != nil {
		return "", "", err
	}

	return s.issueTokens(storedToken.UserID)

}
