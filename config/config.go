package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/yaoapp/yao/lib/exception"
	"github.com/yaoapp/yao/lib/json"
)

// Setting the config setting
var Setting *Config

// Load load the config
func Load(name string, path ...string) *Config {

	var config *Config = nil
	for _, p := range path {
		file := filepath.Join(p, name)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			continue
		}
		config = &Config{}
		json.DecodeFile(file, config)
		break
	}

	if config == nil {
		exception.New("Can't find configuration file", 404).
			Ctx(json.M{"path": path}).
			Throw()
	}
	return config
}

// PWD Get curernt path
func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		exception.Err(err, 500).
			Throw()
	}

	return pwd
}

// HomeDir Get user home path
func HomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}
