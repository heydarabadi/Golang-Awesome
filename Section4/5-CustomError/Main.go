package main

import (
	"errors"
	"time"
)

var errorNotFound error = errors.New("not found")
var errorAccessControll error = errors.New("access controller error")

func (op OpError) Error() string {
	return op.Error()
}

type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}

func DoSomething() error {
	return NewOpError("test", 200, "test Error", time.Now())
}

func NewOpError(op string, code int, message string, time time.Time) *OpError {
	return &OpError{
		Op:      op,
		Code:    code,
		Message: message,
		Time:    time,
	}
}

func main() {

}
