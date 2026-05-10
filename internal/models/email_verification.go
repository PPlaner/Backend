package models

import (
	"database/sql"
	"time"
)

type EmailVerification struct {
	ID        int
	Email     string
	CodeHash  string
	ExpiresAt time.Time
	UsedAt    sql.NullTime
	CreatedAt time.Time
}
