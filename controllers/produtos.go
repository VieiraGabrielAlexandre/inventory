package controllers

import (
	"inventory/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func Cadastrar(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Cadastrar", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preco: ", err)
			return
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade: ", err)
		}

		models.Inserir(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

func Deletar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.Deletar(id)
	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.Editar(id)
	templates.ExecuteTemplate(w, "Editar", produto)
}

func Atualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preco: ", err)
			return
		}

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter id: ", err)
			return
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade: ", err)
			return
		}

		models.Atualizar(idConvertido, quantidadeConvertida, nome, descricao, precoConvertido)
	}

	http.Redirect(w, r, "/", 301)
}
