package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func conectDB() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=root dbname=inventory host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectDB()
	lista, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}
	p := Produto{}

	for lista.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = lista.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
