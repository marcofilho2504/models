package models

import (
	"meuprojeto/db"

	_ "github.com/marcofilho2504/models.git"
)

type produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err)
	}

	p := produto{}
	produtos := []produto{}

	for selectDeTodosOsProdutos.Next() {
		err = selectDeTodosOsProdutos.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		if err != nil {
			panic(err)
		}
		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos

}
