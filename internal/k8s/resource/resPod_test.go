package resource

import (
	"kboard/internal"
	"testing"
)

func TestNewPod(t *testing.T) {
	var pod internal.IResPod

	pod = internal.NewResPod("pod1")
	pod.SetNamespace("namespace")
	labels := map[string]string{
		"app": "app",
		"val": "val",
	}
	pod.SetLabels(labels)
	container := internal.NewContainer("container1", "image1")
	pod.AddContainer(container)
	volume := internal.NewVolume()
	volume.Name = "vol1"
	volume.Secret = &internal.Secret{
		SecretName: "secret1",
		Items:      []map[string]string{},
	}
	pod.AddVolume(volume)
	pod.SetRestartPolicy("policy1")

	t.Fatalf("%v", pod)
}
