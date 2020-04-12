package docker

import (
	"github.com/dop251/goja"
)

func (d *Module) Inject(vm *goja.Runtime) {
	dockerObj := map[string]interface{} {
		"Build": d.Build,
		"Push": d.Push,
		"Network": map[string]interface{} {
			"Create": d.NetworkCreate,
		},
	}
	vm.Set("docker", dockerObj)
}

