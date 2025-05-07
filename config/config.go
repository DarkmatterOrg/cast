package config

import (
	"cast/lib"

	"github.com/BurntSushi/toml"
	"github.com/darkmatterorg/orbit/utils"
)

type Modules struct {
	Horizon bool `toml:"horizon"`
	Experimental bool `toml:"experimental"`
}

type Settings struct {
	Insult bool `toml:"insult"`
	Modules Modules `toml:"modules"`
}

var Config Settings

const fullPath = "/etc/cast/config.toml"

func LoadConfig() {
	_, err := toml.DecodeFile(fullPath, &Config)
	if err != nil {
	
	}
}

func FindConfig() {
	if utils.PathExists(fullPath) {
		lib.Logger.Info("Config file found at " + fullPath)
	} else {
		lib.Logger.Info("Config should be created at " + fullPath)
	}
}