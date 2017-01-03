package postgres

import (
	"database/sql"
	_ "encoding/json"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

var Database *sql.DB

func Open_postgres(defaultUrl string) {
	var db_url = ""

	if db_url == "" {
		db_url = defaultUrl
	}

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal(err)
	}

	Database = db

	//  configure database pooling
	mconns := os.Getenv("MAX_DB_CONNECTIONS")
	maxconns := 20
	if (mconns != "") { maxconns,_ = strconv.Atoi(mconns) }

	SetMaxIdleConns(1)
	SetMaxOpenConns(maxconns)
}
