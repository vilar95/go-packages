package models

import (
	"store/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

type Products struct {
	Products []Product
}

func SearchAllProducts() Products {
	db := db.ConnectDB()

	selectProducts, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	products := Products{}
	for selectProducts.Next() {
		var product Product
		err := selectProducts.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			panic(err.Error())
		}
		products.Products = append(products.Products, product)
	}
	defer db.Close()
	return products
}
