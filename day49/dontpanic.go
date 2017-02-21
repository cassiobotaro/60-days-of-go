package main

import "fmt"

func main() {
	// if panic happens, recover and print the error
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panic: %+v\n", e)
		}
	}()
	// panic crash your program and if not recovered an ugly traceback
	// is returned
	// Avoid panic, try to handle errors gracefully
	// DON'T PANIC, keep calm and keep your towel(see day 28)
	panic("forget the towel")
}
