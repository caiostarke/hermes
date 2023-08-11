package study_case

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Repo struct {
	queries Queries
}

func OpenDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal("error opening database: ", err)
	}

	return db
}
