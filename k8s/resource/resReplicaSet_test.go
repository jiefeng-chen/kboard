package resource

import "testing"

func TestNewReplicaSet(t *testing.T) {
	var replicaSet *ResReplicaSet

	replicaSet = NewReplicaSet()

	replicaSet.SetMetadataName("hello")
	replicaSet.SetNamespace("world")
	labels := map[string]string{
		"app": "app",
		"value": "value",
	}
	replicaSet.SetLabels(labels)

	t.Fatalf("%v", replicaSet)
}

func TestNewContainer(t *testing.T) {
	var container *Container

	container = NewContainer("nginx", "nginx:latest")

	container.Args = "/bin/sh -c"
	env := NewEnv()
	env.Name = "SYS_ENV"
	env.ValueFrom.ResourceFieldRef.Resource = "1"
	env.ValueFrom.ResourceFieldRef.ContainerName = "container"
	container.AppendEnv(env)
	container.Command = "hello"
	container.Resources = NewResource()
	container.VolumeMounts = []map[string]interface{}{}
	container.AppendPort(NewPort("port"))
	container.AppendVolumeMount(map[string]interface{}{
		"volume1": "1313",
	})

	t.Fatalf("%v", container)
}



