package repository

import (
	"PDI-COBRANCA/build/package/database"
)

type users struct {
	name    string
	email   string
	keyword string
}

func InsertUsers(u users) (msg, err error) {

	u.name = "Guilherme"
	u.email = "gui@email.com"
	u.keyword = "12345"
	conn, err := database.OpenConn()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO users(name,email,keyword) VALUES ($1) RETURNING users;`

	err = conn.QueryRow(sql, u).Scan(msg)

	return err, msg
}
