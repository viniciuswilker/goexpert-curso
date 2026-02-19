package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NovoProduto(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, erro := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if erro != nil {
		panic(erro)
	}
	defer db.Close()

	produto := NovoProduto("Notebook", 1554.50)

	erro = inserirProduto(db, produto)
	if erro != nil {
		panic(erro)
	}

}

func inserirProduto(db *sql.DB, produto *Product) error {
	stmt, erro := db.Prepare("insert into products (id, name, price) 	values(?, ?, ?)")
	if erro != nil {
		return erro
	}

	defer stmt.Close()

	_, erro = stmt.Exec(produto.ID, produto.Name, produto.Price)
	if erro != nil {
		return erro
	}

	return nil
}
