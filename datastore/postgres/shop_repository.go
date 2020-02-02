package postgres

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
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
	// write some code here
	var shops []domain.Shop

	sql, args, err := sq.Select("id", "name").From("shops").Where(sq.Eq{"name": search}).ToSql()

	if err != nil {
		log.Printf("creating sql statement for finding shops by query failed: %v", err)
		return nil, err
	}

	rows, err := s.executor.QueryxContext(ctx, sql, args...)

	if err != nil {
		// TODO:
		// if isEmptyResultSet(err) {
		// 	return nil, nil
		// }

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var shop Shop
		err = rows.StructScan(&shop)
		if err != nil {
			return shops, err
		}

		shops = append(shops, *shop.toDomain())
	}

	return shops, err
}
