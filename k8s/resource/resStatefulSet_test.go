package resource

import "testing"

func TestNewResStatefulSet(t *testing.T) {
	var resStatefulSet IResStatefulSet

	resStatefulSet = NewResStatefulSet()
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
	var container IContainer
	container = NewContainer("nginx", "nginx:latest")

	resStatefulSet.AddContainer(container)

	resStatefulSet.ToYamlFile()

	t.Fatalf("%v", resStatefulSet)
}
