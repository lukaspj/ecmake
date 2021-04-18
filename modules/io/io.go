package io

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/pkg/errors"
	"os"
)

type Module struct {
	runtime *goja.Runtime
}

func New() *Module {
	return &Module{}
}

func (io *Module) DeleteFile(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to access %s", path))
	}
	if fi.IsDir() {
		return &FileNotFoundError{path:path, message: "path pointed to a directory"}
	}
	return os.Remove(path)
}

func (io *Module) DeleteFolder(path string, force bool) error {
	fi, err := os.Stat(path)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to access %s", path))
	}
	if !fi.IsDir() {
		return &FileNotFoundError{path:path, message: "path pointed to a file"}
	}
	if force {
		return os.RemoveAll(path)
	} else  {
		return os.Remove(path)
	}
}