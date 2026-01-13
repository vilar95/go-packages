package models

import (
	"log"
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

	// Executa a query de seleção de todos os produtos
	selectProducts, err := db.Query("SELECT * FROM products")

	if err != nil {
		log.Println("Não foi possível fazer a busca de todos os produtos no banco de dados")
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
	// Fecha a conexão com o banco de dados
	defer db.Close()

	return products
}

func CreateNewProduct(name string, description string, price float64, quantity int) {
	log.Println("Iniciando a criação de um novo produto")

	db := db.ConnectDB()

	// Prepara a query de inserção com os valores recebidos
	insertProduct, err := db.Prepare("INSERT INTO products(nome, descricao, preco, quantidade) VALUES($1,$2,$3,$4)")

	if err != nil {
		log.Println("Não foi possível preparar a inserção do novo produto no banco de dados")
		panic(err.Error())
	}
	// Executa a query de inserção
	insertProduct.Exec(name, description, price, quantity)

	log.Println("Produto inserido com sucesso")

	// Fecha a conexão com o banco de dados
	defer db.Close()
}

func DeleteProductByID(id string) {
	log.Println("Iniciando a exclusão do produto")
	db := db.ConnectDB()
	// Prepara a query de exclusão com o ID recebido
	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		log.Println("Não foi possível preparar a exclusão do produto no banco de dados")
		panic(err.Error())
	}
	// Executa a query de exclusão
	deleteProduct.Exec(id)
	log.Println("Produto excluído com sucesso")
	// Fecha a conexão com o banco de dados
	defer db.Close()
}