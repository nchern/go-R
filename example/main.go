package main

import (
	"fmt"
	"github.com/nchern/go-R/R"
)

func eval(expression string) *R.Result {
	r := R.EvalOrDie(expression)
	fmt.Println(expression, "=", r.AsNumeric().Get(0))
	return r
}

func main() {
	fmt.Println("R.init: ", R.Init())

	eval("1+3*4")

	R.EvalOrDie("library(fume)")
	x := R.NewNumericVector([]float64{1, 2, 3, 4, 5, 6, 7, 8})
	R.SetSymbol("x", x)
	eval("sum(x)")

	R.EvalOrDie("print(mkTrend(x))")
	//v := R.EvalOrDie("mkTrend(x)").AsGenericVector()
	//fmt.Println("mkTrend: ", v.Len())
}
