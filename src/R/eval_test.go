package R

import (
	//"log"
	"math"
	"testing"

	"github.com/stretchrcom/testify/assert"
)

func TestEvalBadExpr(t *testing.T) {
	assert.Equal(t, 1, Init())
	r, err := Eval("x+")
	assert.Nil(t, r)
	assert.Error(t, err)

	assert.Panics(t, func() {
		EvalOrDie("+1*(")
	})
}

func TestEval(t *testing.T) {
	assert.Equal(t, 1, Init())
	r, err := Eval("1+3*4")
	assert.Nil(t, err)
	assert.True(t, r.IsNumeric())
	assert.False(t, r.IsComplex())
	assert.Panics(t, func() { r.AsComplex() })
	num := r.AsNumeric()
	assert.Equal(t, 1, num.Len())
	assert.Equal(t, 13, num.Get(0))

	r = EvalOrDie("sqrt(-2+0i)")
	assert.False(t, r.IsNumeric())
	assert.True(t, r.IsComplex())
	assert.Panics(t, func() { r.AsNumeric() })
	cpl := r.AsComplex()
	assert.Equal(t, 1, cpl.Len())
	assert.Equal(t, complex(0, math.Sqrt(2)), cpl.Get(0))
}

func TestEvalWithVariables(t *testing.T) {
	x := NewNumericVector([]float64{10, 13, -15})
	SetSymbol("x", x)
	r := EvalOrDie("sum(x)")
	assert.Equal(t, 8, r.AsNumeric().Get(0))
}

func TestScriptFFT(t *testing.T) {
	assert.Equal(t, 1, Init())
	EvalOrDie("library(stats)")

	data := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	v := NewNumericVector(data)
	SetSymbol("y", v)
	r, err := Eval("fft(y, inverse=FALSE)")
	assert.True(t, r.IsComplex())
	assert.Nil(t, err)
	z := r.AsComplex()
	assert.Equal(t, 8, z.Len())
	assert.Equal(t, complex(36, 0), z.Get(0))
}

func TestMkTrend(t *testing.T) {
	assert.Equal(t, 1, Init())

	Eval("library(stats)")
	Eval("library(chron)")
	Eval("library(fume)")
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	v := NewNumericVector(data)
	SetSymbol("a1", v)
	r, err := Eval("mkTrend(a1)")
	//r, err := Eval("mkTrend(a1)")
	assert.True(t, r.IsGenericVector())
	assert.Nil(t, err)

	//z := r.AsComplex()
	//assert.Equal(t, 8, z.Len())
	//assert.Equal(t, complex(36, 0), z.Get(0))
}
