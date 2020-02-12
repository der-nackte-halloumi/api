package domain

import "github.com/gofrs/uuid"

type Product struct {
	ID   uuid.UUID
	Name string
}
