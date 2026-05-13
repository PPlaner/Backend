package repository

import (
	"context"
	"database/sql"

	"github.com/PPlaner/Backend/internal/dto"
)

type SyncRepository struct {
	db *sql.DB
}

func NewSyncRepository(db *sql.DB) *SyncRepository {
	return &SyncRepository{
		db: db,
	}
}

// --------------------------------------------------------
// ---------------------- Транзакція ----------------------
// --------------------------------------------------------

func (r *SyncRepository) SyncData(ctx context.Context, userID int, req dto.SyncRequest) (dto.SyncResponse, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return dto.SyncResponse{}, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	currentCursor, err := r.GetUserCursorTx(ctx, tx, userID)
	if err != nil {
		return dto.SyncResponse{}, err
	}

	clientCursor := req.Cursor
	hasChanges := len(req.Changes.Note) > 0 || len(req.Changes.Project) > 0

	newCursor := currentCursor

	if hasChanges {
		newCursor = currentCursor + 1

		err = r.UpdateUserCursorTx(ctx, tx, userID, newCursor)
		if err != nil {
			return dto.SyncResponse{}, err
		}

		err = r.UpsertNotesTx(ctx, tx, userID, req.Changes.Note, newCursor)
		if err != nil {
			return dto.SyncResponse{}, err
		}

		err = r.UpsertProjectsTx(ctx, tx, userID, req.Changes.Project, newCursor)
		if err != nil {
			return dto.SyncResponse{}, err
		}
	}

	notes, err := r.GetNotesChangesTx(ctx, tx, userID, clientCursor)
	if err != nil {
		return dto.SyncResponse{}, err
	}

	projects, err := r.GetProjectChangesTx(ctx, tx, userID, clientCursor)
	if err != nil {
		return dto.SyncResponse{}, err
	}

	response := dto.SyncResponse{
		NextCursor: newCursor,
		Changes: dto.SyncChanges{
			Note:    notes,
			Project: projects,
		},
	}

	err = tx.Commit()
	if err != nil {
		return dto.SyncResponse{}, err
	}
	return response, nil
}
