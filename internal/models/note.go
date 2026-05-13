package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID  `json:"id"`
	UserID    int        `json:"userId"`
	ProjectID *uuid.UUID `json:"projectId"`

	EncryptedTitle   []byte `json:"encryptedTitle"`
	EncryptedContent []byte `json:"encryptedContent"`

	Version      int   `json:"version"`
	SyncSequence int64 `json:"syncSequence"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
