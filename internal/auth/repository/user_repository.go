package repository

import (
	"database/sql"

	"github.com/PPlaner/Backend/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *models.User) error {
	query := `
		INSERT INTO users (
			email,
			password_hash,
			salt,
			wmk_pin,
			wmk_recovery,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING id, created_at, updated_at
	`

	return r.db.QueryRow(
		query,
		user.Email,
		user.PasswordHash,
		user.Salt,
		user.WmkPin,
		user.WmkRecovery,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}
func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT 
			id,
			email,
			password_hash,
			salt,
			wmk_pin,
			wmk_recovery,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	var user models.User

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
		&user.WmkPin,
		&user.WmkRecovery,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *UserRepo) GetByID(id int) (*models.User, error) {
	query := `
		SELECT 
			id,
			email,
			password_hash,
			salt,
			wmk_pin,
			wmk_recovery,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	var user models.User

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Salt,
		&user.WmkPin,
		&user.WmkRecovery,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
