package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Produto struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {

	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, erro := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if erro != nil {
		panic(erro)
	}
	db.AutoMigrate(Produto{})

	// create
	// db.Create(&Produto{
	// 	Name:  "Geladeria",
	// 	Price: 2000,
	// })

	// buscar

	// var produto Produto
	// db.First(&produto, 1)
	// fmt.Println(produto)

	// db.First(&produto, "name = ?", "Geladeria")
	// fmt.Println(produto)

	// BUSCAR TODOS
	var produtos []Produto
	db.Find(&produtos)
	fmt.Println(produtos)

}
