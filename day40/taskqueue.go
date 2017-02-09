package main

import (
	"fmt"
	"time"
)

const MagicNumber = 200

// This example is a series of tests based on https://gobyexample.com/worker-pools
// Try change magic numbers, and see the results

func worker(id int, jobs <-chan int, results chan<- int) {
	// each worker block until a new job is received
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}

func main() {
	// queue with MagicNumber slots
	jobs := make(chan int, MagicNumber)
	// job results
	results := make(chan int, MagicNumber)
	// handle MagicNumber workers, waiting for job that comes from jobs channel
	for w := 1; w <= MagicNumber; w++ {
		// jobs is our input channel and results our output channel
		go worker(w, jobs, results)
	}
	// start to fill the jobs queue
	go func() {
		for j := 1; j <= 400; j++ {
			jobs <- j
		}
		// close jobs indicates that i don't have more jobs
		close(jobs)
	}()
	// start to print the results
	for a := 1; a <= MagicNumber; a++ {
		fmt.Println(<-results)
	}
}
