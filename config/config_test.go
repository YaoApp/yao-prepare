package config

import (
	"fmt"
	"testing"
)

func TestSetting(t *testing.T) {
	Setting = Load("config.yset", PWD()+"/assets/")
	fmt.Printf("%#v\n", Setting.Database["main"][0])
}
