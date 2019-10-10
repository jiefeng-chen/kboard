package resource

import (
	"kboard/internal"
	"testing"
)

func TestResStorageClass_SetParameters(t *testing.T) {
	sc := internal.NewStorageClass()
	cephrbd := internal.NewCephRbd()
	data := []map[string]string{
		{
			"key": "adminId",
			"val": "1341234",
		},
		{
			"key": "monitors",
			"val": "192.168.312.31,192.143.131.13",
		},
		{
			"key": "adminSecretName",
			"val": "srtwert",
		},
		{
			"key": "adminSecretNamespace",
			"val": "srtwer123t",
		},
		{
			"key": "pool",
			"val": "srtt",
		},
		{
			"key": "userId",
			"val": "srtwert",
		},
		{
			"key": "userSecretName",
			"val": "srtwert",
		},
		{
			"key": "fsType",
			"val": "xfs",
		},
		{
			"key": "imageFormat",
			"val": "layring",
		},
		{
			"key": "imageFeatures",
			"val": "2",
		},
	}
	err := cephrbd.SetData(data)
	if err != nil {
		t.Error(err)
	}
	sc.SetParameters(*cephrbd)
	t.Fatalf("%+v", sc)
}
