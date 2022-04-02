package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

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
	produtos := []Produto{
		{Nome: "Caneta", Descricao: "Caneta Bic Preta", Preco: 1.79, Quantidade: 10},
		{Nome: "Caderno", Descricao: "Caderno Escolar", Preco: 3.12, Quantidade: 25},
		{Nome: "Lapis", Descricao: "Lapis Preto", Preco: 1.79, Quantidade: 15},
		{Nome: "Borracha", Descricao: "Borracha Azul", Preco: 1.79, Quantidade: 5},
	}

	err := templates.ExecuteTemplate(w, "Index", produtos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
