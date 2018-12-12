package resource

import "testing"

func TestNewResDeployment(t *testing.T) {
	var deploy *ResDeployment
	deploy = NewResDeployment()

	deploy.SetMetadataName("name")
	deploy.SetNamespace("namespace")
	labels := map[string]string{
		"app": "nginx",
	}
	deploy.SetMatchLabels(labels)
	deploy.SetTemplateLabels(labels)
	container := NewContainer("container", "image")
	deploy.AddContainer(container)

	deploy.ToYamlFile()

	t.Fatalf("%v", deploy)
}
