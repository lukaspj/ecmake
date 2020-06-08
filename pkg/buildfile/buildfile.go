package buildfile

import "io"

type Target interface {
	GetName() string
}

type BuildFile interface {
	io.Closer

	GetTargets() []Target
	GetTarget(target string) Target
	RunTarget(target Target, args []string) (int, error)
	AddModuleHost(host *ModuleHost)
}
