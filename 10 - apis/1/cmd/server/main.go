package main

import (
	"1-api/configs"
	"1-api/internal/entity"
	"1-api/internal/infra/database"
	"1-api/internal/webserver/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.CarregarConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configs.JWTExperesIn))

	r.Route("/products", func(r chi.Router) {

		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)

	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}

// func LogRequest(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		r.Context().Value("user")
// 		log.Printf("Request: %s %s", r.Method, r.URL.Path)
// 		next.ServeHTTP(w, r)
// 	})
// }


// go install github.com/swaggo/swag/cmd/swag@latest