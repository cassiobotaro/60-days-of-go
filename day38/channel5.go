package main

import "fmt"

func numbers() chan int {
	// create a channel inside function
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// sends i to c
			c <- i
		}
		// when finnish close the channel to indicate that ends
		close(c)
	}()
	// run the goroutine returns the channel needed to comunicate with goroutine
	return c
}

func main() {
	// iterate over a channel like something iterable
	for n := range numbers() {
		fmt.Println(n)
	}

	// 5 slots buffer
	tasks := make(chan int, 5)
	// insert 4 numbers
	tasks <- 1
	tasks <- 2
	tasks <- 3
	tasks <- 4
	// finnished before channel is full
	close(tasks)
	fmt.Println("Tasks done!")

	// select with default option are non blocking
	status := make(chan string)
	select {
	case msg := <-status:
		fmt.Println("finnished with: ", msg)
	default:
		fmt.Println("still running")
	}
}
