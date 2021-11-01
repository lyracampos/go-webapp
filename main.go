package main

import (
	"go-webapp/routes"
	"net/http"
)

func main() {
	routes.Load()
	http.ListenAndServe(":8080", nil)
}
