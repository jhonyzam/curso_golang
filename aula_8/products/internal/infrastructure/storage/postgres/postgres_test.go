package postgres_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var (
	postgresURL string
	TestDb      *sql.DB
)

func init() {
	postgresURL := os.Getenv("POSTGRES_URL")
	if postgresURL != "" {
		if db, err := sql.Open("postgres", postgresURL); err != nil {
			fmt.Errorf("error on PostgreSQL connection: %q", err)
		} else {
			TestDb = db
		}
	}
}

func TestNewConnection(t *testing.T) {
	//db, err := postgres.NewConnection(os.Getenv("POSTGRES_URL"))
	//assert.NoError(t, err)
	//assert.NotNil(t, db)
	//
	//if err == nil {
	//	defer db.Close()
	//	rows, err := db.Query("select 1")
	//	assert.Nil(t, err, "Connection failed. Check the environment variables.")
	//	assert.NotNil(t, rows)
	//	if rows != nil {docker
	//		defer rows.Close()
	//	}
	//}
}
