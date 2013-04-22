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

SEXP GenericVectorElt(SEXP vec, int i) {
    return VECTOR_ELT(vec, i);
}
*/
import "C"

type Vector struct {
	expression
}

func NewVector(length int) *Vector {

	v := Vector{}
	v.expr = C.allocVector(C.REALSXP, C.R_xlen_t(length))
	v.length = length

	return &v
}

func (this *Vector) Get(i int) *Result {
	this.boundsCheck(i)
	C.Rf_protect(this.expr)
	defer C.Rf_unprotect(1)
	return NewResult(C.GenericVectorElt(this.expr, C.int(i)))
}
