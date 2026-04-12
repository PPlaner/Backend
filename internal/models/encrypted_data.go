package models

import "time"

type EncryptedData struct {
	ID         int       `json:"id"`
	CipherText string    `json:"cipher_text"`
	Nonce      string    `json:"nonce"`
	AuthTag    string    `json:"auth_tag"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
