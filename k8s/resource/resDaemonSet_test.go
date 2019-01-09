package resource

import "testing"

func TestNewResDaemonSet(t *testing.T) {
	var daemonSet IResDaemonSet

	daemonSet = NewResDaemonSet()
	daemonSet.SetNamespace("my")
	daemonSet.SetMetaDataName("")
}
