package repository

import (
	"context"
	"database/sql"

	"github.com/PPlaner/Backend/internal/models"
)

type EmailVerificationRepository struct {
	db *sql.DB
}

func NewEmailVerificationRepository(db *sql.DB) *EmailVerificationRepository {
	return &EmailVerificationRepository{
		db: db,
	}
}

func (r *EmailVerificationRepository) Create(ctx context.Context, verification *models.EmailVerification) error {
	query := `
INSERT INTO email_verifications (email, code_hash, expires_at)
VALUES ($1, $2, NOW() + INTERVAL '10 minutes')
`

	_, err := r.db.ExecContext(
		ctx,
		query,
		verification.Email,
		verification.CodeHash,
	)

	return err
}

func (r *EmailVerificationRepository) GetLatestByEmail(ctx context.Context, email string) (*models.EmailVerification, error) {
	query := `
SELECT id, email, code_hash, expires_at, used_at, created_at
FROM email_verifications
WHERE email = $1
ORDER BY created_at DESC
LIMIT 1`

	var verification models.EmailVerification
	var usedAt sql.NullTime

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&verification.ID,
		&verification.Email,
		&verification.CodeHash,
		&verification.ExpiresAt,
		&usedAt,
		&verification.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	verification.UsedAt = usedAt

	return &verification, nil
}

func (r *EmailVerificationRepository) MarkAsUsed(ctx context.Context, id int) error {
	query := `
UPDATE email_verifications
SET used_at = NOW()
WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err

}

func (r *EmailVerificationRepository) MarkAllAsUsedByEmail(ctx context.Context, email string) error {
	query := `
UPDATE email_verifications
SET used_at = NOW()
WHERE email = $1 AND used_at IS NULL`

	_, err := r.db.ExecContext(ctx, query, email)
	return err

}
