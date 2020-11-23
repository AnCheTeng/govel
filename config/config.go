package config

// Reference: https://github.com/prometheus/prometheus/blob/76cd5f4c7f123041525f101611a2cc04fd3d5382/config/config.go#L138

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	// DefaultConfig is the config with parsing config.yaml
	DefaultConfig = config{}
)

func init() {
	loadFile("config/config.yaml")
}

type config struct {
	Vendors []*vendorConfig `yaml:"vendors,omitempty"`
}

type vendorConfig struct {
	Name     string   `yaml:"name"`
	Hostname string   `yaml:"hostname"`
	Novels   []*Novel `yaml:"novels"`
}

// Novel is novel
type Novel struct {
	Name     string `yaml:"name"`
	IndexURL string `yaml:"indexURL"`
}

func loadFile(filename string) (*config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg := &DefaultConfig
	err = yaml.UnmarshalStrict([]byte(content), cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing YAML file %s", filename)
	}
	return cfg, nil
}
