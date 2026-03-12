package main

import (
	"fmt"

	"api-pizzaria/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/pizzas", GetPizzas)

	r.Run(":8080")

	fmt.Println("Servidor rodando na porta 8080")

}

func GetPizzas(c *gin.Context) {
	Pizzas := []model.Pizza{
		{ID: 1, Name: "Marguerita", Price: 8.99},
		{ID: 2, Name: "Pepperoni", Price: 9.99},
		{ID: 3, Name: "Mussarela", Price: 10.99},
	}

	c.JSON(200, Pizzas)
}
