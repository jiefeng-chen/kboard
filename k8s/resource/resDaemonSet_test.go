package resource

import "testing"

func TestNewResDaemonSet(t *testing.T) {
	var daemonSet *ResDaemonSet

	daemonSet = NewResDaemonSet()
	daemonSet.SetNamespace("my")
	daemonSet.SetMetaDataName("")
}
