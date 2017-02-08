package main

import (
	"fmt"
	"time"
)

func main() {
	// timer will block for 5 seconds
	timer1 := time.NewTimer(time.Second * 5)
	// on each second ticker will "tick"
	ticker := time.NewTicker(time.Second)
	// actual hour
	initial := time.Now()
	go func() {
		// iterate over "ticks" and print number of seconds passed
		for _ = range ticker.C {
			fmt.Printf("\r%d seconds passed...", int(time.Since(initial).Seconds()))
		}
	}()
	// block until timer happens
	<-timer1.C
	fmt.Print("Beep! Beep!")
}
