package main

import "fmt"

func main() {
	// repeat iteration until the expression is true
	// C like loop
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %+v\n", i)
	}

	// omiting initialization and step
	// "while" style
	counter := 1
	for counter < 5 {
		fmt.Printf("counter = %+v\n", counter)
		counter++ // don't forget this step or will run forever
	}

	// omiting only step
	for counter := 1; counter < 5; {
		fmt.Printf("counter = %+v\n", counter)
		counter++ // don't forget this step or will run forever
	}

	// omiting only initialization
	count := 1
	for ; count < 5; count++ {
		fmt.Printf("count = %+v\n", count)
		count++ // don't forget this step or will run forever
	}

	// run forever unless that have a break
	for {
		fmt.Println("loop")
		break // get out of the loop
	}

	// only even numbers
	for n := 0; n <= 5; n++ {
		// if is odd, continue loop
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}

	// for each element -  arrays
	// NOTE: range always returns two elements
	for index, value := range [4]string{"apple", "banana", "orange", "lemon"} {
		fmt.Printf("indice: %d value: %q\n", index, value)
	}
	// slices and map iterations will showed in respective days
}
