package main

import (
	"errors"
	"fmt"
)

// ParamsByValue receives copies from parameters
func ParamsByValue(name string, age int) {
	// we can chage the value of the parameters,
	// but isn't reflected outside the function
	name = name + " Cash"
	age--
	fmt.Printf("Hi, my name is %s, I'm %d.\n", name, age)
}

// ParamsByReference receives a pointer to original parameter
func ParamsByReference(name *string, age *int) {
	// all changes are reflected outside the function
	*name = *name + " Cash"
	*age--
	fmt.Printf("Hi, my name is %s, I'm %d.\n", *name, *age)

}

// NamedReturn show how Go can use named returns
func NamedReturn(ok bool) (result int, err error) {
	// ok, result and err are varibles inside the function
	if !ok {
		// err receives an error, result was initialized with 0
		err = errors.New("An error occurs!")
		// result is 0 and err is an error
		// these values are returned
		return
	}
	// at this point result is 0 and err is nil
	result = 42
	// after this point result is 42 and err still nil
	// you don't need to specify returned values
	return
}

func main() {
	name := "Johnny"
	age := 43
	println("before", name, age)
	ParamsByValue(name, age)
	// outside the functions the values are the same
	println("after", name, age)
	println("before", name, age)
	ParamsByReference(&name, &age)
	println(name, age)
	println("after", name, age)

	r, err := NamedReturn(true)
	fmt.Printf("r = %+v\n", r)
	fmt.Printf("err = %+v\n", err)
}
