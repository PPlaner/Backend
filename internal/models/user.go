package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Salt         string    `json:"salt"`
	WmkPin       string    `json:"wmk_pin"`
	WmkRecovery  string    `json:"wmk_recovery"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
