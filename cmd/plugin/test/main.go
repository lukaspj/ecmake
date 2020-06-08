package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/lukaspj/ecmake/pkg/buildfile"
	"os"
)

type ModuleTest struct {
	logger hclog.Logger
}

func (m *ModuleTest) GetMethods() []string {
	m.logger.Info("GetMethods called")
	return []string{"test1", "test2"}
}

func (m *ModuleTest) Invoke(cmd string, args []interface{}) interface{} {
	m.logger.Info("Invoke called", "cmd", cmd, "args", args)
	switch cmd {
	case "test1":
		m.logger.Info("Test1")
	case "test2":
		m.logger.Info("Test2")
	default:
		m.logger.Error("Unknown")
	}
	return cmd
}

var _ buildfile.Module = &ModuleTest{}


func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:            "TestModule",
		Level:           hclog.Trace,
		Output:          os.Stderr,
		JSONFormat:      true,
	})

	module := &ModuleTest{
		logger: logger,
	}

	var pluginMap = map[string]plugin.Plugin{
		"module": &buildfile.ModulePlugin{Impl: module},
	}

	logger.Debug("Test Module initialized")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: buildfile.HandshakeConfig,
		Plugins: pluginMap,
	})
}

