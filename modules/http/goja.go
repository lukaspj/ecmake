package http

import (
	"github.com/dop251/goja"
)

func Require(runtime *goja.Runtime, module *goja.Object) {
	h := &Module{
		runtime: runtime,
	}

	obj := module.Get("exports").(*goja.Object)
	obj.Set("Get", h.Get)
	obj.Set("Head", h.Head)
	obj.Set("Post", h.Post)
	obj.Set("Put", h.Put)
	obj.Set("Delete", h.Delete)
	obj.Set("Request", h.Request)
}
