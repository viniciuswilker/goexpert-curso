package main

import (
	"1-api/configs"
	"1-api/internal/entity"
	"1-api/internal/infra/database"
	"1-api/internal/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.CarregarConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/products", productHandler.GetProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	http.ListenAndServe(":8000", r)
}
