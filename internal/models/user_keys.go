package models

type UserKeys struct {
	ID               int    `json:"id"`
	UserID           int    `json:"userId"`
	KeyType          int    `json:"keyType"`
	Salt             []byte `json:"salt"`
	WrappedMasterKey []byte `json:"wrappedMasterKey"`
}
