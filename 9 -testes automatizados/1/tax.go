package tax

import "time"

func CalcularTax(valor float64) float64 {

	if valor == 0 {
		return 0
	}

	if valor >= 1000 {
		return 10.0
	}

	return 5.0

}

func CalcularTax2(valor float64) float64 {
	time.Sleep(time.Microsecond)
	if valor == 0 {
		return 0
	}

	if valor >= 1000 {
		return 10.0
	}

	return 5.0

}
