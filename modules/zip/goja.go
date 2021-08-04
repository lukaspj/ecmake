package zip

import (
	"github.com/dop251/goja"
)

func Require() func(runtime *goja.Runtime, module *goja.Object) {
	return func(runtime *goja.Runtime, module *goja.Object) {
		zip := &Module{
			runtime: runtime,
		}

		obj := module.Get("exports").(*goja.Object)
		obj.Set("Writer", zip.Writer)
		obj.Set("WithWriter", zip.WithWriter)
	}
}
