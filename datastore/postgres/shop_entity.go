package postgres

import "github.com/gofrs/uuid"

type Shop struct {
	ID uuid.UUID `db:"id"`
}
