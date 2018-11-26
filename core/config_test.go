package core

import (
	"testing"
)

func TestConfig_LoadConfigFile(t *testing.T) {
	conf := NewConfig().LoadConfigFile("../config/conf.toml")
	t.Errorf("%+v", conf.Data)
}
