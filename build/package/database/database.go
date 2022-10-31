package database

import (
	"PDI-COBRANCA/build/package/configs"
	"database/sql"
	"fmt"
)

var cfg = configs.ConfigsDB{}

func OpenConn() (*sql.DB, error) {

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open(cfg.DBUser, conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}
