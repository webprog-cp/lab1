package models

type Cat struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Age       uint16 `json:"age"`
	Gender    string `json:"gender"`
	IsAdopted bool   `json:"adopted"`
	ShelterID uint64 `json:"shelterID"`
}
