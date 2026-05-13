package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID     uuid.UUID `json:"id"`
	UserID int       `json:"userId"`

	EncryptedContent []byte `json:"encryptedContent"`

	Version      int   `json:"version"`
	SyncSequence int64 `json:"syncSequence"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
