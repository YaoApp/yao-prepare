package sys

import (
	"os"
	"runtime"

	"github.com/yaoapp/yao/lib/exception"
)

// PWD Get curernt path
func PWD() string {
	pwd, err := os.Getwd()
	if err != nil {
		exception.Err(err, 500).
			Throw()
	}

	return pwd
}

// Home Get user home path
func Home() string {
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
