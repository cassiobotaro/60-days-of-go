package main

/*
#include "mystring.h"
*/
import "C"
import "fmt"

func main() {
	// CString wil return a *C.char
	// it's equivalent to char* in C
	original := C.CString("test")
	// length is included when i do #include in comments below
	fmt.Println("Length of \"test\"", C.length(original))
	fmt.Println("First index of \"t\" in \"test\"", C.firstIndexOf(original, C.char('t')))
	fmt.Println("First index of \"b\" in \"test\"", C.firstIndexOf(original, C.char('b')))
	fmt.Println("Last index of \"t\" in \"test\"", C.lastIndexOf(original, C.char('t')))
	fmt.Println("Last index of \"b\" in \"test\"", C.lastIndexOf(original, C.char('b')))
	fmt.Println(`"house" == "house"`, C.equals(C.CString("house"), C.CString("house")))
	fmt.Println(`"house" == "houte"`, C.equals(C.CString("house"), C.CString("houte")))
	C.toUpperCase(original)
	fmt.Println("toUpperCase(\"test\")", C.GoString(original))
	C.toLowerCase(original)
	fmt.Println("toLowerCase(\"TEST\")", C.GoString(original))
	fmt.Println("\"test\" == \"TEST\" ignoring case", C.equalsIgnoreCase(original, C.CString("TEST")))
	fmt.Println("\"test\" == \"TAST\" ignoring case", C.equalsIgnoreCase(original, C.CString("TAST")))
	sub := C.CString("")
	C.substring(original, sub, 1, 4)
	fmt.Println("teste (0, 4)", C.GoString(sub))
	C.replace(original, C.char('e'), C.char('a'))
	// Have to convert to a GoString
	fmt.Println("Replace \"e\" to \"a\" in \"test\"", C.GoString(original))
}
