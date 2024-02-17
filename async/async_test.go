package async

import (
	"errors"
	"testing"
	"time"
)

func TestRunAsync(t *testing.T) {
	// Test case for success scenario
	successFunc := func() (interface{}, error) {
		time.Sleep(1 * time.Second)
		return "success", nil
	}

	result, err := RunAsync(successFunc, WithTimeout(2*time.Second))
	if err != nil || result != "success" {
		t.Errorf("Expected success, got %v, error: %v", result, err)
	}

	// Test case for timeout scenario
	timeoutFunc := func() (interface{}, error) {
		time.Sleep(2 * time.Second)
		return nil, nil
	}

	_, err = RunAsync(timeoutFunc, WithTimeout(1*time.Second))
	if err == nil || err.Error() != "operation timed out" {
		t.Errorf("Expected timeout error, got %v", err)
	}

	// Test case for error scenario
	errorFunc := func() (interface{}, error) {
		return nil, errors.New("error occurred")
	}

	_, err = RunAsync(errorFunc)
	if err == nil || err.Error() != "error occurred" {
		t.Errorf("Expected error 'error occurred', got %v", err)
	}
}

func TestRunAllAsync(t *testing.T) {
	asyncFuncSuccess := func() (interface{}, error) {
		time.Sleep(1 * time.Second)
		return "success", nil
	}
	asyncFuncError := func() (interface{}, error) {
		return nil, errors.New("error occurred")
	}
	asyncFuncTimeout := func() (interface{}, error) {
		time.Sleep(3 * time.Second)
		return nil, nil
	}

	// Test case for all success scenario
	results, errs := RunAllAsync([]AsyncFunc{asyncFuncSuccess, asyncFuncSuccess}, WithTimeout(2*time.Second))
	for i, err := range errs {
		if err != nil {
			t.Errorf("Expected no error for func %d, got %v", i, err)
		}
	}
	for i, result := range results {
		if result != "success" {
			t.Errorf("Expected success for func %d, got %v", i, result)
		}
	}

	// Test case for mixed scenarios: success, error, and timeout
	results, errs = RunAllAsync([]AsyncFunc{asyncFuncSuccess, asyncFuncError, asyncFuncTimeout}, WithTimeout(2*time.Second))
	if errs[1].Error() != "error occurred" {
		t.Errorf("Expected error 'error occurred' for func 1, got %v", errs[1])
	}
	if errs[2] == nil || errs[2].Error() != "operation timed out" {
		t.Errorf("Expected timeout error for func 2, got %v", errs[2])
	}
}
