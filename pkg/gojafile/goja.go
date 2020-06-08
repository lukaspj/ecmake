package gojafile

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func expandFileName(path string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return path, errors.Wrap(err, fmt.Sprintf("failed to access %s", path))
	}
	if fi.Mode().IsDir() {
		path = filepath.Join(path, "ecmake.js")
	}

	fi, err = os.Stat(path)
	return path, errors.WithStack(err)
}

func GetGojaFile(vm *goja.Runtime, path string) (file *GojaFile, err error) {
	path, err = expandFileName(path)
	if err != nil {
		return nil, err
	}

	return NewGojaFile(vm, path), nil
}
