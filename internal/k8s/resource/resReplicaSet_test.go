package resource

import (
	"kboard/internal"
	"testing"
)

func TestNewReplicaSet(t *testing.T) {
	var replicaSet internal.IResReplicaSet

	replicaSet = internal.NewResReplicaSet()

	replicaSet.SetMetadataName("hello")
	replicaSet.SetNamespace("world")
	labels := map[string]string{
		"app":   "app",
		"value": "value",
	}
	replicaSet.SetLabels(labels)

	t.Fatalf("%v", replicaSet)
}
