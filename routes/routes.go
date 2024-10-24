package routes

import (
	"ApiRestGo/controllers"
	"net/http"
)

func ListRoutes() {
	http.HandleFunc("/", controllers.Index)

}
