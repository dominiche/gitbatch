package config

import (
	"fmt"
	"github.com/kylelemons/go-gypsy/yaml"
	"os"
	"path/filepath"
)

var GitType string
var GitPath string
var PrivateToken string

func init() {
	//dir := filepath.Dir(os.Args[0]) //can't work if we set env path and invoke in other folder
	ex, err := os.Executable() //this works
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(ex)
	absPath, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	configPath := absPath + string(filepath.Separator) + "config.yaml"
	config, err := yaml.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
	}

	GitType = getConfig(config, "git.type")
	PrivateToken = getConfig(config, "git.token")
	GitPath = getConfig(config, "git.path")
}

func getConfig(config *yaml.File, prop string) string {
	value, e := config.Get(prop)
	if e != nil {
		fmt.Println(prop+" can't be empty! ", e)
		os.Exit(1)
	}
	return value
}
