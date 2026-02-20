package main

import (
	"database/sql"
	"fmt"

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

	// produto := NovoProduto("Celular", 2000.50)

	// erro = inserirProduto(db, produto)
	// if erro != nil {
	// 	panic(erro)
	// }

	// produto.Price = 100.00
	// erro = atualizarProduto(db, produto)
	// if erro != nil {
	// 	panic(erro)
	// }

	// p, erro := buscarPorId(db, produto.ID)
	// if erro != nil {
	// 	panic(erro)
	// }
	// fmt.Printf("Produto: %v, possui o preço de R$ %.2f", p.Name, p.Price)

	produtos, erro := buscarTodosProdutos(db)
	if erro != nil {
		panic(erro)
	}

	for _, p := range produtos {
		fmt.Printf("Produto: %v, Price: %.2f\n", p.Name, p.Price)
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

func atualizarProduto(db *sql.DB, produto *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.Name, produto.Price, produto.ID)
	if err != nil {
		return err
	}

	return nil
}

func buscarPorId(db *sql.DB, id string) (*Product, error) {

	stmt, erro := db.Prepare("select id, name, price from products where id = ?")
	if erro != nil {
		return nil, erro
	}

	defer stmt.Close()

	var produto Product

	if erro = stmt.QueryRow(id).Scan(&produto.ID, &produto.Name, &produto.Price); erro != nil {
		return nil, erro
	}

	return &produto, nil

}

func buscarTodosProdutos(db *sql.DB) ([]Product, error) {
	linhas, erro := db.Query("select id, name, price from products")
	if erro != nil {
		return []Product{}, erro
	}

	defer linhas.Close()

	var produtos []Product

	for linhas.Next() {

		var p Product
		if erro := linhas.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
		); erro != nil {
			return nil, erro
		}

		produtos = append(produtos, p)

	}

	return produtos, nil
}

func deletarProduto(db *sql.DB, id string) error {
	stmt, erro := db.Prepare("delete from products where id = ?")

	if erro != nil {
		panic(erro)
	}

	defer stmt.Close()

	_, erro = stmt.Exec(id)
	if erro != nil {
		panic(erro)
	}
	return nil

}
