package core

import "testing"

func TestConfig_LoadConfigFile(t *testing.T) {
	conf := NewConfig().LoadConfigFile("../config/conf.yaml")
	t.Errorf("%+v", conf.Data)
}
