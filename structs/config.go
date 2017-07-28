package structs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v1"
)

//Configurator ...
type Configurator struct {
	*Config
}

//Config ...
type Config struct {
	Environment string
	ENVPrefix   string
}

// GetEnvironment get environment
func (configor *Configurator) GetEnvironment() string {
	if configor.Environment == "" {
		if env := os.Getenv("CONFIG_ENV"); env != "" {
			return env
		}

		return "development"
	}
	return configor.Environment
}

// Load will unmarshal configurations to struct from files that you provide
func (configor *Configurator) Load(config interface{}, files string) error {
	file, _ := configor.getConfigurationFiles(files)

	return processFile(config, file)

}

func (configor *Configurator) getConfigurationFiles(file string) (string, error) {
	return getConfigurationFileWithENVPrefix(file, configor.GetEnvironment())
}

func (configor *Configurator) getENVPrefix(config interface{}) string {
	if configor.Config.ENVPrefix == "" {
		if prefix := os.Getenv("CONFIG_ENV_PREFIX"); prefix != "" {
			return prefix
		}
		return "config"
	}
	return configor.Config.ENVPrefix
}

func getConfigurationFileWithENVPrefix(file, env string) (string, error) {
	var (
		envFile string
		extname = path.Ext(file)
	)

	if extname == "" {
		envFile = fmt.Sprintf("%v.%v", file, env)
	} else {
		envFile = fmt.Sprintf("%v.%v%v", strings.TrimSuffix(file, extname), env, extname)
	}

	absPath, err := filepath.Abs((os.Getenv("GOPATH") + "/archer/config/" + envFile))

	return absPath, err
}

func processFile(config interface{}, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, config)
}

func getPrefixForStruct(prefixes []string, fieldStruct *reflect.StructField) []string {
	if fieldStruct.Anonymous && fieldStruct.Tag.Get("anonymous") == "true" {
		return prefixes
	}
	return append(prefixes, fieldStruct.Name)
}
