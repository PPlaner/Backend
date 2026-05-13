package dto

import (
	"time"
)

type SyncRequest struct {
	Cursor  int64       `json:"cursor"`
	Changes SyncChanges `json:"changes"`
}

type SyncResponse struct {
	NextCursor int64       `json:"next_cursor"`
	Changes    SyncChanges `json:"changes"`
}

type SyncChanges struct {
	Project []ProjectSyncDTO `json:"project"`
	Note    []NoteSyncDTO    `json:"note"`
}

type ProjectSyncDTO struct {
	ID               string     `json:"id"`
	EncryptedContent []byte     `json:"encrypted_content"`
	Version          int        `json:"version"`
	SyncSequence     int64      `json:"sync_sequence"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}

type NoteSyncDTO struct {
	ID               string     `json:"id"`
	ProjectID        *string    `json:"project_id"`
	EncryptedTitle   []byte     `json:"encrypted_title"`
	EncryptedContent []byte     `json:"encrypted_content"`
	Version          int        `json:"version"`
	SyncSequence     int        `json:"sync_sequence"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}
