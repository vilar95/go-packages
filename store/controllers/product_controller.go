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
		log.Println("Nome do Produto:", name)
		log.Println("Descrição do Produto:", description)
		log.Println("Preço do Produto:", price)
		log.Println("Quantidade do Produto:", quantity)

		parsePrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			http.Error(w, "Ocorreu um erro ao converter o preço", http.StatusBadRequest)
			return
		}

		parseQuantity, err := strconv.Atoi(quantity)
		if err != nil {
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

func EditProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.GetProductByID(idProduct)
	temp.ExecuteTemplate(w, "EditProduct", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		log.Println("Inicia a atualização do produto")
		log.Println("ID do Produto:", id)
		log.Println("Nome do Produto:", name)
		log.Println("Descrição do Produto:", description)
		log.Println("Preço do Produto:", price)
		log.Println("Quantidade do Produto:", quantity)
		parseID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Ocorreu um erro ao converter o ID", http.StatusBadRequest)
			return
		}
		parsePrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			http.Error(w, "Ocorreu um erro ao converter o preço", http.StatusBadRequest)
			return
		}
		parseQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Ocorreu um erro ao converter a quantidade", http.StatusBadRequest)
			return
		}
		models.UpdateProductByID(parseID, name, description, parsePrice, parseQuantity)
	}

	// Redireciona para a página inicial após a atualização do produto
	// Retorna status 301 Movido Permanentemente
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
