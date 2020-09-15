// +build e2e

package e2e

import "fmt"

type TestError struct {
	Message string
	Err     error
}

func newTestError(message string, err error) *TestError {
	return &TestError{
		Message: message,
		Err:     err,
	}
}

func (e *TestError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}
