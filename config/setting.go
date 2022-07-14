package config

import "github.com/banksalad/go-banksalad"

type Setting struct {
	ServiceName        string
	GRPCServerEndpoint string
	GRPCServerPort     string
	HTTPServerPort     string
	StatsdUDPPort      string

	Env       string
	SubEnvID  string
	Namespace string

	GracefulShutdownTimeoutMs int

	AuthGRPCEndpoint string
}

func NewSetting() Setting {
	return Setting{
		ServiceName:        "myapp",
		GRPCServerEndpoint: banksalad.MustGetEnv("GRPC_SERVER_ENDPOINT", "localhost:18081"),
		GRPCServerPort:     banksalad.MustGetEnv("GRPC_SERVER_PORT", "18081"),
		HTTPServerPort:     banksalad.MustGetEnv("HTTP_SERVER_PORT", "18082"),
		StatsdUDPPort:      banksalad.MustGetEnv("STATSD_UDP_PORT", "8125"),

		Env:       banksalad.MustGetEnv("ENV", "development"),
		SubEnvID:  banksalad.MustGetEnv("SUB_ENV_ID", "local"),
		Namespace: banksalad.MustGetEnv("NAMESPACE", "development-local"),

		GracefulShutdownTimeoutMs: banksalad.MustAtoi(banksalad.MustGetEnv("GRACEFUL_SHUTDOWN_TIMEOUT_MS", "10000")),

		AuthGRPCEndpoint: banksalad.MustGetEnv("AUTH_GRPC_ENDPOINT", "dns:///auth-headless:18081"),
	}
}
