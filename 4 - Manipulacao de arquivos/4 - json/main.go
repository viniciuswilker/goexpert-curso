package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {

	conta := Conta{Numero: 1, Saldo: 100}

	res, err := json.Marshal(conta)
	if err != nil {
		println(err.Error())
	}

	println(string(res))

	if err := json.NewEncoder(os.Stdout).Encode(conta); err != nil {
		println(err.Error())
	}

	jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)
	var contaX Conta
	if erro := json.Unmarshal(jsonPuro, &contaX); erro != nil {
		println(err)
	}

	println(contaX.Numero)
	println(contaX.Saldo)

}
