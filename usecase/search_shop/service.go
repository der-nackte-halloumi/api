package search_shop

import (
	"context"

	"github.com/der-nackte-halloumi/api/domain"
)

type Service interface {
	FindShopsByQuery(ctx context.Context, search string, lat float32, long float32) ([]domain.Shop, error)
	GetShops(ctx context.Context) ([]domain.Shop, error)
}

type service struct {
	// logger
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindShopsByQuery(ctx context.Context, search string, lat float32, long float32) ([]domain.Shop, error) {
	// TODO: logging, access control, validation
	return s.repository.FindShopsByQuery(ctx, search, lat, long)
}

func (s *service) GetShops(ctx context.Context) ([]domain.Shop, error) {
	// TODO: logging, access control, validation
	return s.repository.GetShops(ctx)
}
