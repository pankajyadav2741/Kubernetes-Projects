package error

import "fmt"

type myError struct {
	StatusCode int
	Message    string
}

func HandleError(code int, msg string) myError {
	return myError{
		StatusCode: code,
		Message:    fmt.Sprintf("error: %w", msg),
	}
}
