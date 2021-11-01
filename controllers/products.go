package controllers

import (
	"go-webapp/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAll()
	temp.ExecuteTemplate(w, "Index", products)
}
