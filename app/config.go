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

	GithubToken           string `split_words:"true"`
	TwitterConsumerKey    string `split_words:"true" required:"true"`
	TwitterConsumerSecret string `split_words:"true" required:"true"`
	TwitterAccessKey      string `split_words:"true" required:"true"`
	TwitterAccessSecret   string `split_words:"true" required:"true"`
	DNSimpleAccountID     string `envconfig:"dnsimple_account_id" required:"true"`
	DNSimpleToken         string `envconfig:"dnsimple_token" required:"true"`
}
