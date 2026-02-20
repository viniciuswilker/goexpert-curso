package tax

import "testing"

func TestCalcularTax(t *testing.T) {
	valor := 500.0

	esperado := 5.0

	resultado := CalcularTax(valor)

	if resultado != esperado {
		t.Errorf("Esperado %f mas retornou: %f", esperado, resultado)
	}
}

func TestCalculateTaxBatch(t *testing.T) {

	type calcTax struct {
		valor, esperado float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0, 0},
	}

	for _, item := range table {

		resultado := CalcularTax(item.valor)

		if resultado != item.esperado {
			t.Errorf("Esperado %f mas retornou: %f", item.esperado, resultado)

		}
	}

}

func BenchmarkCalculcarTaxt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalcularTax(500.0)
	}

}

func BenchmarkCalculcarTaxt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalcularTax2(500.0)
	}

}

func FuzzCalcularTax(f *testing.F) {

	seed := []float64{-1, -2, -3.5, 500.0, 1000.0, 1545.0}

	for _, valor := range seed {
		f.Add(valor)
	}
	f.Fuzz(func(t *testing.T, valor float64) {
		CalcularTax(valor)
	})

}

//  go test -coverprofile=coverage.out
//  go test -v
//  go tool cover -html=coverage.out
//  go test -bench=. -run=^#
