package postgres

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/der-nackte-halloumi/api/datastore/sql_datastore"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var queryBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewDatastore(host string, port string, user string, name string, password string) (sql_datastore.Datastore, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		name,
		password,
	)

	return sqlx.Connect("postgres", connectionString)
}
