module github.com/lukaspj/ecmake

go 1.14

replace (
	github.com/docker/docker => github.com/docker/engine v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190830141801-acfa387b8d69
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/hcsshim v0.8.7 // indirect
	github.com/containerd/containerd v1.3.3 // indirect
	github.com/containerd/continuity v0.0.0-20200228182428-0f16d7a0959c // indirect
	github.com/dlclark/regexp2 v1.2.0 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v17.12.0-ce-rc1.0.20200309214505-aa6a9891b09c+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/dop251/goja v0.0.0-20200326102500-6438c8ddc517
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/hashicorp/go-plugin v1.3.0
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v0.1.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.7
	golang.org/x/sync v0.0.0-20190423024810-112230192c58 // indirect
	golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae // indirect
)
