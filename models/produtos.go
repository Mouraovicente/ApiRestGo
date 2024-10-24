package models

import (
	"ApiRestGo/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func SelectAll() []Produto {
	db := db.ConnectDB()
	selectAll, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectAll.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func InsertProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()
	insertProduct, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}
	insertProduct.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

func DeleteProduct(id string) {
	db := db.ConnectDB()
	deleteProduto, err := db.Prepare("Delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}
	deleteProduto.Exec(id)
	defer db.Close()

}
