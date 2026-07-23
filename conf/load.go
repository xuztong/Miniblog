package conf

import (
	"os"

	"github.com/goccy/go-yaml"
)

var config *Config

func C() *Config {
	if config == nil {
		config = Default()
	}
	return config
}

func FromYaml(path string) error {
	contont,err:=os.ReadFile(path)
	if err!=nil{
		return err
	}
	config=C()
	return yaml.Unmarshal(contont,config)
}
