package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/PPlaner/Backend/internal/auth/repository"
	"github.com/PPlaner/Backend/internal/models"
)

type EmailVerificationService struct {
	repo *repository.EmailVerificationRepository
}

func NewEmailVerificationService(repo *repository.EmailVerificationRepository) *EmailVerificationService {
	return &EmailVerificationService{repo: repo}
}

func generateVerificationCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func hashVerificationCode(code string) string {
	hash := sha256.Sum256([]byte(code))
	return hex.EncodeToString(hash[:])
}

func (s *EmailVerificationService) CreateVerification(ctx context.Context, email string) (string, error) {
	err := s.repo.MarkAllAsUsedByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	code := generateVerificationCode()
	codeHash := hashVerificationCode(code)

	verification := &models.EmailVerification{
		Email:    email,
		CodeHash: codeHash,
	}

	err = s.repo.Create(ctx, verification)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (s *EmailVerificationService) ConfirmCode(ctx context.Context, email string, code string) error {
	verification, err := s.repo.GetLatestByEmail(ctx, email)
	if err != nil {
		return err
	}

	if verification.UsedAt.Valid {
		return fmt.Errorf("verification code already used")
	}

	if time.Now().After(verification.ExpiresAt) {
		return fmt.Errorf("verification code expired")
	}

	codeHash := hashVerificationCode(code)
	if verification.CodeHash != codeHash {
		return fmt.Errorf("invalid verification code")
	}

	err = s.repo.MarkAsUsed(ctx, verification.ID)
	if err != nil {
		return err
	}

	return s.repo.MarkAsUsed(ctx, verification.ID)
}
