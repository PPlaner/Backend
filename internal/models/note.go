package models

import "time"

type Note struct {
	ID              int        `json:"id"`
	UserID          int        `json:"user_id"`
	ProjectID       int        `json:"project_id"`
	EncryptedDataID int        `json:"encrypted_data_id"`
	Version         int        `json:"version"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
