package models

import "time"

type Project struct {
	ID              int        `json:"id"`
	UserID          int        `json:"userId"`
	EncryptedDataID int        `json:"encryptedDataId"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	DeletedAt       *time.Time `json:"deletedAt"`
	SyncSequence    int        `json:"syncSequence"`
}
