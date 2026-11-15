package config

import "sync"

var (
	cfg     *config
	cfgOnce sync.Once
)

func init() {
	cfgOnce.Do(
		func() {
			// todo :- add defaults here
			cfgObj := config{}
			cfg = &cfgObj
		},
	)
}

type config struct {
	httpConfig *ConnCfg
	rcConfig   *ConnCfg
}

type ConnCfg struct {
	Port    *int
	Timeout *int
	Host    *string
}

func GetHTTPPort() *int {
	return cfg.httpConfig.Port
}
