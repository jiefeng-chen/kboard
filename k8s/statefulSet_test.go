package k8s

import (
	config2 "kboard/config"
	"kboard/k8s/resource"
	"log"
	"testing"
)

func TestNewStatefulSet(t *testing.T) {
	config := config2.NewConfig()
	lib := NewStatefulSet(config)
	statefulSet := resource.NewResStatefulSet()
	statefulSet.SetMetaDataName("mystateful")
	statefulSet.SetNamespace("namespace")
	statefulSet.SetReplicas(3)
	container := resource.NewContainer("mycontainer", "image")
	statefulSet.AddContainer(container)

	yamlData, err := statefulSet.ToYamlFile()
	if err != nil {
		log.Printf("%v", err)
	}
	res := lib.WriteToEtcd("myapp", "mystateful", yamlData)
	t.Errorf("%v", res)
}
