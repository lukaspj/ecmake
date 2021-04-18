package io

import (
	"github.com/dop251/goja"
)

func Require(runtime *goja.Runtime, module *goja.Object) {
	ioObj := &Module{
		runtime: runtime,
	}

	obj := module.Get("exports").(*goja.Object)
	obj.Set("DeleteFile", ioObj.DeleteFile)
	obj.Set("DeleteFolder", ioObj.DeleteFolder)
}
