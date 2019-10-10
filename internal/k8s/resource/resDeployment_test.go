package resource

import (
	"kboard/internal"
	"testing"
)

func TestNewResDeployment(t *testing.T) {
	var deploy internal.IResDeployment
	deploy = internal.NewResDeployment()

	deploy.SetMetadataName("name")
	deploy.SetNamespace("namespace")
	labels := map[string]string{
		"app": "nginx",
	}
	deploy.SetMatchLabels(labels)
	deploy.SetTemplateLabels(labels)
	container := internal.NewContainer("container", "image")
	deploy.AddContainer(container)

	deploy.ToYamlFile()

	t.Fatalf("%v", deploy)
}
