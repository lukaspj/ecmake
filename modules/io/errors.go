package io

import (
	"fmt"
)

type FileNotFoundError struct {
	path    string
	message string
}

func (t *FileNotFoundError) Error() string {
	return fmt.Sprintf(`File: "%s" was not found: %s`, t.path, t.message)
}
