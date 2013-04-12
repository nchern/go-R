package R

import (
	"testing"

	"github.com/stretchrcom/testify/assert"
)

func TestNumericVector(t *testing.T) {
	Init()
	data := []float64{10, 2, 51, 48}
	actual := NewNumericVector(data)
	assert.Equal(t, len(data), actual.Len())
	assert.Panics(t, func() {
		actual.boundsCheck(-1)
	})
	for i, x := range data {
		assert.Equal(t, x, actual.Get(i))
	}

	actual.Set(2, 31)
	assert.Equal(t, 31, actual.Get(2))
	assert.Panics(t, func() {
		actual.Set(5, 1)
	})
}

func TestComplexVector(t *testing.T) {
	Init()
	data := []complex128{complex(1, 2), complex(1.1, 4.5)}
	actual := NewComplexVector(data)
	assert.Equal(t, len(data), actual.Len())

	for i, x := range data {
		assert.Equal(t, x, actual.Get(i))
	}
	actual.Set(1, complex(10, 11.2))
	assert.Equal(t, complex(10, 11.2), actual.Get(1))
}
