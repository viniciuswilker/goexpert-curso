package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularTax(t *testing.T) {
	tax, err := CalcularTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalcularTax(0)
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, tax)
}

func TestCalcularTaxESalvar(t *testing.T) {

	repositorio := &TaxRepositorioMock{}
	repositorio.On("SaveTax", 10.0).Return(nil)
	// repositorio.On("SaveTax", 0).Return(errors.New("erro ao salvar tax"))
	erro := CalcularTaxESalvar(1000.0, repositorio)
	assert.Nil(t, erro)

	repositorio.AssertExpectations(t)
}
