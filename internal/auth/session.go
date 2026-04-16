package auth

import (
	"time"

	"github.com/PPlaner/Backend/internal/models"
)

func (s *AuthService) issueTokens(userID int) (string, string, error) {
	accessToken, err := GenerateAccessToken(userID, s.jwtSecret, s.accessTokenTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	tokenHash := HashToken(refreshToken)

	session := &models.RefreshToken{
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(s.refreshTokenTTL),
	}

	err = s.refreshTokenRepo.Create(session)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
