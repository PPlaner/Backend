package repository

import (
	"context"
	"database/sql"

	"github.com/PPlaner/Backend/internal/dto"
)

func (r *SyncRepository) UpsertNotesTx(ctx context.Context, tx *sql.Tx, userID int, notes []dto.NoteSyncDTO, syncSequence int) error {
	for _, note := range notes {
		var encryptedDataID int

		encryptedDataID, err := r.insertEncryptedData(ctx, note.EncryptedData)
		if err != nil {
			return err
		}

		_, err = r.db.ExecContext(
			ctx,
			`INSERT INTO notes (id, user_id, project_id, encrypted_data_id, version, sync_sequence, created_at, updated_at, deleted_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			 ON CONFLICT (id) DO UPDATE SET
				project_id = EXCLUDED.project_id,
				encrypted_data_id = EXCLUDED.encrypted_data_id,
				version = EXCLUDED.version,
				sync_sequence = EXCLUDED.sync_sequence,
				updated_at = EXCLUDED.updated_at,
				deleted_at = EXCLUDED.deleted_at`,
			note.ID,
			userID,
			note.ProjectID,
			encryptedDataID,
			note.Version,
			syncSequence,
			note.CreatedAt,
			note.UpdatedAt,
			note.DeletedAt,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SyncRepository) GetNotesChangesTx(ctx context.Context, tx *sql.Tx, userID int, cursor int) ([]dto.NoteSyncDTO, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT 
			n.id,
			n.project_id,
			n.version,
			n.sync_sequence,
			n.created_at,
			n.updated_at,
			n.deleted_at,
			e.cipher_text,
			e.nonce,
			e.auth_tag
		FROM notes n
		JOIN encrypted_data e ON n.encrypted_data_id = e.id
		WHERE n.user_id = $1 AND n.sync_sequence > $2`,
		userID,
		cursor,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []dto.NoteSyncDTO

	for rows.Next() {
		var note dto.NoteSyncDTO

		err := rows.Scan(
			&note.ID,
			&note.ProjectID,
			&note.Version,
			&note.SyncSequence,
			&note.CreatedAt,
			&note.UpdatedAt,
			&note.DeletedAt,
			&note.EncryptedData.CipherText,
			&note.EncryptedData.Nonce,
			&note.EncryptedData.AuthTag,
		)
		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}
