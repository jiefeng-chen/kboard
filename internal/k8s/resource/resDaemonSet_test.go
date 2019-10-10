package resource

import (
	"kboard/internal"
	"testing"
)

func TestNewResDaemonSet(t *testing.T) {
	var daemonSet internal.IResDaemonSet

	daemonSet = internal.NewResDaemonSet()
	daemonSet.SetNamespace("my")
	daemonSet.SetMetaDataName("")
}
