package sh

import (
	"github.com/dop251/goja"
)

func (sh *Module) Inject(vm *goja.Runtime) {
	shObj := map[string]interface{} {
		"RunCmd": sh.RunCmd,
		"OutCmd": sh.OutCmd,
		"Run": sh.Run,
		"RunV": sh.RunV,
		"RunWith": sh.RunWith,
		"RunWithV": sh.RunWithV,
		"Output": sh.Output,
		"OutputWith": sh.OutputWith,
		"Exec": sh.Exec,
		"CmdRan": sh.CmdRan,
		"ExitStatus": sh.ExitStatus,
	}
	vm.Set("sh", shObj)
}

