package controllers

import (
	"ApiRestGo/models"
	"net/http"
	"text/template"
)

var tamp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SelectAll()
	tamp.ExecuteTemplate(w, "Index", allProducts)
}
