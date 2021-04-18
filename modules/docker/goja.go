package docker

import (
	"github.com/dop251/goja"
)

func Require(runtime *goja.Runtime, module *goja.Object) {
	docker := &Module{
		runtime: runtime,
	}

	obj := module.Get("exports").(*goja.Object)
	obj.Set("Build", docker.Build)
	obj.Set("Push", docker.Push)
	obj.Set("NetworkCreate", docker.NetworkCreate)
}

