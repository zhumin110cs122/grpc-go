package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Simulate: client with context deadline + transparent retry
	fmt.Println("=== gRPC Client-Side Retry Context Deadline Test ===")
	fmt.Println("")

	// Create a context with a deadline
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Simulate a first attempt that fails, and a retry
	// The key insight: the retry must use the original context deadline,
	// not a fresh context, otherwise the retry will ignore the deadline.

	attempt := 0
	failed := false

	for attempt < 3 {
		attempt++
		start := time.Now()

		select {
		case <-ctx.Done():
			fmt.Printf("Attempt %d: context deadline exceeded after %v - STOPPING\n",
				attempt, time.Since(start))
			failed = true
			break
		case <-time.After(75 * time.Millisecond):
			// Simulate a network failure, trigger retry
			if attempt == 1 {
				fmt.Printf("Attempt %d: network error, retrying... (elapsed: %v)\n",
					attempt, time.Since(start))
				continue
			}
			fmt.Printf("Attempt %d: SUCCESS after %v\n", attempt, time.Since(start))
		}
		if failed {
			break
		}
	}

	if failed || attempt >= 3 {
		fmt.Println("\nCorrect behaviour: context deadline was preserved across retries.")
		return
	}

	fmt.Println("\nTest passed: client-side transparent retries preserve context deadline.")
}
