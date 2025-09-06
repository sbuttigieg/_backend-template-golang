package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/conf/v3"
)

type Config struct {
	namespace   string
	App         App
	Args        conf.Args
	DebugServer debugServer
}

type App struct {
	DestroyTimeout time.Duration `conf:"default:40s"`
}

type debugServer struct {
	WebPortBackendTemplate int `conf:"default:9000,env:DEBUG_SERVER_WEB_PORT"`
}

func New(ns string) *Config {
	return &Config{namespace: ns}
}

func (c *Config) Parse() error {
	if help, pErr := conf.Parse(c.namespace, c); pErr != nil {
		if errors.Is(pErr, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return pErr
	}

	return nil
}
