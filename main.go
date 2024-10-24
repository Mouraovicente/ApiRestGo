package main

import (
	"ApiRestGo/routes"
	"net/http"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {

	routes.ListRoutes()
	http.ListenAndServe(":8080", nil)

}
