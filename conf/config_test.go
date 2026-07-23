package conf_test

import (
	"blog/conf"
	"testing"
)

func TestC(t *testing.T) {
	t.Log(conf.C().Yaml())
}

func TestYaml(t *testing.T) {
	err := conf.FromYaml("./app.yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().Yaml())
}
