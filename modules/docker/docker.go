package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"os"

	// Docker
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

type Module struct {
}

func New() *Module {
	return &Module{}
}

type DockerRegistryConfig struct {
	Username string
	Password string
	Host     string
}

type DockerBuildConfig struct {
	ContextPath string `json:"context_path"`
	Tag         string `json:"tag"`
	DockerFile  string `json:"docker_file,omitempty"`
}

func (d *Module) Build(config DockerBuildConfig) error {
	if config.ContextPath == "" {
		config.ContextPath = "."
	}
	if config.DockerFile == "" {
		config.DockerFile = "Dockerfile"
	}

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	dockerBuildContext, err := archive.TarWithOptions(config.ContextPath, &archive.TarOptions{})
	defer dockerBuildContext.Close()

	buildOptions := types.ImageBuildOptions{
		Dockerfile: config.DockerFile,
		Tags:       []string{config.Tag},
	}

	buildResponse, err := cli.ImageBuild(ctx, dockerBuildContext, buildOptions)
	if err != nil {
		return err
	}
	defer buildResponse.Body.Close()

	termFd, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
	return nil
}

func (d *Module) Push(config DockerRegistryConfig, tag string) error {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	auth := struct {
		Username string
		Password string
	}{
		Username: config.Username,
		Password: config.Password,
	}

	authBytes, _ := json.Marshal(auth)

	registryAuth := base64.StdEncoding.EncodeToString(authBytes)

	bodyReadCloser, err := cli.ImagePush(ctx, tag, types.ImagePushOptions{
		RegistryAuth: registryAuth,
	})
	if err != nil {
		return errors.Wrapf(err, "failed to push image to %s", config.Host)
	}
	defer bodyReadCloser.Close()

	termFd, isTerm := term.GetFdInfo(os.Stderr)
	jsonmessage.DisplayJSONMessagesStream(bodyReadCloser, os.Stderr, termFd, isTerm, nil)

	return nil
}
