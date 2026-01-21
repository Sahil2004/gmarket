package database

import (
	"github.com/Sahil2004/gmarket/server/queries"
)

type Queries struct {
	*queries.UserQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgresDBConnection()

	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQueries: &queries.UserQueries{DB: db},
	}, nil
}