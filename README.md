
# GoAsync Library

## Introduction
Elevate your Go applications with our async/await and promises library—streamlining concurrency with an intuitive API for maximum performance and readability. Dive into next-level Go programming; your journey to efficient, clean code starts here.

## Overview
The GoAsync library is a concise, powerful package designed to simplify the execution of asynchronous tasks in Go applications. Drawing inspiration from the promise-based patterns in JavaScript, it offers a straightforward way to run async functions with support for timeouts, error handling, and concurrent execution of multiple tasks.

## Features
- **Asynchronous Function Execution**: Run asynchronous functions that return a result and an error, similar to JavaScript promises.
- **Timeouts**: Execute async functions with customizable timeouts to prevent hanging operations.
- **Concurrent Execution**: Run multiple async functions concurrently, with control over timeouts and error collection.
- **Simple API**: The library provides a simple and intuitive API, making it easy to integrate asynchronous operations into your Go applications.

## Installation

To use the GoAsync library, first, ensure you have a working Go environment. Then, install the library using `go get`:

```
go get github.com/angelospillos/goasync/async
```

## Usage

### Importing the Package

```go
import (
	"github.com/angelospillos/goasync/async"
)
```

### Running a Single Asynchronous Function

```go
asyncFunc := func() (interface{}, error) {
	time.Sleep(2 * time.Second) // Simulate work
	return "Result", nil
}

result, err := async.RunAsync(asyncFunc)
fmt.Printf("RunAsync: result = %v, error = %v
", result, err)
```

### Running a Single Asynchronous Function with Custom Timeout

```go
result, err := async.RunAsync(asyncFunc, async.WithTimeout(1*time.Second))
fmt.Printf("RunAsync with custom timeout: result = %v, error = %v
", result, err)
```

### Running Multiple Asynchronous Functions Concurrently

```go
results, errs := async.RunAllAsync([]async.AsyncFunc{asyncFunc, asyncFunc})
for i, res := range results {
	fmt.Printf("RunAllAsync: result[%d] = %v, error = %v
", i, res, errs[i])
}
```

### Running Multiple Asynchronous Functions with Custom Timeout

```go
results, errs := async.RunAllAsync([]async.AsyncFunc{asyncFunc, asyncFunc}, async.WithTimeout(5*time.Second))
for i, res := range results {
	fmt.Printf("RunAllAsync with custom timeout: result[%d] = %v, error = %v
", i, res, errs[i])
}
```

## API Reference

### AsyncFunc

Type definition for an asynchronous function.

```go
type AsyncFunc func() (interface{}, error)
```

### RunAsync

Executes a single asynchronous function with a default or custom timeout.

```go
func RunAsync(asyncFunc AsyncFunc, opts ...Option) (interface{}, error)
```

### RunAllAsync

Executes multiple asynchronous functions concurrently, with context control and customizable worker pool size.

```go
func RunAllAsync(asyncFuncs []AsyncFunc, opts ...Option) ([]interface{}, []error)
```

### Option and WithTimeout

Customizable options for asynchronous operations, such as setting a timeout.

```go
type Option func(*options)

func WithTimeout(t time.Duration) Option
```

## Examples

Refer to the `example.go` file for practical examples on how to use the GoAsync library in your projects.

## Contributing

Contributions to the GoAsync library are welcome! Please refer to the project's GitHub repository for contribution guidelines.

## License

The GoAsync library is open-source software licensed under the [GNU GPLv3](LICENSE.md).
