package console

import (
	"fmt"
	"github.com/dop251/goja"
)
// Inspired by github.com/dop251/goja_nodejs/console
type Console struct {
}

func (c *Console) log(call goja.FunctionCall) goja.Value {
	if len(call.Arguments) == 0 {
		fmt.Println()
	}

	fmtStr := call.Arguments[0].String()
	var args []interface{}
	for _, s := range call.Arguments[1:] {
		args = append(args, s.String())
	}

	fmt.Println(fmt.Sprintf(fmtStr, args...))
	return nil
}

func (c *Console) Inject(vm *goja.Runtime) {
	consoleObj := map[string]interface{} {
		"log": c.log,
		"error": c.log,
		"warn": c.log,
	}
	vm.Set("console", consoleObj)
}

func NewConsole() *Console {
	return &Console{}
}
