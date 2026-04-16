package auth

import (
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *AuthService) Login(email, password string) (string, string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", ErrInvalidCredentials
	}

	err = CheckPasswordHash(password, user.PasswordHash)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}
	return s.issueTokens(user.ID)
}
