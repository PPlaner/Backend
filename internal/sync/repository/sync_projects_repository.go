package repository

import (
	"context"
	"database/sql"

	"github.com/PPlaner/Backend/internal/dto"
)

func (r *SyncRepository) UpsertProjectsTx(ctx context.Context, tx *sql.Tx, userID int, projects []dto.ProjectSyncDTO, syncSequence int) error {
	for _, project := range projects {
		var encryptedDataID int

		err := tx.QueryRowContext(
			ctx,
			`INSERT INTO encrypted_data (cipher_text, nonce, auth_tag)
			 VALUES ($1, $2, $3)
			 RETURNING id`,
			project.EncryptedData.CipherText,
			project.EncryptedData.Nonce,
			project.EncryptedData.AuthTag,
		).Scan(&encryptedDataID)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(
			ctx,
			`INSERT INTO projects (id, user_id, encrypted_data_id, sync_sequence, created_at, updated_at, deleted_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7)
			 ON CONFLICT (id) DO UPDATE SET
				encrypted_data_id = EXCLUDED.encrypted_data_id,
				sync_sequence = EXCLUDED.sync_sequence,
				updated_at = EXCLUDED.updated_at,
				deleted_at = EXCLUDED.deleted_at`,
			project.ID,
			userID,
			encryptedDataID,
			syncSequence,
			project.CreatedAt,
			project.UpdatedAt,
			project.DeletedAt,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SyncRepository) GetProjectChangesTx(ctx context.Context, tx *sql.Tx, userID int, cursor int) ([]dto.ProjectSyncDTO, error) {
	rows, err := tx.QueryContext(
		ctx,
		`SELECT 
			p.id,
			p.sync_sequence,
			p.created_at,
			p.updated_at,
			p.deleted_at,
			e.cipher_text,
			e.nonce,
			e.auth_tag
		FROM projects p
		JOIN encrypted_data e ON p.encrypted_data_id = e.id
		WHERE p.user_id = $1 AND p.sync_sequence > $2`,
		userID,
		cursor,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []dto.ProjectSyncDTO

	for rows.Next() {
		var project dto.ProjectSyncDTO

		err := rows.Scan(
			&project.ID,
			&project.SyncSequence,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.DeletedAt,
			&project.EncryptedData.CipherText,
			&project.EncryptedData.Nonce,
			&project.EncryptedData.AuthTag,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
