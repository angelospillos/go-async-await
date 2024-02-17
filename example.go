package main

import (
	"github.com/angelospillos/goasync/async"
	"fmt"
	"time"
)

func main() {
	// Async function example
	asyncFunc := func() (interface{}, error) {
		time.Sleep(2 * time.Second) // Simulate work
		return "Result", nil
	}

	// Run a single async function with default settings
	result, err := async.RunAsync(asyncFunc)
	fmt.Printf("RunAsync: result = %v, error = %v\n", result, err)

	// Run a single async function with a custom timeout
	result, err = async.RunAsync(asyncFunc, async.WithTimeout(1*time.Second))
	fmt.Printf("RunAsync with custom timeout: result = %v, error = %v\n", result, err)

	// Run multiple async functions with default settings
	results, errs := async.RunAllAsync([]async.AsyncFunc{asyncFunc, asyncFunc})
	for i, res := range results {
		fmt.Printf("RunAllAsync: result[%d] = %v, error = %v\n", i, res, errs[i])
	}

	// Run multiple async functions with a custom timeout
	results, errs = async.RunAllAsync([]async.AsyncFunc{asyncFunc, asyncFunc}, async.WithTimeout(5*time.Second))
	for i, res := range results {
		fmt.Printf("RunAllAsync with custom timeout: result[%d] = %v, error = %v\n", i, res, errs[i])
	}
}
