package main

import (
	"errors"
	"fmt"
)

// SimpleSum sum two numbers
// receive two params and returns one
func SimpleSum(a, b int) int {
	return a + b
}

// MultipleReturn returns multiple values
// don't receive params but return multiple values
func MultipleReturn() (int, string) {
	return 0, "another value"
}

// Division will  raise error when b is 0
// it's a common practice returns error as second value
func Division(a, b float64) (result float64, err error) {
	if b == 0.0 {
		return 0.0, errors.New("division by zero")
	}
	return a / b, nil
}

// Sum can receive a variadic number of parameters
// Variadic params are a slice of that type
// sum can receive one or many params for exemple sum(1,2) or sum(1,2,3,4)
func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

// Params are copy of passed arguments
// except in some cases like interfaces, slices, maps
// if you want the real item, uses pointer
func referenceExample(a *int, b []int, c map[string]string) {
	*a++
	b = append(b, 4)
	c["some"] = "item"
}

func main() {

	// function call
	println(SimpleSum(2, 2))

	// call a function but ignore the results
	_ = SimpleSum(2, 2)

	// ignore the first value returned
	_, second := MultipleReturn()
	println(second)

	// variadic params - one or more
	println(sum(1, 2, 3, 4, 5))
	println(sum(1, 2))
	println(sum(1))
	number := 0
	list := []int{}
	amap := make(map[string]string)
	// the address of number is passed to function that change it value
	// then when we print, the value is 1
	// list is passed by copy, and is not changed by function
	// but maps are passed by reference, changes will be reflected
	// outside function
	referenceExample(&number, list, amap)
	fmt.Println(number, list, amap)
}
