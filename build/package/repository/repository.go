package repository

import (
	"PDI-COBRANCA/build/package/database"
	"PDI-COBRANCA/build/package/model"
	"fmt"
)

type RepositoryInterface interface {
	InsertUsers(u model.Users) (msg, err error)
	GetUsers() (us []model.Users, err error)
}

func InsertUsers(u model.Users) (msg, err error) {

	conn, err := database.OpenConn()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO users(name,email,keyword) VALUES ($1,$2,$3) RETURNING users;`

	err = conn.QueryRow(sql, u.Name, u.Email, u.Keyword).Scan(msg)

	return err, msg
}

func GetUsers() (us []model.Users, err error) {

	conn, err := database.OpenConn()

	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM users;`

	rows, err := conn.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u model.Users
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Keyword)
		if err != nil {
			continue
		}

		us = append(us, u)
	}
	fmt.Println(us)
	return us, err
}

func GetUserByEmail(email string) (us []model.Users, err error) {

	conn, err := database.OpenConn()

	fmt.Println("repositório: " + email)
	if err != nil {
		return
	}
	defer conn.Close()

	//sql := `SELECT * FROM users WHERE email=$1`

	rows, err := conn.Query(`SELECT * FROM users WHERE email=$1`, email)
	for rows.Next() {
		var u model.Users
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Keyword)
		if err != nil {
			continue
		}

		us = append(us, u)
	}

	fmt.Println(email)
	return us, nil

}
