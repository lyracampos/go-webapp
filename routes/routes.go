package routes

import (
	"go-webapp/controllers"
	"net/http"
)

func Load() {
	http.HandleFunc("/", controllers.Index)
}
