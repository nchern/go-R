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

Rcomplex ComplexVectorElt(SEXP vec, int i) {
    return COMPLEX(vec)[i];
}

void SetComplexVectorElt(SEXP vec, int i, Rcomplex val) {
    COMPLEX(vec)[i] = val;
}
*/
import "C"

type ComplexVector struct {
	expression
}

func NewComplexVector(vector []complex128) *ComplexVector {

	length := len(vector)
	v := ComplexVector{}
	//v.expr = C.allocVector(C.CPLXSXP, C.R_len_t(length))
	v.expr = C.allocVector(C.CPLXSXP, C.R_xlen_t(length))
	v.length = length

	v.CopyFrom(vector)

	return &v
}

func (this *ComplexVector) Get(i int) complex128 {
	this.boundsCheck(i)
	C.Rf_protect(this.expr)
	defer C.Rf_unprotect(1)

	c := C.ComplexVectorElt(this.expr, C.int(i))
	return complex(float64(c.r), float64(c.i))
}

func (this *ComplexVector) Set(i int, val complex128) {
	this.boundsCheck(i)
	C.Rf_protect(this.expr)
	defer C.Rf_unprotect(1)

	var c C.Rcomplex
	c.r = C.double(real(val))
	c.i = C.double(imag(val))
	C.SetComplexVectorElt(this.expr, C.int(i), c)
}

func (this *ComplexVector) CopyFrom(src []complex128) {
	C.Rf_protect(this.expr)
	defer C.Rf_unprotect(1)
	for i := 0; i < this.length; i++ {
		var c C.Rcomplex
		c.r = C.double(real(src[i]))
		c.i = C.double(imag(src[i]))
		C.SetComplexVectorElt(this.expr, C.int(i), c)
	}
}
