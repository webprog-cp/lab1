package models

type Shelter struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Capacity uint64 `json:"capacity"`
}
