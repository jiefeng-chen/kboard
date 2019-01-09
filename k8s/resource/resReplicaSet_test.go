package resource

import "testing"

func TestNewReplicaSet(t *testing.T) {
	var replicaSet IResReplicaSet

	replicaSet = NewResReplicaSet()

	replicaSet.SetMetadataName("hello")
	replicaSet.SetNamespace("world")
	labels := map[string]string{
		"app":   "app",
		"value": "value",
	}
	replicaSet.SetLabels(labels)

	t.Fatalf("%v", replicaSet)
}
