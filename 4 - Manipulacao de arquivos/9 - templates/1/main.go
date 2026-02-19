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

func main() {

	curso := Curso{"Go", 40}

	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}} -- Carga Horário: {{.CargaHoraria}}")
	erro := tmp.Execute(os.Stdout, curso)
	if erro != nil {
		panic(erro)
	}

	mux := http.NewServeMux()

	http.ListenAndServe(":8080", mux)
}
