package dto

import (
	"time"
)

type SyncRequest struct {
	Cursor  int         `json:"cursor"`
	Changes SyncChanges `json:"changes"`
}

type SyncResponse struct {
	NextCursor int         `json:"next_cursor"`
	Changes    SyncChanges `json:"changes"`
}

type SyncChanges struct {
	Project []ProjectSyncDTO `json:"project"`
	Note    []NoteSyncDTO    `json:"note"`
}

type EncryptedDataDTO struct {
	CipherText string `json:"cipher_text"`
	Nonce      string `json:"nonce"`
	AuthTag    string `json:"auth_tag"`
}
type ProjectSyncDTO struct {
	ID            int              `json:"id"`
	EncryptedData EncryptedDataDTO `json:"encrypted_data"`
	Version       int              `json:"version"`
	SyncSequence  int              `json:"sync_sequence"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	DeletedAt     *time.Time       `json:"deleted_at,omitempty"`
}

type NoteSyncDTO struct {
	ID            int              `json:"id"`
	ProjectID     *int             `json:"project_id"`
	EncryptedData EncryptedDataDTO `json:"encrypted_data"`
	Version       int              `json:"version"`
	SyncSequence  int              `json:"sync_sequence"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
	DeletedAt     *time.Time       `json:"deleted_at,omitempty"`
}
