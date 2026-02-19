package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))

		erro := t.Execute(w, Cursos{
			{"Go", 40},
			{"java", 20},
			{"python", 100},
		})
		if erro != nil {
			panic(erro)
		}
	})

	http.ListenAndServe(":8080", nil)
}
