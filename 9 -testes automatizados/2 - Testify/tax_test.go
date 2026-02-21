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
