package postgres

import (
	"context"

	"github.com/der-nackte-halloumi/api/datastore/sql_datastore"
	"github.com/der-nackte-halloumi/api/domain"
)

type shopRepository struct {
	executor sql_datastore.Executor
}

func NewShopRepository(executor sql_datastore.Executor) *shopRepository {
	return &shopRepository{executor: executor}
}

func (s *shopRepository) FindShopsByQuery(ctx context.Context, search string, lat float32, long float32) ([]domain.Shop, error) {
	return nil, nil
}
