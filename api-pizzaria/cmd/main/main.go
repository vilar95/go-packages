package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type Pizza struct {
	ID		  int     `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}

func main() {
	r := gin.Default()
	r.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, []Pizza{
			{ID: 1, Name: "Marguerita", Price: 8.99},
			{ID: 2, Name: "Pepperoni", Price: 9.99},
			{ID: 3, Name: "Mussarela", Price: 10.99},
		})
	})

	r.Run(":8080")
	
	fmt.Println("Servidor rodando na porta 8080")

}