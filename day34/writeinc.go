package main

/*
#include "mystring.h"

int sum(int a, int b){
	return a + b;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// sum two integers
	fmt.Println("sum(2, 3)", C.sum(2, 3))
	// CString wil return a *C.char
	// it's equivalent to char* in C
	original := C.CString("test")
	// length is included when i do #include in comments below
	fmt.Println("Length of \"test\"", C.length(original))
	fmt.Println("First index of \"t\" in \"test\"", C.firstIndexOf(original, C.char('t')))
	fmt.Println("First index of \"b\" in \"test\"", C.firstIndexOf(original, C.char('b')))
	fmt.Println("Last index of \"t\" in \"test\"", C.lastIndexOf(original, C.char('t')))
	fmt.Println("Last index of \"b\" in \"test\"", C.lastIndexOf(original, C.char('b')))
	house := C.CString("house")
	houte := C.CString("houte")
	fmt.Println(`"house" == "house"`, C.equals(house, house))
	fmt.Println(`"house" == "houte"`, C.equals(house, houte))
	C.toUpperCase(original)
	fmt.Println("toUpperCase(\"test\")", C.GoString(original))
	C.toLowerCase(original)
	fmt.Println("toLowerCase(\"TEST\")", C.GoString(original))
	testUpper := C.CString("TEST")
	tastUpper := C.CString("TAST")
	fmt.Println("\"test\" == \"TEST\" ignoring case", C.equalsIgnoreCase(original, testUpper))
	fmt.Println("\"test\" == \"TAST\" ignoring case", C.equalsIgnoreCase(original, tastUpper))
	sub := C.CString("")
	C.substring(original, sub, 1, 4)
	fmt.Println("teste (0, 4)", C.GoString(sub))
	C.replace(original, C.char('e'), C.char('a'))
	// Have to convert to a GoString
	fmt.Println("Replace \"e\" to \"a\" in \"test\"", C.GoString(original))
	// CString allocate memory using malloc
	// you have to free memory manually
	C.free(unsafe.Pointer(original))
	C.free(unsafe.Pointer(house))
	C.free(unsafe.Pointer(houte))
	C.free(unsafe.Pointer(sub))
	C.free(unsafe.Pointer(testUpper))
	C.free(unsafe.Pointer(tastUpper))
}
