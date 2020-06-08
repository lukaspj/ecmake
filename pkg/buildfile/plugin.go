package buildfile

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"log"
	"net/rpc"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

type Module interface {
	GetMethods() []string
	Invoke(cmd string, args []interface{}) interface{}
}

type InvokeArgs struct {
	Cmd string
	Args []interface{}
}

var _ Module = &ModuleRPC{}

type ModuleRPC struct{ client *rpc.Client }

func (g *ModuleRPC) Invoke(cmd string, args []interface{}) interface{} {
	invokeArgs := &InvokeArgs{
		Cmd: cmd,
		Args: args,
	}
	var resp interface{}
	err := g.client.Call("Plugin.Invoke", invokeArgs, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

func (g *ModuleRPC) GetMethods() []string {
	var resp []string
	err := g.client.Call("Plugin.GetMethods", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

type ModuleRPCServer struct {
	Impl Module
}

func (s *ModuleRPCServer) GetMethods(args interface{}, resp *[]string) error {
	*resp = s.Impl.GetMethods()
	return nil
}

func (s *ModuleRPCServer) Invoke(args *InvokeArgs, resp *interface{}) error {
	*resp = s.Impl.Invoke(args.Cmd, args.Args)
	return nil
}

var _ plugin.Plugin = &ModulePlugin{}

type ModulePlugin struct {
	// Impl Injection
	Impl Module
}

func (p *ModulePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ModuleRPCServer{Impl: p.Impl}, nil
}

func (ModulePlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ModuleRPC{client: c}, nil
}

type ModuleHost struct {
	logger    hclog.Logger
	client    *plugin.Client
	rpcClient plugin.ClientProtocol
}

func (m *ModuleHost) Dispense() Module {
	// Request the plugin
	raw, err := m.rpcClient.Dispense("module")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	module := raw.(Module)
	return module
}

func (m *ModuleHost) Close() {
	m.client.Kill()
}

func InitializeModule(logger hclog.Logger, path string) *ModuleHost {
	// We're a host! Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(fmt.Sprintf("%s", path)),
		Logger:          logger,
	})

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	return &ModuleHost{
		logger: logger.ResetNamed("plugin-host"),
		client: client,
		rpcClient: rpcClient,
	}
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "ECMAKE_JSONRPC2_PLUGIN",
	MagicCookieValue: "ECMAKE_JSONRPC2_PLUGIN_SYSTEM",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"module": &ModulePlugin{},
}
