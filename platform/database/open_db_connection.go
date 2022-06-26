package database

import (
	query_v1 "github.com/hotrungnhan/go-fiber-template/app/v1/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*query_v1.UserQueries // load queries from User model
	*query_v1.BookQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}
	return &Queries{
		// Set queries from models:
		UserQueries: &query_v1.UserQueries{DB: db}, // from User model
		BookQueries: &query_v1.BookQueries{DB: db}, // from Book model
	}, nil
}
