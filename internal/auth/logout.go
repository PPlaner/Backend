package auth

import (
	"errors"
)

var ErrInvalidLogoutToken = errors.New("invalid token")

func (s *AuthService) Logout(refreshToken string) error {
	if refreshToken == "" {
		return ErrInvalidLogoutToken
	}

	tokenHash := HashToken(refreshToken)

	err := s.refreshTokenRepo.RevokeByTokenHash(tokenHash)
	if err != nil {
		return err
	}

	return nil
}
