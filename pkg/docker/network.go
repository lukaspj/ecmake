package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func (d *Module) NetworkCreate(name string, config *types.NetworkCreate) error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	_, err = cli.NetworkCreate(ctx, name, *config)

	return err
}