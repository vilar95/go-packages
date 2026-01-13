package routes

import (
	"net/http"
	"store/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/add-product", controllers.AddProduct)
	http.HandleFunc("/insert-product", controllers.InsertProduct)
	http.HandleFunc("/delete-product", controllers.DeleteProduct)
}
