package postgres

import (
	"database/sql"
	"log"

	// driver postgres
	_ "github.com/lib/pq"
)

func NewConnection(postgresURL string) (*sql.DB, error) {

	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		log.Printf("ERROR: on PostgreSQL connection: %q\n", err)
		return nil, err
	}
	return db, nil
}
