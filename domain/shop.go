package domain

import "github.com/gofrs/uuid"

type Shop struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Lat      float32   `json:"lat"`
	Lng      float32   `json:"lng"`
	Address  string    `json:"address"`
	Products []Product `json:"products"`
}
