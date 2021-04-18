package sh

import (
	"github.com/dop251/goja"
)

func Require(verbose bool) func(runtime *goja.Runtime, module *goja.Object) {
	return func(runtime *goja.Runtime, module *goja.Object) {
		sh := &Module{
			Verbose: verbose,
			runtime: runtime,
		}

		obj := module.Get("exports").(*goja.Object)
		obj.Set("RunCmd", sh.RunCmd)
		obj.Set("OutCmd", sh.OutCmd)
		obj.Set("Run", sh.Run)
		obj.Set("RunV", sh.RunV)
		obj.Set("RunWith", sh.RunWith)
		obj.Set("RunWithV", sh.RunWithV)
		obj.Set("Output", sh.Output)
		obj.Set("OutputWith", sh.OutputWith)
		obj.Set("Exec", sh.Exec)
		obj.Set("CmdRan", sh.CmdRan)
		obj.Set("ExitStatus", sh.ExitStatus)
	}
}
