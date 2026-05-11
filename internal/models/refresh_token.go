package models

import "time"

type RefreshToken struct {
	ID        int        `json:"id"`
	UserID    int        `json:"userId"`
	TokenHash string     `json:"tokenHash"`
	ExpiresAt time.Time  `json:"expiresAt"`
	CreatedAt time.Time  `json:"createdAt"`
	RevokedAt *time.Time `json:"revokedAt"`
}
