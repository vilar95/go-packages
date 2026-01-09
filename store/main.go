package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connection := "user=postgres dbname=insider_store password=@12345678 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

type PageData struct {
	Products []Product
}

func main() {
	http.HandleFunc("/", index)
	println("Servidor rodando em http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()

	selectProducts, err := db.Query("SELECT * FROM products")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := PageData{}
	for selectProducts.Next() {
		var product Product
		err := selectProducts.Scan(&product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products.Products = append(products.Products, product)
	}
	defer db.Close()

	err = temp.ExecuteTemplate(w, "index", products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
