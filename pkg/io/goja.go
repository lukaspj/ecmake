package io

import (
	"github.com/dop251/goja"
)

func (io *Module) Inject(vm *goja.Runtime) {
	ioObj := map[string]interface{} {

	}
	vm.Set("io", ioObj)
}

