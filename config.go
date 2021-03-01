package main

import (
	"github.com/juju/juju/api"
	"github.com/juju/names/v4"
)

// config is configuration for the Juju exporter.
type config struct {
	Default string                 `yaml:"default"`
	Models  map[string]*jujuConfig `yaml:"models"`
}

// jujuConfig is configuration for connecting to Juju.
type jujuConfig struct {
	APIEndpoints []string `yaml:"api-endpoints"`
	CACert       string   `yaml:"ca-cert"`
	Username     string   `yaml:"username"`
	Password     string   `yaml:"password"`
	ModelUUID    string   `yaml:"model-uuid"`
	Patterns     []string `yaml:"patterns"`

	registry *registry
}

// newClient initiates a new Juju client from configuration.
func (c jujuConfig) newClient() (api.Connection, error) {
	return api.Open(&api.Info{
		Addrs:    c.APIEndpoints,
		CACert:   c.CACert,
		ModelTag: names.NewModelTag(c.ModelUUID),
		Tag:      names.NewUserTag(c.Username),
		Password: c.Password,
	}, api.DefaultDialOpts())
}
