package db

import (
	"Market/backend"
	error2 "Market/error"
	"database/sql"
)

func GetAllProduct(DB *sql.DB, id int) (backend.Products, error) {
	query := `SELECT name,description,picture,price from products where id=$1`
	rows := DB.QueryRow(query, id)

	var products backend.Products

	if err := rows.Scan(&products.Name, &products.Description, &products.Picture_URL, &products.Price); err != nil {
		return backend.Products{}, error2.Wrap("Can`t scan data from db", err)
	}

	return products, nil

}
