package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// Sleep 2 seconds A detail, this instruction will block process
	// until finnish
	time.Sleep(2 * time.Second)

	// christmas
	christmas := time.Date(2017, 12, 25, 0, 0, 0, 0, time.UTC)
	println(christmas.Day())
	println(christmas.Year())

	// week day
	fmt.Println(christmas.Weekday())

	// duration until christmas
	fmt.Println("days until christmas: ", math.Trunc(time.Until(christmas).Hours()/24))

	// check if date already passed
	fmt.Println("already passed: ", christmas.Before(time.Now()))

	// date after 3 days
	now := time.Now()
	days := time.Duration(24 * 3)
	future := now.Add(days * time.Hour)
	fmt.Println(future)

	// timer will block for 5 seconds
	timer := time.NewTimer(time.Second * 5)
	// on each second ticker will "tick"
	ticker := time.NewTicker(time.Second)
	// actual hour
	initial := time.Now()
	go func() {
		// iterate over "ticks" and print number of seconds passed
		for _ = range ticker.C {
			fmt.Printf("\r try %d: Dormammu, i've come to bargain............☠", int(time.Since(initial).Seconds()))
		}
	}()
	// block until timer happens
	<-timer.C
	fmt.Printf("\rtry %d: Dormammu, i've come to bargain............OK", int(time.Since(initial).Seconds()))
	fmt.Print("\nThe earth is safe again.")

	// If you want to see more about dates, see the docs: https://golang.org/pkg/time/
	// here too: https://gobyexample.com/time
}
