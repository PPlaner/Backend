package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/PPlaner/Backend/internal/dto"
)

type KeyRepository struct {
	db *sql.DB
}

func NewKeyRepository(db *sql.DB) *KeyRepository {
	return &KeyRepository{db: db}
}

func (r *KeyRepository) Upsert(userID int, key dto.KeyDTO) error {
	query := `
		INSERT INTO user_keys (user_id, key_type, salt, wrapped_master_key, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id, key_type) DO UPDATE SET
			salt = EXCLUDED.salt,
			wrapped_master_key = EXCLUDED.wrapped_master_key,
			updated_at = EXCLUDED.updated_at
	`

	_, err := r.db.Exec(query, userID, key.Type, key.Salt, key.WrappedMasterKey)
	return err
}

func (r *KeyRepository) UpsertMany(userID int, keys []dto.KeyDTO) error {
	if len(keys) == 0 {
		return nil
	}

	numFields := 5
	placeholderGroups := make([]string, len(keys))
	args := make([]any, 0, len(keys)*numFields)

	for i, key := range keys {
		base := i * numFields
		placeholderGroups[i] = fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", base+1, base+2, base+3, base+4, base+5)

		args = append(args, userID, key.Type, key.Salt, key.WrappedMasterKey, key.UpdatedAt)
	}

	query := fmt.Sprintf(`
        INSERT INTO user_keys (user_id, key_type, salt, wrapped_master_key, updated_at)
        VALUES %s
        ON CONFLICT (user_id, key_type) DO UPDATE SET
            salt = EXCLUDED.salt,
            wrapped_master_key = EXCLUDED.wrapped_master_key,
			updated_at = EXCLUDED.updated_at
    `, strings.Join(placeholderGroups, ", "))

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *KeyRepository) GetByUserID(userID int) ([]dto.KeyDTO, error) {
	query := `
		SELECT key_type, salt, wrapped_master_key, updated_at
		FROM user_keys WHERE user_id = $1
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keys []dto.KeyDTO
	for rows.Next() {
		var k dto.KeyDTO

		if err := rows.Scan(&k.Type, &k.Salt, &k.WrappedMasterKey, &k.UpdatedAt); err != nil {
			return nil, err
		}

		keys = append(keys, k)
	}

	return keys, nil
}
