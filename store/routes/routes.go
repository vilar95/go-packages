package routes

import (
	"net/http"
	"store/controllers"
)

func SetupRoutes() {
	http.HandleFunc("/", controllers.Index)
}
