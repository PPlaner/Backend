package auth

import (
	"database/sql"
	"time"

	"github.com/PPlaner/Backend/internal/models"
)

type RefreshTokenRepo struct {
	db *sql.DB
}

func NewRefreshTokenRepo(db *sql.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{db: db}
}

func (r *RefreshTokenRepo) Create(token *models.RefreshToken) error {
	query := `
		INSERT INTO refresh_tokens (
			user_id,
			token_hash,
			expires_at,
			created_at,
			revoked_at
		)
		VALUES ($1, $2, $3, NOW(), NULL)
		RETURNING id, created_at, revoked_at
	`

	return r.db.QueryRow(
		query,
		token.UserID,
		token.TokenHash,
		token.ExpiresAt,
	).Scan(
		&token.ID,
		&token.CreatedAt,
		&token.RevokedAt,
	)
}

func (r *RefreshTokenRepo) GetByTokenHash(hash string) (*models.RefreshToken, error) {
	query := `
		SELECT
			id,
			user_id,
			token_hash,
			expires_at,
			created_at,
			revoked_at
		FROM refresh_tokens
		WHERE token_hash = $1
	`

	var token models.RefreshToken

	err := r.db.QueryRow(query, hash).Scan(
		&token.ID,
		&token.UserID,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.CreatedAt,
		&token.RevokedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *RefreshTokenRepo) RevokeByTokenHash(hash string) error {
	query := `
		UPDATE refresh_tokens
		SET revoked_at = $1
		WHERE token_hash = $2
	`

	_, err := r.db.Exec(query, time.Now(), hash)
	return err
}

func (r *RefreshTokenRepo) Update(token *models.RefreshToken) error {
	query := `
		UPDATE refresh_tokens
		SET
			token_hash = $1,
			expires_at = $2,
			revoked_at = $3
		WHERE id = $4
	`

	_, err := r.db.Exec(
		query,
		token.TokenHash,
		token.ExpiresAt,
		token.RevokedAt,
		token.ID,
	)

	return err
}
