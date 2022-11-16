package database

import (
	"PDI-COBRANCA/build/package/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var cfg = configs.ConfigsDB{}

func OpenConn() (*sql.DB, error) {

	conn := fmt.Sprint("host=localhost port=5432 user=postgres password=admin dbname=pdi_cobranca_db sslmode=disable")
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}
