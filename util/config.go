package util

import (
	"github.com/BurntSushi/toml"
	"github.com/darkmatterorg/orbit/utils"
)

type Modules struct {
	Experimental bool `toml:"experimental"`
	Fixes bool `toml:"fixes"`
	Umbra bool `toml:"umbra"`
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
		Logger.Info("Config file found at " + fullPath)
	} else {
		Logger.Info("Config should be created at " + fullPath)
	}
}