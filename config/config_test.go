package config

import (
	"testing"

	"github.com/yaoapp/yao/lib/json"
)

func TestSetting(t *testing.T) {
	Setting = Load("config.yset", PWD()+"/assets/")
	json.PrettyPrint(Setting)
}
