package resource

import (
	"kboard/internal"
	"testing"
)

func TestResConfigMap_SetData(t *testing.T) {
	data := []map[string]string{
		{"key": "1", "val": "1"},
		{"key": "2", "val": "2"},
	}
	confMap := internal.NewResConfigMap()
	confMap.SetData(data)
	confMap.SetNamespace("helle")
	confMap.SetMetadataName("132413")
	t.Fatalf("%+v", confMap)
}
