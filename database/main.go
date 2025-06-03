package main

import (
	"database/sql"
	// O blank identifier é usado para importar o pacote sem usar diretamente
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(id, name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	Product := NewProduct(uuid.New().String(), "NoteBook", 1999.99)

	err = insertProduct(db, Product)
	if err != nil {
		panic(err)
	}
	Product.Price = 1799.99
	// err = updateProduct(db, Product)
	// if err != nil {
	// 	panic(err)
	// }

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil

}
