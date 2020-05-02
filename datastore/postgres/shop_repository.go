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

func (s *shopRepository) GetShops(ctx context.Context) ([]domain.Shop, error) {
	var shops []domain.Shop

	sql, args, err := queryBuilder.Select("id", "name", "address", "lat", "long").From("shops").ToSql()

	if err != nil {
		log.Printf("creating sql statement for retrieving all shops failed: %v", err)
		return nil, err
	}

	rows, err := s.executor.QueryxContext(ctx, sql, args...)

	if err != nil {
		log.Printf("error when executing query for retrieving all shops: %v", err)
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

	return shops, nil
}

func (s *shopRepository) FindShopsByQuery(ctx context.Context, search string, lat float32, long float32) ([]domain.Shop, error) {
	var shops []domain.Shop
	var productIds []string

	sql, args, err := queryBuilder.
		Select("products.id").
		Distinct().
		From("products").
		LeftJoin("categories_translations on categories_translations.category_id = products.category_id").
		Join("products_translations on products_translations.product_id = products.id").
		Where(sq.Or{
			sq.ILike{"categories_translations.value": "%" + search + "%"},
			sq.ILike{"products_translations.value": "%" + search + "%"},
		}).ToSql()

	if err != nil {
		log.Printf("creating sql statement for finding product ids by query failed: %v", err)
		return nil, err
	}

	rows, err := s.executor.QueryxContext(ctx, sql, args...)

	if err != nil {
		log.Printf("error when executing query for product ids by query", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		productIds = append(productIds, id)
	}

	if len(productIds) == 0 {
		return nil, nil
	}

	sql, args, err = queryBuilder.
		Select("shops.id", "shops.name", "shops.address", "shops.lat", "shops.long").
		Distinct().
		From("shops").
		Join("shops_products on shops.id = shops_products.shop_id").
		Where(sq.Eq{"shops_products.product_id": productIds}).ToSql()

	if err != nil {
		log.Printf("creating sql statement for finding shops by query failed: %v", err)
		return nil, err
	}

	rows, err = s.executor.QueryxContext(ctx, sql, args...)

	if err != nil {
		log.Println("error when executing query for finding shops by query", err)
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
