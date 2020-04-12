package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type NetworkCreateConfig struct {

}

func (d *Module) NetworkCreate(name string, config NetworkCreateConfig) error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	_, err = cli.NetworkCreate(ctx, name, types.NetworkCreate{
	})

	return err
}