package gojafile

import "github.com/dop251/goja"

type Module interface {
	Inject(runtime *goja.Runtime)
}