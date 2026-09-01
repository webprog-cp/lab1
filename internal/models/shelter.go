package models

type Shelter struct {
	ID       uint64 `json:"id"`
	name     string `json:"name"`
	city     string `json:"city"`
	address  string `json:"address"`
	capacity uint64 `json:"capacity"`
}
