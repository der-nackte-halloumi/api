package postgres

import "github.com/gofrs/uuid"

import "github.com/der-nackte-halloumi/api/domain"

type Shop struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
	// TODO: add products here
}

func (s *Shop) toDomain() *domain.Shop {
	return &domain.Shop{
		ID:   s.ID,
		Name: s.Name,
	}
}
