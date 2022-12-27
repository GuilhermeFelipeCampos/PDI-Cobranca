package repository

import (
	"PDI-COBRANCA/build/package/database"
	"PDI-COBRANCA/build/package/model"
	"fmt"
	"strings"
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
	formatedEmail := strings.ReplaceAll(email, "{", "")
	formatedEmail = strings.ReplaceAll(formatedEmail, "}", "")
	conn, err := database.OpenConn()

	if err != nil {
		return
	}
	defer conn.Close()
	sql := `SELECT * FROM users WHERE email=$1`
	rows, err := conn.Query(sql, formatedEmail)
	for rows.Next() {
		var u model.Users
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.Keyword)
		if err != nil {
			continue
		}

		us = append(us, u)
	}

	return us, nil

}

// Alterar função para que receba um id e altere os campos referente ao id
func UpdateUser(user model.Users) (u model.Users, err error) {
	conn, err := database.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `UPDATE users SET name=$1,email=$2,Keyword=$3 WHERE id=$4 RETURNING id, name, email, keyword`
	err = conn.QueryRow(sql, user.Name, user.Email, user.Keyword, user.Id).Scan(&u.Id, &u.Name, &u.Email, &u.Keyword)

	if err != nil {
		panic(err)
	}

	return u, nil
}

func DeleteUser(id string) (userId string, err error) {
	conn, err := database.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()
	sql := `DELETE FROM users WHERE id=$1`
	err = conn.QueryRow(sql, id).Scan(&userId)
	if err != nil {
		panic(err)
	}

	return userId, nil
}
