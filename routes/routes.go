package routes

import (
	"inventory/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.Cadastrar)
	http.HandleFunc("/insert", controllers.Inserir)
	http.HandleFunc("/delete", controllers.Deletar)
	http.HandleFunc("/edit", controllers.Editar)
	http.HandleFunc("/update", controllers.Atualizar)
}
