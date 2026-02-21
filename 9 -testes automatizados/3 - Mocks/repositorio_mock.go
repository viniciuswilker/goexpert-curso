package tax

import "github.com/stretchr/testify/mock"

type TaxRepositorioMock struct {
	mock.Mock
}

func (m *TaxRepositorioMock) SaveTax(tax float64) error {
	args := m.Called(tax)

	return args.Error(0)
}
