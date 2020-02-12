package search_shop

import (
	"context"

	"github.com/der-nackte-halloumi/api/domain"
)

type Repository interface {
	FindShopsByQuery(ctx context.Context, search string, lat float32, long float32) ([]domain.Shop, error)
	GetShops(ctx context.Context) ([]domain.Shop, error)
}
