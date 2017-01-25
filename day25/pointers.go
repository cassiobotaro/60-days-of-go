package main

func decr(number *int) {
	*number--
}

func main() {
	// simple integer
	number := 42
	// point to a number
	var numberPointer *int

	// while not initialized numberPointer is null
	if numberPointer == nil {
		println("initialize null")
	}

	// assign the adress of number to numberPointer
	// "&" means "address of"
	numberPointer = &number

	// "*" means "value pointed by"
	println(*numberPointer)

	// without asterisk will print the adress pointed by numberPointer(address of number)
	println(numberPointer)

	// if changes the value of number
	number++

	// numberPointer value also changes(are the same value - memory address)
	println(number)
	println(*numberPointer)

	// pointer can also be used to change the value of a variable inside a function
	// pass "address of" number as parameter
	decr(&number)

	println(number)
}
