package resource

import (
	"kboard/internal"
	"testing"
)

func TestNewResStatefulSet(t *testing.T) {
	var resStatefulSet internal.IResStatefulSet

	resStatefulSet = internal.NewResStatefulSet()
	// 1. name
	resStatefulSet.SetMetaDataName("name")
	resStatefulSet.SetNamespace("namespace")
	annos := map[string]string{
		"app": "nginx",
	}
	resStatefulSet.SetAnnotations(annos)
	labels := map[string]string{
		"app": "nginx",
	}
	resStatefulSet.SetLabels(labels)
	resStatefulSet.SetServiceName("service name")
	resStatefulSet.SetStorage("1Gi")
	resStatefulSet.SetReplicas(3)
	resStatefulSet.SetVolumeClaimName("volume claim name")
	resStatefulSet.SetAccessMode("ReadWriteOnce")
	var container internal.IContainer
	container = internal.NewContainer("nginx", "nginx:latest")

	resStatefulSet.AddContainer(container)

	resStatefulSet.ToYamlFile()

	t.Fatalf("%v", resStatefulSet)
}
