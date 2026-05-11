package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	Salt         string    `json:"salt"`
	WmkPin       string    `json:"wmkPin"`
	WmkRecovery  string    `json:"wmkRecovery"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	SyncCursor   int       `gorm:"default:0" json:"syncCursor"`
}
