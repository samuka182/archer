package config

import (
	"archer/structs"
)

// New initialize a Configor
func New() *structs.Configurator {
	return &structs.Configurator{}
}

// ENV return environment
func ENV() string {
	return New().GetEnvironment()
}

// Load will unmarshal configurations to struct from files that you provide
func Load(config *structs.Config, files string) error {
	return New().Load(config, files)
}

//Get ...
func Get() *structs.Config {
	var conf = &structs.Config{}
	Load(conf, "config.yml")

	return conf
}
