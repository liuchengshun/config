package config

import "fmt"

type filelineError struct {
	line string
}

func newFilelineError(line string) filelineError {
	return filelineError{line: line}
}

func (err filelineError) error() string {
	return fmt.Sprintf("does not support the file line config meesage: %v", err.line)
}
