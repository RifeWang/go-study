package main

import (
	"bytes"
	"debug"
	"fmt"
	"net/http"
	"sync"
)

// GenericFunction ... Alias for a function that takes no args and returns an error
type GenericFunction func() error

// RunAsyncAllowErrors ...
func RunAsyncAllowErrors(functions ...GenericFunction) []error {
	// Channel for communicating when execution has finished.
	finished := make(chan struct{}, 1)

	// Create a wait group for the functions to be executed.
	// Note: wg.Add sets the number of goroutines to wait on.
	//       wg.Done decrements the number of goroutines left to wait on by one.
	var wg sync.WaitGroup
	wg.Add(len(functions))

	// Wait group in go routine. Will not close until all functions have completed
	go func() {
		wg.Wait()
		close(finished)
	}()

	// Spawn a goroutine for each function
	errors := make([]error, len(functions))
	for j := range functions {
		go func(j int) {
			defer wg.Done()
			// Clause for handling panic errors
			defer func() {
				if r := recover(); r != nil {
					// Skip 4 stack frames:
					// 1) debug.Stack()
					// 2) formatStack()
					// 3) this anonymous func
					// 4) runtime/panic
					err := fmt.Errorf(
						"panic in async function: %v\n%s",
						r, formatStack(4))
					errors[j] = err
				}
			}()
			// Execute function and log error at appropriate
			// position if applicable.
			if err := functions[j](); err != nil {
				errors[j] = err
				return
			}
		}(j)
	}

	// The code below this will not be reached until the wait group has finished
	<-finished
	return errors
}

// Return formatted stack trace, skipping "skip" leading stack frames
func formatStack(skip int) string {
	lines := bytes.Split(bytes.TrimSpace(debug.Stack()), []byte("\n"))
	formatted := bytes.Join(lines[1+2*skip:], []byte("\n"))
	return string(formatted)
}

// --------------------------------------------
// 使用示例
// --------------------------------------------
// func makeRequest() error {
// 	if _, err := http.Get("http://foo.com"); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func  main() {
// 	funcs := []GenericFunction{
// 		makeRequest,
// 		makeRequest,
// 		makeRequest,
// 	}
// 	RunAsyncAllowErrors(funcs...)
// }

// --------------------------------------------
// 使用闭包传递参数
// --------------------------------------------
// func makeRequest(url string) GenericFunction {
// 	return func() error {
// 		if _, err := http.Get(url); err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }

// func main() {
// 	funcs := []GenericFunction{
// 		makeRequest("http://abc.com"),
// 		makeRequest("http://def.com"),
// 		makeRequest("http://ghi.com"),
// 	}
// 	RunAsyncAllowErrors(funcs...)
// }

// --------------------------------------------
// 使用闭包传递参数，同时接受返回结果
// --------------------------------------------
func makeRequest(url string, response *http.Response) GenericFunction {
	return func() error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		response = resp
		return nil
	}
}
func main() {
	var resp1 http.Response
	var resp2 http.Response
	var resp3 http.Response
	funcs := []GenericFunction{
		makeRequest("http://abc.com", &resp1),
		makeRequest("http://def.com", &resp2),
		makeRequest("http://ghi.com", &resp3),
	}
	RunAsyncAllowErrors(funcs...)
}
