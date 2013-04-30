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

	//eval("1+3*4")

	R.EvalOrDie("library(stats)")
	R.EvalOrDie("library(chron)")
	R.EvalOrDie("library(fume)")
	R.EvalOrDie("library(zyp)")
	x := R.NewNumericVector([]float64{11.0, 2.0, 31.0, 14.0, 51.0, 16.0, 7.0, 28.0})
	//x := R.NewNumericVector([]float64{1, 2, 3, 4, 5.0, 6.0, 7.0, 8.0})
	p := R.Protect(x.ToSexp())
	defer p.Unprotect()

	R.SetSymbol("x", x)
	//	eval("sum(x)")

	//R.EvalOrDie("print(mkTrend(c(1,2,3,4,5,6,7,8), 0.95))")
	//R.EvalOrDie("print(MannKendall(c(1,2,3,4,5,6,7,8)))")

	//	l := x.Len()
	//R.EvalOrDie("print(x)")
	//R.EvalOrDie(fmt.Sprintf("print(acf(rank(lm(x ~ I(1:8))$resid), lag.max = 7, plot = FALSE)$acf[-1])"))
	//R.EvalOrDie(fmt.Sprintf("print(rank(lm(x ~ I(1:8))$resid))"))
	//R.EvalOrDie(fmt.Sprintf("print(lm(c(1,2,3,4,5,6,7,8) ~ I(1:8))$residuals)"))
	//R.EvalOrDie(fmt.Sprintf("x<-c(1,2,3,8,5,6,7,1)"))
	R.EvalOrDie(fmt.Sprintf("x<-c(1.001,2.001,3,4,5,6,7,8)"))
	//R.EvalOrDie(fmt.Sprintf("print(lm(x ~ I(1:8)))"))
	//R.EvalOrDie(fmt.Sprintf("print(lm(x ~ I(1:8))$residuals)"))
	//R.EvalOrDie(fmt.Sprintf("print(acf(rank(lm(x ~ I(1:8))$resid), lag.max = 7, plot = FALSE)$acf[-1])"))
	//R.EvalOrDie(fmt.Sprintf("print(lm(c(1,2,3,4,5,6,7,8) ~ I(1:8))$residuals)"))
	//R.EvalOrDie(fmt.Sprintf("print(rank(lm(c(1,2,3,4,5,6,7,8) ~ I(1:8))$resid))"))
	//R.EvalOrDie(fmt.Sprintf("print(lm(x ~ I(1:8)))"))
	//R.EvalOrDie(fmt.Sprintf("print(lm(x ~ I(1:8))$residuals)"))
	//R.EvalOrDie(fmt.Sprintf("print(summary(lm(x ~ I(1:8))))"))
	//R.EvalOrDie(fmt.Sprintf("print(I(1:8))"))
	//R.EvalOrDie(fmt.Sprintf("print(list(name=123, 4, foo=c(3,8,9))$name)"))

	R.EvalOrDie("library(fume)")
	y := R.EvalOrDie("mkTrend(x)").AsGenericVector()
	fmt.Println(y.Get(3).AsNumeric().Get(0))

}

func bg(proc func()) chan int {
	c := make(chan int)
	go func() {
		proc()
		c <- 1
	}()
	return c
}

func main() {

	fmt.Println("R.init: ", R.Init())
	test()

	//	c := bg(test)
	//	<-c
	//v := R.EvalOrDie("mkTrend(x)").AsGenericVector()
	//fmt.Println("mkTrend: ", v.Len())
}
