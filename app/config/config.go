package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	PortStart int `json:"port_start" yaml:"port_start"`

	DatabaseHost string `json:"database_host" yaml:"database_host"`
	DatabasePort int    `json:"database_port" yaml:"database_port"`
	DatabaseUser string `json:"database_user" yaml:"database_user"`
	DatabasePass string `json:"database_pass" yaml:"database_pass"`
	DatabaseName string `json:"database_name" yaml:"database_name"`
}

func GetConfig(configDir string) (*Configuration, error) {
	var err error
	var filename string

	if configDir != "" {
		filename, _ = filepath.Abs(configDir)
	} else {
		filename, _ = filepath.Abs("config.yaml")
	}

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	cfg := new(Configuration)
	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
