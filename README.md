go-R
===

Go(golang) bindings for R language

This is simple binding to eval R expressions and pass results to/from Go code. 

**WARNING!**

Project in the early stage, memory leaks and even SIGFAULTs are possible. Use it on your own risk.

Known issues
===

* You should call `R.Init()` exactly in main goroutine of the app(i.e. not in goroutine created by your app). This was found by experiments. Possible reason is clashing between process thread stack and goroutine stack. As a result, its impossible to run R-related code in tests as go test run its testing function in custom goroutine per test.

Getting started
====

1. Install https://github.com/stretchrcom/testify testing package.
1. Install R environment: http://cran.r-project.org/
2. Check if you have R header files under correct path(check `#cgo CFLAGS:` directive in sources.
2. Make sure `libR.so` is avaliable. Set `$LD_LIBRARY_PATH` if R istallation is not system-wide.
2. Make sure `$R_HOME` is set and pointed to your R location before you run any R-related code.
3. Run `go run main.go` under go-R/example directory.

Basic usage
====

```
package main

import (
    "fmt"

    "github.com/nchern/go-R/R"
)

func main() {
    R.Init()

    x := R.NewNumericVector([]float64{1, 2, 3})
    R.SetSymbol("x", x)
    r := R.EvalOrDie("sum(x)").AsNumeric()
    fmt.Println(r.Get(0))
}
```

More examples are in test code.
