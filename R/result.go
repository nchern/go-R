package R

/*

#cgo LDFLAGS: -lm -lR
#cgo CFLAGS: -I /usr/share/R/include/

#include <stdlib.h>
#include <R.h>
#include <Rinternals.h>
#include <Rdefines.h>
#include <R_ext/Parse.h>
#include <Rembedded.h>

*/
import "C"

type Result struct {
	expr C.SEXP
}

func NewResult(expr C.SEXP) *Result {
	return &Result{expr: expr}
}

func (this *Result) IsNumeric() bool {
	return C.Rf_isReal(this.expr) != 0
}

func (this *Result) IsComplex() bool {
	return C.Rf_isComplex(this.expr) != 0
}

func (this *Result) AsComplex() *ComplexVector {
	if !this.IsComplex() {
		panic("Not a complex vector")
	}
	v := ComplexVector{}
	v.length = int(C.Rf_length(this.expr))
	v.expr = this.expr
	return &v

}

func (this *Result) AsNumeric() *NumericVector {
	if !this.IsNumeric() {
		panic("Not a numeric vector")
	}
	v := NumericVector{}
	v.length = int(C.Rf_length(this.expr))
	v.expr = this.expr
	return &v
}

func (this *Result) IsGenericVector() bool {
	return C.Rf_isVector(this.expr) != 0
}

func (this *Result) AsGenericVector() *Vector {
	if !this.IsGenericVector() {
		panic("Not a complex vector")
	}
	v := Vector{}
	v.length = int(C.Rf_length(this.expr))
	v.expr = this.expr
	return &v

}
