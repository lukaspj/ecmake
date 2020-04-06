package gojafile

import (
	"github.com/dop251/goja"
	"github.com/lukaspj/ecmake/pkg/buildfile"
	"github.com/pkg/errors"
)

type GojaFile struct {
	runtime *goja.Runtime
	script  goja.Value
	targets map[string]goja.Value
}

type GojaTarget struct {
	name string
}

func (g GojaTarget) GetName() string {
	return g.name
}

func (g GojaFile) GetTargets() []buildfile.Target {
	var targets []buildfile.Target
	for n, _ := range g.targets {
		targets = append(targets, GojaTarget{name: n})
	}

	return targets
}

func (g GojaFile) GetTarget(target string) buildfile.Target {
	targets := g.GetTargets()
	for _, t := range targets {
		if t.GetName() == target {
			return t
		}
	}
	return nil
}

func (g GojaFile) RunTarget(target buildfile.Target, args []string) (int, error) {
	var fn goja.Callable
	var ok bool
	for n, t := range g.targets {
		if n != target.GetName() {
			continue
		}

		fn, ok = goja.AssertFunction(t)
		if !ok {
			return -1, errors.WithStack(TargetInvalidError{Target: n})
		}

		break
	}
	if fn == nil {
		return -1, errors.WithStack(TargetNotFoundError{Target: target.GetName()})
	}

	returnRaw, err := fn(nil, g.runtime.ToValue(args))
	if err != nil {
		return -1, errors.WithStack(TargetExecutionError{Target: target.GetName(), Args: args, Cause: err})
	}

	return int(returnRaw.ToInteger()), nil
}

func NewGojaFile(vm *goja.Runtime, path, content string) (*GojaFile, error) {
	var targets map[string]goja.Value

	setTargets := func(ts map[string]goja.Value) {
		targets = ts
	}
	vm.Set("SetTargets", setTargets)

	script, err := vm.RunScript(path, content)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &GojaFile{
		runtime: vm,
		script:  script,
		targets: targets,
	}, nil
}
