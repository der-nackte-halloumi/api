package domain

import "github.com/gofrs/uuid"

type Shop struct {
	ID       uuid.UUID
	Name     string
	Lat      float32
	Lng      float32
	Address  string
	Products []Product
}
