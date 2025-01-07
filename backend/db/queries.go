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
