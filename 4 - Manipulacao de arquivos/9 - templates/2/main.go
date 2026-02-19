package main

import (
	"html/template"
	"net/http"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	erro := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"java", 20},
		{"python", 100},
	})
	if erro != nil {
		panic(erro)
	}

	mux := http.NewServeMux()

	http.ListenAndServe(":8080", mux)
}
