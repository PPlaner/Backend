package dto

import "time"

type KeyDTO struct {
	Type             string    `json:"type" binding:"required"`
	Salt             []byte    `json:"salt" binding:"required"`
	WrappedMasterKey []byte    `json:"wrappedMasterKey" binding:"required"`
	UpdatedAt        time.Time `json:"updatedAt" binding:"required"`
}

type KeysDTO struct {
	Keys []KeyDTO `json:"keys" binding:"required,dive"`
}
