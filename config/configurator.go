package config

import "archer/structs"

//Config ...
type Config struct {
	APPName string

	DB struct {
		Name  string `default:"archerdb"`
		Local string `default:"mongodb://localhost"`
	}
}

// New initialize a Configor
func New(config *structs.Config) *structs.Configurator {
	if config == nil {
		config = &structs.Config{}
	}
	return &structs.Configurator{Config: config}
}

// ENV return environment
func ENV() string {
	return New(nil).GetEnvironment()
}

// Load will unmarshal configurations to struct from files that you provide
func Load(config interface{}, files string) error {
	return New(nil).Load(config, files)
}

//Get ...
func Get() *Config {
	var conf = &Config{}
	Load(conf, "config.yml")

	return conf
}
