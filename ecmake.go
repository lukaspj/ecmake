package main

import (
	"github.com/lukaspj/ecmake/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.Execute(cmd.Config{
		Version: version,
		Commit:  commit,
		Date:    date,
	})
}
