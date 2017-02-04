package main

import (
	"fmt"
	"time"
)

func greet() {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello")
}

func greet2() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello 2")
}

func main() {
	// "A goroutine is a function running independently in the same address space as other goroutines
	// Like launching a function with shell's & notation." - Rob Pike
	// https://talks.golang.org/2012/waza.slide
	// called synchronously
	greet()
	// called asynchronously
	go greet()
	go greet2()
	// greet 2 is invoked after greet but executed before
	// wait for enter
	// for while we needed this to avoid lost our goroutines
	// in a next day we learning how to wait for a computation in a better way
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
