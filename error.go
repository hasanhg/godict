package godict

import "fmt"

const (
	errPrefix = "godict:"
)

var (
	ErrInvalidElement = errorf("invalid element")
)

func errorf(format string, a ...any) error {
	format = fmt.Sprintf("%s %s", errPrefix, format)
	return fmt.Errorf(format, a...)
}
