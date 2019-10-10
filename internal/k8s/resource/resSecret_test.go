package resource

import (
	"kboard/internal"
	"testing"
)

func TestResSecret_SetData(t *testing.T) {
	secretStr := internal.NewSecret()
	data := []map[string]string{
		{"key": "h", "val": "hello"},
		{"key": "w", "val": "world"},
	}
	err := secretStr.SetData(data)
	if err != nil {
		t.Error(err)
	}
	t.Fatalf("%+v", secretStr)
}

func TestResSecret_GetData(t *testing.T) {
	secretStr := internal.NewSecret()
	data := []map[string]string{
		{"key": "h", "val": "hello"},
		{"key": "w", "val": "world"},
	}
	secretStr.SetData(data)
	d, err := secretStr.GetData()
	if err != nil {
		t.Error(err.Error())
	}
	for k, v := range d {
		t.Logf("%+v => %+v", k, v)
	}
}
