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

int initR() {
    char *argv[] = {"REmbeddedMy", "--gui=none", "--silent", "--slave"};
    int argc = sizeof(argv)/sizeof(argv[0]);

    return Rf_initEmbeddedR(argc, argv);
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

var (
	isInitialized int
)

type RProtector interface {
	Protect(sexpr C.SEXP) RProtector
	Unprotect()
}

type protector struct {
	count int
}

func Protect(sexpr C.SEXP) RProtector {
	p := protector{}
	p.Protect(sexpr)
	return &p
}

func (this *protector) Protect(sexpr C.SEXP) RProtector {
	C.Rf_protect(sexpr)
	this.count++
	return this
}

func (this *protector) Unprotect() {
	C.Rf_unprotect(C.int(this.count))
	this.count = 0
}

func EvalOrDie(expression string) *Result {
	r, err := Eval(expression)
	if err != nil {
		panic(fmt.Sprintf("FAILED: %s", err))
	}
	return r
}

func Eval(expression string) (*Result, error) {

	var status C.ParseStatus

	cmd := C.CString(expression)
	defer C.free(unsafe.Pointer(cmd))

	cmdRChar := C.mkChar(cmd)
	protector := Protect(cmdRChar)

	cmdSexp := C.allocVector(C.STRSXP, 1)
	protector.Protect(cmdSexp)
	C.SET_STRING_ELT(cmdSexp, 0, cmdRChar)

	parsedCmd := C.R_ParseVector(cmdSexp, -1, (*C.ParseStatus)(unsafe.Pointer(&status)), C.R_NilValue)
	if status != C.PARSE_OK {
		return nil, fmt.Errorf("Invalid command: %s", C.GoString(cmd))
	}

	protector.Protect(parsedCmd)
	defer protector.Unprotect()

	var result C.SEXP
	/* Loop is needed here as EXPSEXP will be of length > 1 */
	for i := 0; i < int(C.Rf_length(parsedCmd)); i++ {
		result = C.eval(C.VECTOR_ELT(parsedCmd, C.int(i)), C.R_GlobalEnv)
	}
	return NewResult(result), nil
}

func SetSymbol(name string, val Expression) {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))
	C.Rf_protect(val.toSexp())
	defer C.Rf_unprotect(1)
	C.defineVar(C.install(nameC), val.toSexp(), C.R_GlobalEnv)
}

func Init() int {
	if isInitialized != 0 {
		return isInitialized
	}
	isInitialized = int(C.initR())
	return isInitialized
}

/*
func testMem(n int) {
	data := randVector(50000)
	v := NewNumericVector(data)
	for i := 0; i < n; i++ {
		data = randVector(50000)
		v.CopyFrom(data)
		SetSymbol("x", v)
		res, err := Eval("sum(x)")
		if err != nil {
			panic(fmt.Sprintf("Eval error: %s", err))
		}
		if i%10 == 0 {
			fmt.Printf("%d Result: %f\nEnter to continue...\n", i, float64(C.asReal(res.expr)))
			var s string
			fmt.Scanf("%s", &s)
			if s == "q" {
				return
			}
			if s == "gc" {
				fmt.Printf("Call gc()")
				_, err = Eval("gc()")
				if err != nil {
					panic(fmt.Sprintf("Eval error: %s", err))
				}
			}
		}
	}
}*/
