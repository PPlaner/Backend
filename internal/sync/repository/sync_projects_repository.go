package repository

import (
	"context"
	"database/sql"

	"github.com/PPlaner/Backend/internal/dto"
)

func (r *SyncRepository) UpsertProjectsTx(ctx context.Context, tx *sql.Tx, userID int, projects []dto.ProjectSyncDTO, syncSequence int64) error {
	for _, project := range projects {
		_, err := tx.ExecContext(
			ctx,
			`INSERT INTO projects (id, user_id, encrypted_content, version, sync_sequence, created_at, updated_at, deleted_at)
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			 ON CONFLICT (id) DO UPDATE SET
				encrypted_content = EXCLUDED.encrypted_content,
				version = EXCLUDED.version,
				sync_sequence = EXCLUDED.sync_sequence,
				updated_at = EXCLUDED.updated_at,
				deleted_at = EXCLUDED.deleted_at`,
			project.ID,
			userID,
			project.EncryptedContent,
			project.Version,
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

func (r *SyncRepository) GetProjectChangesTx(ctx context.Context, tx *sql.Tx, userID int, cursor int64) ([]dto.ProjectSyncDTO, error) {
	rows, err := tx.QueryContext(
		ctx,
		`SELECT id, encrypted_content, version, sync_sequence, created_at, updated_at, deleted_at
		FROM projects WHERE user_id = $1 AND sync_sequence > $2`,
		userID, cursor,
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
			&project.EncryptedContent,
			&project.Version,
			&project.SyncSequence,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	return projects, nil
}
