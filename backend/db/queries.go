package db

import (
	"Market/backend"
	error2 "Market/error"
	"database/sql"
)

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

func AddUser(DB *sql.DB, user backend.User) error {
	query := `INSERT INTO users (login, password,email) VALUES ($1,$2,$3)`
	_, err := DB.Exec(query, user.Login, user.Password, user.Email)
	if err != nil {
		return error2.Wrap("Can`t add user to db", err)
	}

	return nil
}
