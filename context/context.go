package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("fundamentals: context")

	// Create a context
	ctx := context.Background()

	// Add a value to context
	ctx = context.WithValue(ctx, "name", "eric")

	// Get a value from context
	name := ctx.Value("name")
	fmt.Println(name)

	// Add a cancel to context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	// Send the context to a routine that we want to be cancellable
	go doSomething(ctx)

	// Simulate an event that cancels after a delay
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	// Block until the context is cancelled
	select {
	case <-ctx.Done():
		fmt.Println("context: cancel()")
	}

	// Cancel a context after a timeout
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)

	// Block until the context is cancelled
	select {
	case <-ctx.Done():
		fmt.Println("context: cancel()")
	}

}

func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("doSomething: cancel()")
			return
		default:
			fmt.Println("hello")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
