package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"sync/atomic"

	"github.com/hashicorp/go-plugin"
	"github.com/odigos-io/odigos/instrumentation"
	odigosPlugin "github.com/odigos-io/odigos/instrumentation/plugin/shared"
	"github.com/odigos-io/odigos/odiglet/pkg/ebpf"
	"go.opentelemetry.io/auto"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
)

type instrumentedProcess struct {
	i      *auto.Instrumentation
	cp     *ebpf.ConfigProvider[auto.InstrumentationConfig]
	closed *atomic.Bool
}

var _ auto.ConfigProvider = (*ebpf.ConfigProvider[auto.InstrumentationConfig])(nil)

type goInstrumentationServer struct {
	instrumentationsByPid map[int]*instrumentedProcess
	logger                *slog.Logger
	endpoint              string
}

func (s *goInstrumentationServer) Start(ctx context.Context, pid int, settings instrumentation.Settings) error {
	defaultExporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(s.endpoint),
	)
	if err != nil {
		return fmt.Errorf("failed to create exporter: %w", err)
	}

	// initialConfig, err := convertToGoInstrumentationConfig(settings.InitialConfig)
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid initial config type, expected *odigosv1.SdkConfig, got %T", settings.InitialConfig)
	// }

	// cp := ebpf.NewConfigProvider(initi alConfig)

	inst, err := auto.NewInstrumentation(
		ctx,
		auto.WithEnv(), // for OTEL_LOG_LEVEL
		auto.WithPID(pid),
		auto.WithResourceAttributes(settings.ResourceAttributes...),
		auto.WithServiceName(settings.ServiceName),
		auto.WithTraceExporter(defaultExporter),
		auto.WithGlobal(),
		// auto.WithConfigProvider(cp),
	)
	if err != nil {
		s.logger.Error("instrumentation setup failed", "error", err)
		return err
	}

	err = inst.Load(ctx)
	if err != nil {
		s.logger.Error("instrumentation load failed", "error", err)
		return err
	}

	closed := &atomic.Bool{}

	s.instrumentationsByPid[pid] = &instrumentedProcess{
		i:      inst,
		closed: closed,
		// cp: nil,
	}

	runCtx := context.Background()
	go func() {
		err := inst.Run(runCtx)
		if err != nil && !errors.Is(err, context.Canceled) {
			s.logger.Error("instrumentation run failed", "pid", pid, "error", err)
		}
		closed.Store(true)
	}()

	fmt.Println("#############################")
	s.logger.Info("instrumentation started", "pid", pid, "serviceName", settings.ServiceName)
	fmt.Println("#############################")
	return nil
}

func (s *goInstrumentationServer) ApplyConfig(ctx context.Context, pid int, config instrumentation.Config) error {
	return nil
}

func (s *goInstrumentationServer) Close(ctx context.Context, pid int) error {
	if inst, ok := s.instrumentationsByPid[pid]; ok {
		if inst.closed.Load() {
			s.logger.Warn("instrumentation already closed", "pid", pid)
			return nil
		}
		if err := inst.i.Close(); err != nil {
			s.logger.Error("failed to shutdown instrumentation", "pid", pid, "error", err)
			return err
		}
		delete(s.instrumentationsByPid, pid)
		s.logger.Info("instrumentation closed", "pid", pid)
	} else {
		// TODO: should we return an error here?
		s.logger.Warn("no instrumentation found for pid", "pid", pid)
	}
	return nil
}

func main() {
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	logger := slog.New(handler)
	logger = logger.WithGroup("golang-instrumentation-plugin")

	args := os.Args
	if len(args) < 2 {
		logger.Error("missing endpoint argument")
		os.Exit(1)
	}
	endpoint := args[1]
	if endpoint == "" {
		logger.Error("empty endpoint argument")
		os.Exit(1)
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: odigosPlugin.Handshake,
		Plugins: map[string]plugin.Plugin{
			odigosPlugin.InstrumentationPluginName: &odigosPlugin.InstrumentationPlugin{
				Impl: &goInstrumentationServer{
					logger:                logger,
					endpoint:              endpoint,
					instrumentationsByPid: make(map[int]*instrumentedProcess),
				},
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
