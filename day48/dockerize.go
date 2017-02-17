package main

import "fmt"

func main() {
	greeting := "Hello world"
	greet(greeting)
}

func greet(greetingArg string) {
	fmt.Println("This is my greeting:")
	fmt.Println(greetingArg)
}
