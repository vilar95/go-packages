package main

import (
	"database/sql"
	"fmt"
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

	// p, err := selectProduct(db, "2013f5a6-1657-409a-878c-2229e02344c1")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Product found:", p)
	// if p != nil {
	// 	fmt.Printf("Name: %s, Price: %.2f\n", p.Name, p.Price)
	// } else {
	// 	fmt.Println("No product found with the given ID.")
	// }

	fmt.Println("Product inserted successfully. ID:", Product.ID)
	fmt.Println("Listing all products:")
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	Product.Price = 4799.99
	err = updateProduct(db, Product)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product updated successfully. ID:", Product.ID)
	fmt.Println("Listing all products:")
	products, err = selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

	// err = deleteProduct(db, "3bd4a5ff-1c4e-44fc-bc2f-478cb9dfd934")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Product deleted successfully. ID:", "3bd4a5ff-1c4e-44fc-bc2f-478cb9dfd934")
	// }

	err = deleteAllProducts(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("All products deleted successfully.")

	fmt.Println("Listing all products:")
	products, err = selectAllProducts(db)
	if err != nil {
		panic(err)
	} else if len(products) == 0 {
		fmt.Println("No products found.")	
		return
	}
	
	for _, p := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}

}

// A função insertProduct insere um novo produto no banco de dados
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

// A função updateProduct atualiza um produto existente no banco de dados
func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// A função selectProduct busca um produto pelo ID no banco de dados
func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	//err = stmt.QueryRowContext(nil, id).Scan(&p.ID, &p.Name, &p.Price)
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// A função selectAllProducts busca todos os produtos no banco de dados
func selectAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*Product
	// O Next é usado para iterar sobre as linhas retornadas pela consulta
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		// O append é usado para adicionar o produto à lista de produtos
		products = append(products, &p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

// A função deleteProduct remove um produto do banco de dados pelo ID
func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// O Exec é usado para executar a instrução preparada com o ID do produto
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// A função deleteAllProducts remove todos os produtos do banco de dados
func deleteAllProducts(db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM products")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
