package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	products := []Product{
		{Name: "Camiseta", Description: "Bem bonita", Price: 29, Quantity: 10},
		{"Notebook", "Muito rápido", 1999, 2},
		{"Teste 1", "Descrição 1", 1, 1},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
