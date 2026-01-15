package main

import (
	"net/http"
	"store/routes"
)

func main() {
	routes.SetupRoutes()
	http.ListenAndServe(":8000", nil)
}
