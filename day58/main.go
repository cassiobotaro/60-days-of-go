package main

import (
	"context"
	"fmt"
	"time"
)

// FindKeyInContext receives a context and retrieves a key
func FindKeyInContext(ctx context.Context, k string) {
	if value := ctx.Value(k); value != nil {
		fmt.Println(value)
		return
	}
	fmt.Println("not found")
}

// Custom type needed to be key in context
type Custom string

func main() {
	// initialize the context
	ctx := context.WithValue(context.Background(), Custom("key"), "value")
	// Verify if context contains key "key"
	FindKeyInContext(ctx, "key")
	// Verify if context contains key "wrongkey"
	FindKeyInContext(ctx, "wrongkey")
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	i := 0
	running := true
	// run until deadline exceed
	for running {
		select {
		case <-ctx.Done():
			fmt.Println("deadline:", ctx.Err())
			running = false
		default:
			// print number of loops
			fmt.Println(i)
		}
		i++
	}
	// Another type of context is WithTimeout
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	// uncomment line 49 and comment line 50 if you want to see timeout
	// ctx, cancel = context.WithTimeout(context.Background(), 50*time.Millisecond)
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	exp := func() chan bool {
		c := make(chan bool)
		// simulate an expensive computation
		go func() {
			time.Sleep(2 * time.Second)
			c <- true
		}()
		return c
	}

	select {
	case <-exp():
		fmt.Println("Finnish him!")
	case <-ctx.Done():
		fmt.Println("timeout: ", ctx.Err()) // prints "context deadline exceeded"
	}

	// For more about context, see the docs: https://golang.org/pkg/context
}
