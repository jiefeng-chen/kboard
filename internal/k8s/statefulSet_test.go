package k8s

import (
	config2 "kboard/config"
	"kboard/internal"
	"log"
	"testing"
)

func TestNewStatefulSet(t *testing.T) {
	config := config2.NewConfig()
	lib := internal.NewStatefulSet(config)
	statefulSet := internal.NewResStatefulSet()
	statefulSet.SetMetaDataName("mystateful")
	statefulSet.SetNamespace("namespace")
	statefulSet.SetReplicas(3)
	container := internal.NewContainer("mycontainer", "image")
	statefulSet.AddContainer(container)

	yamlData, err := statefulSet.ToYamlFile()
	if err != nil {
		log.Printf("%v", err)
	}
	res := lib.WriteToEtcd("myapp", "mystateful", yamlData)
	t.Errorf("%v", res)
}
