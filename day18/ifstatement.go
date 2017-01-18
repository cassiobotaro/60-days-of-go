package main

import "fmt"

func main() {
	rainingToday := true
	cloudy := true

	// simple conditional
	if rainingToday {
		fmt.Println("It's raining today!")
	}
	// conditional with 2 conditions
	if rainingToday {
		fmt.Println("It's raining today")
	} else {
		fmt.Println("Go to the beach!")
	}
	// multiple conditions
	if rainingToday {
		fmt.Println("It's raining today")
	} else if cloudy {
		fmt.Println("It's cold today")
	} else {
		fmt.Println("Go to the beach!")
	}
	// conditional with enclosing scope
	number := 5
	if number := 6; true {
		number++
		fmt.Printf("number = %+v\n", number)
	}
	// still 5 because of scope
	fmt.Printf("number = %+v\n", number)

}
