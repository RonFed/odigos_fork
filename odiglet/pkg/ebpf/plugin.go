package ebpf

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/odigos-io/odigos/instrumentation"
	odigosPlugin "github.com/odigos-io/odigos/instrumentation/plugin/shared"
	"github.com/odigos-io/odigos/k8sutils/pkg/service"
	"github.com/odigos-io/odigos/odiglet/pkg/env"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type instrumentedProcess struct {
	pid    int
	plugin odigosPlugin.Instrumentation
}

type pluginFactory struct {
	plugin odigosPlugin.Instrumentation
}

func CleanupPlugins() {
	plugin.CleanupClients()
}

func NewInstrumentationPluginFactory(path string) (instrumentation.Factory, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: odigosPlugin.Handshake,
		Plugins:         odigosPlugin.PluginMap,

		// point the plugin stdout and stderr to ours
		SyncStdout: os.Stdout,
		SyncStderr: os.Stderr,

		// we are going to call CleanupClients once the Odiglet goes down
		// we might want to change this in the future if we want to support dynamically adding and removing plugins.
		Managed: true,

		Logger: hclog.New(&hclog.LoggerOptions{
			Name: "plugin",
			Level: hclog.DefaultLevel,
			JSONFormat: true,
		}),

		// passing the endpoint to export to as an argument to the plugin
		// currently assuming a gRPC exporter will be used by the plugin
		Cmd:              exec.Command(path, service.LocalTrafficOTLPGrpcDataCollectionEndpoint(env.Current.NodeIP)),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})

	// use the gRPC client
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(odigosPlugin.InstrumentationPluginName)
	if err != nil {
		return nil, err
	}

	inst, ok := raw.(odigosPlugin.Instrumentation)
	if !ok {
		return nil, fmt.Errorf("unexpected type from plugin: %T", raw)
	}

	return &pluginFactory{plugin: inst}, nil
}

var _ instrumentation.Factory = (*pluginFactory)(nil)

func (f *pluginFactory) CreateInstrumentation(ctx context.Context, pid int, settings instrumentation.Settings) (instrumentation.Instrumentation, error) {
	if f.plugin == nil {
		return nil, errors.New("instrumentation plugin not initialized")
	}

	err := f.plugin.Start(ctx, pid, settings)
	if err != nil {
		return nil, fmt.Errorf("failed to start instrumentation: %w", err)
	}

	return &instrumentedProcess{pid: pid, plugin: f.plugin}, nil
}

func (g *instrumentedProcess) Load(ctx context.Context) error {
	return nil
}
func (g *instrumentedProcess) Run(ctx context.Context) error {
	return nil
}
func (g *instrumentedProcess) ApplyConfig(ctx context.Context, config instrumentation.Config) error {
	return nil
}

func (g *instrumentedProcess) Close(ctx context.Context) error {
	if g.plugin == nil {
		return errors.New("instrumentation plugin not initialized")
	}

	return g.plugin.Close(ctx, g.pid)
}
