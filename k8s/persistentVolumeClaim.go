package k8s

import (
	"github.com/revel/config"
	"kboard/resource"
)

type IPersistentVolumeClaim interface {
	IK8sCore
	ReadStatus(string, string) error
}

func NewPersistentVolumeClaim(Config *config.Context) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		K8sCore{
			Config: Config,
			Kind:   resource.RESOURCE_PERSISTENT_VOLUME_CLAIM,
			Urls: Urls{
				Read:   "/api/v1/namespaces/%s/persistentvolumeclaims/%s",
				Create: "/api/v1/namespaces/%s/persistentvolumeclaims",
			},
		},
	}
}

type PersistentVolumeClaim struct {
	K8sCore
}

func (l *PersistentVolumeClaim) ReadStatus(ns string, name string) error {
	return nil
}
