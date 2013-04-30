#include <stdio.h>
#include <stdlib.h>

#define CSTACK_DEFNS 1

#include <R.h>
#include <Rinternals.h>
#include <Rinterface.h>
#include <Rdefines.h>
#include <Rembedded.h>
#include <R_ext/Parse.h>


int initR() {
    char *argv[] = {"REmbeddedMy", "--gui=none", "--silent"};
    int argc = sizeof(argv)/sizeof(argv[0]);

    return Rf_initEmbeddedR(argc, argv);
}

SEXP execCmd(const char *cmd) {
    SEXP cmdSexp, cmdExpr, ans = R_NilValue;
    ParseStatus status;
    int i;

    PROTECT(cmdSexp = allocVector(STRSXP, 1));
    SET_STRING_ELT(cmdSexp, 0, mkChar(cmd));
    cmdExpr = PROTECT(R_ParseVector(cmdSexp, -1, &status, R_NilValue));
    if (status != PARSE_OK) {
        UNPROTECT(2);
        error("invalid command: %s", cmd);
    }
    /* Loop is needed here as EXPSEXP will be of length > 1 */
    for(i = 0; i < length(cmdExpr); i++)
         ans = eval(VECTOR_ELT(cmdExpr, i), R_GlobalEnv);
    UNPROTECT(2);
    return ans;
}

int execScript(int line_count, char *lines[]) {
    int i;
    for (i = 0; i < line_count; i++) {
        execCmd(lines[i]);
    } 
    return 0;
}

void lmExample() {
        execCmd("library(fume)");
        execCmd("print(mkTrend(c(1,2,3,4,5,6,7,8), 0.95))");
	    execCmd("print(lm(x ~ I(1:8))$resid)");
        //execCmd("print(mkTrend(x, 0.95))");
}

int main(int argc, char **argv) {
    int r = initR();
    printf("Hello R: %d\n", r);
    SEXP e, val;
    int errorOccurred, i;
    int result = -1;
    
    SEXP ans, x, y;

    PROTECT(x = allocVector(REALSXP, 8));
    for (i = 0; i < 8; i++)
        REAL(x)[i] = i + 1;
    SEXP xSym = install("x");
    SEXP ySym = install("y");
    defineVar(xSym, x, R_GlobalEnv);
    PROTECT(xSym);
    printf("x: %d, y: %d\n", findVar(xSym, R_GlobalEnv) ==  R_UnboundValue, findVar(ySym, R_GlobalEnv) == R_UnboundValue);
    //char *scridddpt[] = {
        //"c<-file(\"R.log\", \"w\")",
        //"sink(file=c, type=\"message\")",
        //"1+3",
    //};
    //execScript(4, script);
    UNPROTECT(2);
}
