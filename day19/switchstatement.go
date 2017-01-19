package main

import (
	"fmt"
	"time"
)

func main() {
	// multiple conditionals
	option := 1
	switch option {
	// case 1 or 4
	case 1, 4:
		fmt.Println("option 1 or 4")

	case 2:
		// fallthrough is the same as "continue to next case without evaluate it"
		fmt.Printf("option 2 and ")
		fallthrough

	case 3:
		fmt.Println("option 3\n")
	}

	// switch statement "if style"
	t := time.Now().Day()
	switch {
	// if 1 < t <= 15
	case t > 1 && t <= 15:
		fmt.Println("$$$")
	// if 15 < t <= 20
	case t > 15 && t <= 20:
		fmt.Println("$$$ has gone")
	// 20+
	default:
		fmt.Println("Waiting for the next month")

	}

	// switch type
	// Anonymous function that receive something that implements interface{}(all!)
	func(some_var interface{}) {
		// case about type
		switch some_var.(type) {
		// is int
		case int:
			fmt.Println("Some var is int")
		// is string
		case string:
			fmt.Println("Some var is string")
		// None of previous
		default:
			fmt.Println("Sorry, unknown type!")
		}
	}(true)
}
