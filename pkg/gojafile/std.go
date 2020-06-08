package gojafile

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/lukaspj/ecmake/pkg/buildfile"
	"strings"
)

type StdModule struct {
	logger hclog.Logger
	file   *GojaFile
}

func NewStdModule(logger hclog.Logger) *StdModule {
	return &StdModule{logger: logger}
}

type ExecutionError struct {
	Command  string
	Args     []string
	ExitCode int
}

func (t ExecutionError) Error() string {
	return fmt.Sprintf(`running "%s %s" failed with exit code %d`, t.Command, strings.Join(t.Args, " "), t.ExitCode)
}

func (std *StdModule) Inject(file *GojaFile) *StdModule {
	std.file = file

	stdObj := map[string]interface{}{
		"LoadPlugin": std.LoadPlugin,
	}
	file.runtime.Set("std", stdObj)
	return std
}

func (std *StdModule) LoadPlugin(path string) map[string]interface{} {
	std.logger.Info("Trying to load plugin", "path", path)
	plugin := buildfile.InitializeModule(std.logger, path)
	std.file.AddModuleHost(plugin)

	module := plugin.Dispense()
	std.logger.Info("Fetching methods")
	methods := module.GetMethods()

	obj := map[string]interface{}{}
	for _, m := range methods {
		obj[m] = getModuleCall(module, m)
	}

	return obj
}

func getModuleCall(module buildfile.Module, method string) func(args... interface{}) interface{} {
	return func(args... interface{}) interface{} {
		return module.Invoke(method, args)
	}
}
