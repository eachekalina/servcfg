package servcfg

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ServerConfig struct {
	Host    string
	Port    int
	Timeout int `yaml:"timeout,omitempty"`
}

func ParseServerConfig(path string) (ServerConfig, error) {
	var cfg ServerConfig
	bytes, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
