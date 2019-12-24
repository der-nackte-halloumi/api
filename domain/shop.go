package domain

import "github.com/gofrs/uuid"

type Shop struct {
	ID       uuid.UUID
	Name     string
	lat      float32
	lng      float32
	Products []Product
}
