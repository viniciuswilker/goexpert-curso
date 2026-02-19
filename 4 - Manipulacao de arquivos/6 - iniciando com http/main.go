package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	fmt.Println("Escutando na porta 8080")

	http.HandleFunc("/", BuscarCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscarCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("BAD REQUEST")
		return
	}

	cepBusca, erro := BuscaCep(cepParam)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Erro ao chamar a função")
		return
	}
	json.NewEncoder(w).Encode(cepBusca)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Olá mundo"))
}

func BuscaCep(cep string) (*ViaCEP, error) {

	resp, erro := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if erro != nil {
		fmt.Println("BAD REQUEST")
		return nil, erro
	}
	defer resp.Body.Close()

	corpoRequisicao, erro := io.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("Erro ao ler o corpo da requisicao")
		return nil, erro
	}

	var cepResultado ViaCEP

	if erro = json.Unmarshal(corpoRequisicao, &cepResultado); erro != nil {
		fmt.Println("Erro ao converter o valor")
		return nil, erro
	}

	return &cepResultado, nil

}
