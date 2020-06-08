package cmd

import (
	"encoding/gob"
	"fmt"
	"github.com/dop251/goja"
	"github.com/hashicorp/go-hclog"
	"github.com/lukaspj/ecmake/pkg/buildfile"
	"github.com/lukaspj/ecmake/pkg/console"
	"github.com/lukaspj/ecmake/pkg/docker"
	"github.com/lukaspj/ecmake/pkg/gojafile"
	"github.com/lukaspj/ecmake/pkg/io"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func exitWithError(err error) {
	fmt.Println(fmt.Sprintf("%+v", err))
	os.Exit(-1)
}

func getBuildFile() buildfile.BuildFile {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:            "ecmake",
		Level:           hclog.Trace,
		Output:          os.Stderr,
		JSONFormat:      false,
	})

	vm := goja.New()
	console.NewConsole().Inject(vm)
	docker.New().Inject(vm)
	io.New().Inject(vm)
	gob.Register(map[string]interface{}{})

	wd, err := os.Getwd()
	if err != nil {
		exitWithError(errors.Wrap(err, "failed to get current directory"))
	}

	file, err := gojafile.GetGojaFile(vm, wd)
	if err != nil {
		exitWithError(err)
	}

	gojafile.NewStdModule(logger).Inject(file)

	err = file.Initialize()
	if err != nil {
		exitWithError(err)
	}

	return file
}

func GetRootCmd(config Config) *cobra.Command {
	return &cobra.Command{
		Use:   "ecmake",
		Short: "ECMAke is a build-tool leveraging JavaScript for build logic",
		Args: func(cmd *cobra.Command, args []string) error {
			file := getBuildFile()
			targets := file.GetTargets()

			if len(args) == 0 {
				fmt.Println("Goja Targets:")
				for _, t := range file.GetTargets() {
					fmt.Println(fmt.Sprintf(" * %s", t.GetName()))
				}
				os.Exit(0)
			}

			found := false
			for _, t := range targets {
				found = found || t.GetName() == args[0]
			}
			if !found {
				return errors.Errorf("target %s does not exist in list %v", args[0], targets)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			file := getBuildFile()

			target := file.GetTarget(args[0])
			errorCode, err := file.RunTarget(target, args[1:])
			if err != nil {
				exitWithError(err)
			}

			err = file.Close()
			if err != nil {
				exitWithError(err)
			}

			os.Exit(errorCode)
		},
	}
}

type Config struct {
	Version string
	Commit  string
	Date    string
}

func Execute(config Config) {
	rootCmd := GetRootCmd(config)
	rootCmd.AddCommand(GetVersionCommand(config))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
