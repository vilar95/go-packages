package controllers

import (
	"html/template"
	"net/http"
	"store/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	AllProduct := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", AllProduct)
}