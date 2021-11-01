package routes

import (
	"go-webapp/controllers"
	"net/http"
)

func Load() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
