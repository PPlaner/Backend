package models

import "time"

type EncryptedData struct {
	ID         int       `json:"id"`
	CipherText string    `json:"cipherText"`
	Nonce      string    `json:"nonce"`
	AuthTag    string    `json:"authTag"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
