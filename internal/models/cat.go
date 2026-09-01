package models

type Cat struct {
	ID        uint64 `json:"id"`
	name      string `json:"name"`
	age       uint16 `json:"age"`
	gender    string `json:"gender"`
	isAdopted bool   `json:"adopted"`
	shelterID uint64 `json:"shelterID"`
}
