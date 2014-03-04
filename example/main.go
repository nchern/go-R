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

func test() {

	eval("1+3*4")

	R.EvalOrDie("library(stats)")

	x := R.NewNumericVector([]float64{11.0, 2.0, 31.0, 14.0, 51.0, 16.0, 7.0, 28.0})
	fmt.Println("x: ", x.ToArray())
	p := R.Protect(x.ToSexp())
	defer p.Unprotect()

	R.SetSymbol("x", x)
	eval("sum(x)")

}

func main() {

	fmt.Println("R.init: ", R.Init())
	test()
}
