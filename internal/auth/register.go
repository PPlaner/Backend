package auth

import (
	"errors"

	"github.com/PPlaner/Backend/internal/models"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func (s *AuthService) Register(email, password string) (string, string, error) {
	existingUser, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", "", err
	}

	if existingUser != nil {
		return "", "", ErrUserAlreadyExists
	}

	passwordHash, err := HashPassword(password)
	if err != nil {
		return "", "", err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return "", "", err
	}

	return s.issueTokens(user.ID)

}
