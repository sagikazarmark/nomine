package app

import "time"

// Configuration holds any kind of config that is necessary for running
type Configuration struct {
	// Recommended values are: production, development, staging, release/123, etc
	Environment string `default:"production"`
	Debug       bool   `split_words:"true"`

	GrpcServiceAddr string        `ignored:"true"`
	RestServiceAddr string        `ignored:"true"`
	HealthAddr      string        `ignored:"true"`
	DebugAddr       string        `ignored:"true"`
	ShutdownTimeout time.Duration `ignored:"true"`

	FluentdEnabled bool   `split_words:"true"`
	FluentdHost    string `split_words:"true"`
	FluentdPort    int    `split_words:"true" default:"24224"`
}
