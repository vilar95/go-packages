package controllers

import (
	"html/template"
	"log"
	"net/http"
	"store/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	AllProduct := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", AllProduct)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "AddProduct", nil)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		log.Println("Incia a inserção do produto")
		log.Println("Nome do Produto:", name, "\nDescrição:", description, "\nPreço:", price, "\nQuantidade:", quantity)

		parsePrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Ocorreu um erro ao converter o preço")
			http.Error(w, "Ocorreu um erro ao converter o preço", http.StatusBadRequest)
			return
		}

		parseQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Ocorreu um erro ao converter a quantidade")
			http.Error(w, "Ocorreu um erro ao converter a quantidade", http.StatusBadRequest)
			return
		}

		models.CreateNewProduct(name, description, parsePrice, parseQuantity)
	}

	// Redireciona para a página inicial após a inserção do produto
	// Retorna status 301 Movido Permanentemente
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProductByID(idProduct)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}