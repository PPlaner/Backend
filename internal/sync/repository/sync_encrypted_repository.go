package repository

import (
	"context"

	"github.com/PPlaner/Backend/internal/dto"
)

func (r *SyncRepository) insertEncryptedData(ctx context.Context, data dto.EncryptedDataDTO) (int, error) {
	var encryptedDataID int

	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO encrypted_data (cipher_text, nonce, auth_tag)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		data.CipherText,
		data.Nonce,
		data.AuthTag,
	).Scan(&encryptedDataID)

	if err != nil {
		return 0, err
	}

	return encryptedDataID, nil
}
