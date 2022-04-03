package models

import "inventory/db"

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
	Id         int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectDB()
	lista, err := db.Query("select * from produtos order by id asc")

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
		p.Id = id

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func Inserir(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectDB()
	stmt, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(nome, descricao, preco, quantidade)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func Deletar(id string) {
	db := db.ConectDB()
	stmt, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func Editar(id string) Produto {
	db := db.ConectDB()

	produtoCapturado, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for produtoCapturado.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoCapturado.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
		produto.Id = id

	}

	defer db.Close()

	return produto
}

func Atualizar(id, quantidade int, nome, descricao string, preco float64) {
	db := db.ConectDB()

	AtualizaProduto, err := db.Prepare("update produtos set id = $1, quantidade = $2, nome = $3, descricao = $4 where preco = $5")
	if err != nil {
		panic(err.Error())
	}

	_, err = AtualizaProduto.Exec(id, quantidade, nome, descricao, preco)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
