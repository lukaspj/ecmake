package gojafile

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/lukaspj/ecmake/pkg/buildfile"
	"github.com/pkg/errors"
	"io/ioutil"
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

func GetGojaFile(vm *goja.Runtime, path string) (file buildfile.BuildFile, err error) {
	path, err = expandFileName(path)
	if err != nil {
		return nil, err
	}

	return LoadGojaFile(vm, path)
}

func LoadGojaFile(vm *goja.Runtime, path string) (file *GojaFile, err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to read file %s", path))
	}

	return NewGojaFile(vm, path, string(content))
}
