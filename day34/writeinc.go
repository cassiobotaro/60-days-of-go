package main

/*
#include "mystring.h"
*/
import "C"
import "fmt"

func main() {
	// CString wil return a *C.char
	// it's equivalent to char* in C
	original := C.CString("teste")
	// length is included when i do #include in comments below
	fmt.Println(C.length(C.CString("teste")))
	C.replace(original, C.char('e'), C.char('a'))
	// Have to convert to a GoString
	fmt.Println(C.GoString(original))
}
