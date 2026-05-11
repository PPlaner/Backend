package models

import "time"

type Note struct {
	ID              int        `json:"id"`
	UserID          int        `json:"userId"`
	ProjectID       int        `json:"projectId"`
	EncryptedDataID int        `json:"encryptedDataId"`
	Version         int        `json:"version"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	SyncSequence    int        `json:"syncSequence"`
}
