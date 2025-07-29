package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID         int        `gorm:"primaryKey"`                    // Chave primária da categoria
	Name       string     `gorm:"size:100"`                      // Tamanho máximo de 100 caracteres para o nome da categoria
	Products   []Products `gorm:"many2many:products_categories"` // Relacionamento com Products
	gorm.Model            // O Model serve para adicionar campos de timestamp (CreatedAt, UpdatedAt, DeletedAt)
}

type Products struct {
	ID           int          `gorm:"primaryKey"`                    // Chave primária do produto
	Name         string       `gorm:"size:100"`                      // Tamanho máximo de 100 caracteres para o nome do produto
	Price        float64      `gorm:"size:10"`                       // Tamanho máximo de 10 dígitos para o preço
	CategoryID   int          `gorm:"not null"`                      // Chave estrangeira para Category
	Categories   []Category   `gorm:"many2many:products_categories"` // Relacionamento com Category
	SerialNumber SerialNumber `gorm:"foreignKey:ProductID"`          // Relacionamento com SerialNumber
	gorm.Model                // O Model serve para adicionar campos de timestamp (CreatedAt, UpdatedAt, DeletedAt)
}

type SerialNumber struct {
	ID        int    `gorm:"primaryKey"` // Chave primária do número de série
	Number    string `gorm:"size:100"`   // Tamanho máximo de 100 caracteres para o número de série
	ProductID int    `gorm:"not null"`   // Chave estrangeira para Products

	gorm.Model // O Model serve para adicionar campos de timestamp (CreatedAt, UpdatedAt, DeletedAt)
}

func main() {
	// Configuração do banco de dados MySQL
	// O dsn (Data Source Name) é a string de conexão com o banco de dados
	// O dsn deve ser ajustado de acordo com suas credenciais e configurações do MySQL
	// o chartset=utf8mb4 é usado para suportar caracteres especiais
	// parseTime=True é usado para converter os campos de data e hora em tipos do Go
	// loc=Local é usado para definir o fuso horário local
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrar os modelos para o banco de dados
	// Isso cria as tabelas no banco de dados com base nas estruturas definidas
	db.AutoMigrate(&Products{}, &Category{}, &SerialNumber{})

	//Criar uma nova categoria
	// category := Category{Name: "POS"}
	// db.Create(&category)
	// category2 := Category{Name: "Gertec"}
	// db.Create(&category2)
	// // Criar um novo produto
	// product := Products{Name: "GPOS700X", Price: 999.99, Categories: []Category{category, category2}}
	// db.Create(&product)

	// var categories []Category
	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	// if err != nil {
	// 	panic(err)
	// }
	// for _, category := range categories {
	// 	fmt.Printf("Category Name: %s\n", category.Name)
	// 	for _, product := range category.Products {
	// 		fmt.Printf("Product Name: %s, Price: %.2f\n", product.Name, product.Price)
	// 	}
	// }

	// // Criar um novo número de série
	// serialNumber := SerialNumber{Number: "SN654321", ProductID: 2}
	// db.Create(&serialNumber)

	// var products []Products
	// Listar todos os produtos com suas categorias
	// Belongs to é usado para carregar as categorias associadas aos produto
	// Preload é usado para carregar as associações de forma antecipada
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Printf("ID: %d, Name: %s, Price: %.2f, Category: %s, Serial: %s\n", product.ID, product.Name, product.Price, product.Category.Name, product.SerialNumber.Number)
	// }

	// Listar todas as categorias com seus produtos e números de série
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Printf("Category Name: %s\n", category.Name)
		for _, product := range category.Products {
			fmt.Printf("Product Name: %s, Price: %.2f\n", product.Name, product.Price)
		}
	}
	// Criar um lote de produtos
	// products := []Products{
	// 	{Name: "NoteBook", Price: 1999.99, CategoryID: category.ID},
	// 	{Name: "Smartphone", Price: 999.99, CategoryID: category.ID},
	// 	{Name: "Tablet", Price: 499.99, CategoryID: category.ID},
	// 	{Name: "Smartwatch", Price: 299.99, CategoryID: category.ID},
	// 	{Name: "Headphones", Price: 199.99, CategoryID: category.ID},
	// }
	// db.Create(&products)

	// Selecionar o primeiro produto
	// var product Products
	// db.First(&product, 3)
	// fmt.Print("Primeiro produto: ", product.Name, " - Preço: ", product.Price, "\n")

	// Selecionar um smartphone
	// var smartphone Products
	// db.Where("name = ?", "Smartphone").First(&smartphone)
	// fmt.Println("Smartphone encontrado:", smartphone.Name, " - Preço:", smartphone.Price)

	// selecionar todos os produtos
	// var products []Products
	// // Limitando a 2 produtos e pulando os primeiros 2
	// db.Limit(2).Offset(2).Find(&products)
	// fmt.Println("Lista de produtos:")
	// for _, p := range products {
	// 	fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	// }

	// Quando usamos o LIKE, o GORM converte automaticamente para o operador LIKE do SQL
	// var expensiveProducts []Products
	// db.Where("name LIKE ?", "%Book%").Find(&expensiveProducts)
	// fmt.Println("Produtos caros:")
	// for _, p := range expensiveProducts {
	// 	fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	// }

	// var p Products
	// db.First(&p, 2)
	// p.Name = "Smartphone2"
	// db.Save(&p)
	// fmt.Println("Produto atualizado:", p.Name, " - Preço:", p.Price)

	// db.Delete(&p, 2)
	// fmt.Println("Produto deletado:", p.Name)

	// var allProducts []Products
	// db.Find(&allProducts)
	// fmt.Println("Lista de todos os produtos:")
	// for _, product := range allProducts {
	// 	fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)
	// }

}
