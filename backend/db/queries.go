package db

import (
	"Market/backend"
	error2 "Market/error"
	"database/sql"
)

type Password struct {
	Password string `json:"password"`
}

func GetAllProduct(DB *sql.DB) ([]backend.Products, error) {
	query := `SELECT name,description,picture,price from products`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, error2.Wrap("Can`t receive products from db", err)
	}

	defer rows.Close()

	var products []backend.Products

	for rows.Next() {
		var product backend.Products

		if err := rows.Scan(&product.Name, &product.Description, &product.Picture_URL, &product.Price); err != nil {
			return nil, error2.Wrap("Can`t scan products from db", err)
		}

		products = append(products, product)
	}

	return products, nil

}

func GetAllLoginAndEmail(DB *sql.DB) ([]backend.User, error) {
	query := `SELECT login from users`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, error2.Wrap("Can`t receive users from db", err)
	}

	defer rows.Close()

	var logins []backend.User

	for rows.Next() {
		var user backend.User

		if err := rows.Scan(&user.Login); err != nil {
			return nil, error2.Wrap("Can`t scan products from db", err)
		}

		logins = append(logins, user)
	}

	return logins, nil

}

func AddUser(DB *sql.DB, user backend.User) error {
	query := `INSERT INTO users (login,password,email) VALUES ($1,$2,$3)`
	_, err := DB.Exec(query, user.Login, user.Password, user.Email)
	if err != nil {
		return error2.Wrap("Can`t add user to db", err)
	}
	return nil
}

//func GetUser(DB *sql.DB, login string) (string, error) {
//	query := `SELECT password FROM users WHERE login=$1`
//
//	row := DB.QueryRow(query, login)
//
//	var password Password
//	if err := row.Scan(&password.Password); err != nil {
//		return "", error2.Wrap("Can`t get user password from db", err)
//	}
//
//	if
//	return password.Password, nil
//
//}
//
