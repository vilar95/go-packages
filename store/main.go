package main

import (
	"net/http"
	"store/routes"
)

func main() {
	routes.SetupRoutes()
	println("Servidor rodando em http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
