package async

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// AsyncFunc is a function type that returns a result and an error, resembling a Promise in JavaScript.
type AsyncFunc func() (interface{}, error)

// RunAsync executes a single asynchronous function with a default timeout.
func RunAsync(asyncFunc AsyncFunc, opts ...Option) (interface{}, error) {
	options := makeOptions(opts...)
	ctx, cancel := context.WithTimeout(context.Background(), options.timeout)
	defer cancel()

	resultChan := make(chan interface{}, 1)
	errorChan := make(chan error, 1)

	go func() {
		result, err := asyncFunc()
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errorChan:
		return nil, err
	case <-ctx.Done():
		return nil, fmt.Errorf("operation timed out")
	}
}

// RunAllAsync executes multiple asynchronous functions concurrently, with context control and worker pool size.
func RunAllAsync(asyncFuncs []AsyncFunc, opts ...Option) ([]interface{}, []error) {
	options := makeOptions(opts...)
	ctx, cancel := context.WithTimeout(context.Background(), options.timeout)
	defer cancel()

	var wg sync.WaitGroup
	results := make([]interface{}, len(asyncFuncs))
	errors := make([]error, len(asyncFuncs))

	for i, asyncFunc := range asyncFuncs {
		wg.Add(1)
		go func(i int, asyncFunc AsyncFunc) {
			defer wg.Done()
			result, err := asyncFunc()
			select {
			case <-ctx.Done():
				if err == nil { // Only set timeout error if no other error has occurred
					errors[i] = fmt.Errorf("operation timed out")
				} else {
					errors[i] = err // Preserve original error if occurred
				}
			default:
				results[i] = result
				errors[i] = err
			}
		}(i, asyncFunc)
	}

	wg.Wait()
	return results, errors
}

// Option configuration options for async functions.
type Option func(*options)

type options struct {
	timeout time.Duration
}

// WithTimeout sets the timeout for the async operation.
func WithTimeout(t time.Duration) Option {
	return func(opts *options) {
		opts.timeout = t
	}
}

func makeOptions(opts ...Option) options {
	// Default options
	defaultOptions := options{
		timeout: 10 * time.Second, // Default timeout
	}

	for _, opt := range opts {
		opt(&defaultOptions)
	}

	return defaultOptions
}
