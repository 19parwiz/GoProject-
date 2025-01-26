// pkg/database/db.go
package database

import (
	"bookstore/pkg/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.GetDBConnectionString())
	if err != nil {
		return nil, err
	}
	return db, nil
}
