package routes

import (
	"ApiRestGo/controllers"
	"net/http"
)

func ListRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)

}
