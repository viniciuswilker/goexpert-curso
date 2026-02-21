package tax

import "errors"

func CalcularTax(valor float64) (float64, error) {

	if valor == 0 {
		return 0.0, errors.New("deve ser maior que zero")

	}

	if valor >= 1000 && valor < 20000 {
		return 10.0, nil
	}

	if valor >= 20000 {
		return 20.0, nil
	}

	return 5.0, nil

}

type Repositorio interface {
	SaveTax(valor float64) error
}

func CalcularTaxESalvar(valor float64, repositorio Repositorio) error {
	tax := CalcularTax2(valor)
	return repositorio.SaveTax(tax)

}

func CalcularTax2(valor float64) float64 {

	if valor == 0 {
		return 0.0

	}

	if valor >= 1000 && valor < 20000 {
		return 10.0
	}

	if valor >= 20000 {
		return 20.0
	}

	return 5.0

}
