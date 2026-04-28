package repository

import (
	"context"
	"database/sql"
)

func (r *SyncRepository) GetUserCursorTx(ctx context.Context, tx *sql.Tx, userID int) (int, error) {
	var cursor int

	err := r.db.QueryRowContext(
		ctx,
		"SELECT sync_cursor FROM users WHERE id = $1",
		userID,
	).Scan(&cursor)

	if err != nil {
		return 0, err
	}

	return cursor, nil
}

func (r *SyncRepository) UpdateUserCursorTx(ctx context.Context, tx *sql.Tx, userID int, newCursor int) error {
	_, err := r.db.ExecContext(
		ctx,
		"UPDATE users SET sync_cursor = $1 WHERE id = $2",
		newCursor,
		userID,
	)

	return err
}
