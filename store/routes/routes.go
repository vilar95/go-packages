package routes

import (
	"net/http"
	"store/controllers"
)

func SetupRoutes() {
	// Static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/add-product", controllers.AddProduct)
	http.HandleFunc("/insert-product", controllers.InsertProduct)
	http.HandleFunc("/delete-product", controllers.DeleteProduct)
	http.HandleFunc("/edit-product", controllers.EditProduct)
	http.HandleFunc("/update-product", controllers.UpdateProduct)
}
