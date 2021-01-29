package config

import (
	"fmt"
	"testing"
)

func TestSetting(t *testing.T) {
	Setting = Load("config.json", PWD())
	fmt.Printf("%#v\n", Setting)
}
