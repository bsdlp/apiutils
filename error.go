package apiutils

import "net/http"

// Error is an error
type Error interface {
	StatusCode() int
	Error() string
}

type err struct {
	status int
	msg    string
}

func (e err) Error() string { return e.msg }

// StatusCode returns the recommended http status code for this error
func (e err) StatusCode() int { return e.status }

// NewError returns an Error
func NewError(statusCode int, msg string) Error {
	if msg == "" {
		msg = http.StatusText(statusCode)
	}
	return err{
		status: statusCode,
		msg:    msg,
	}
}
