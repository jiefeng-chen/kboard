package resource

import "testing"

func TestNewPod(t *testing.T) {
	var pod IResPod

	pod = NewResPod("pod1")
	pod.SetNamespace("namespace")
	labels := map[string]string{
		"app": "app",
		"val": "val",
	}
	pod.SetLabels(labels)
	container := NewContainer("container1", "image1")
	pod.AddContainer(container)
	volume := NewVolume()
	volume.Name = "vol1"
	volume.Secret = &Secret{
		SecretName: "secret1",
		Items:      []map[string]string{},
	}
	pod.AddVolume(volume)
	pod.SetRestartPolicy("policy1")

	t.Fatalf("%v", pod)
}
