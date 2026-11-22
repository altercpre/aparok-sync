package synctest

import (
	"github.com/altercpre/aparok-sync/app"
	"github.com/altercpre/aparok-sync/net/streampool"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init(a *app.App) (err error) {
	return
}

func (c *Config) Name() (name string) {
	return "config"
}

func (c *Config) GetStreamConfig() streampool.StreamConfig {
	return streampool.StreamConfig{
		SendQueueSize:    100,
		DialQueueWorkers: 100,
		DialQueueSize:    100,
	}
}
