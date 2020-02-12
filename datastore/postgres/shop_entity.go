package postgres

import "github.com/gofrs/uuid"

import "github.com/der-nackte-halloumi/api/domain"

type Shop struct {
	ID      uuid.UUID `db:"id"`
	Name    string    `db:"name"`
	Address string    `db:"address"`
	Lat     float32   `db:"lat"`
	Lng     float32   `db:"long"`
	// TODO: add products here
}

func (s *Shop) toDomain() *domain.Shop {
	return &domain.Shop{
		ID:      s.ID,
		Name:    s.Name,
		Address: s.Address,
		Lat:     s.Lat,
		Lng:     s.Lng,
	}
}
