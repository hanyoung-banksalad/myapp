package config

import (
	"github.com/smira/go-statsd"

	"github.com/banksalad/idl/gen/go/apis/v1/auth"
)

type Config interface {
	Setting() Setting

	StatsdClient() *statsd.Client

	AuthCli() auth.AuthClient
}

// verify DefaultConfig implements all Config interface methods
var _ Config = (*DefaultConfig)(nil)

type DefaultConfig struct {
	setting Setting

	statsdClient *statsd.Client

	authCli       auth.AuthClient
}

func (c *DefaultConfig) Setting() Setting {
	return c.setting
}

func (c *DefaultConfig) StatsdClient() *statsd.Client {
	return c.statsdClient
}

func (c *DefaultConfig) AuthCli() auth.AuthClient {
	return c.authCli
}

func NewConfig(
	setting Setting,
	statsdClient *statsd.Client,
	authCli auth.AuthClient,
) Config {
	return &DefaultConfig{
		setting:      setting,
		statsdClient: statsdClient,
		authCli:      authCli,
	}
}
