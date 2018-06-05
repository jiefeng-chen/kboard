package resource

import "gopkg.in/yaml.v2"

type resPersistentVolumeClaim interface {
	SetMetaDataName(string) bool
	SetNamespace(string) bool
	SetAccessMode(string) bool
	SetStorage(string) bool
	SetVolumeName(string) bool
	SetStorageClassName(string) bool
}

type ResPersistentVolumeClaim struct {
	ApiVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec struct {
		AccessModes string `yaml:"accessModes"`
		Resources   struct {
			Request struct {
				Storage string
			}
		}
		StorageClassName string `yaml:"storageClassName"`
		VolumeName       string `yaml:"volumeName"`
	}
	Kind string
}

func NewPersistentVolumeClaim() *ResPersistentVolumeClaim {
	return &ResPersistentVolumeClaim{
		ApiVersion: "v1",
		Kind:       "PersistentVolumeClaim",
	}
}

func (r *ResPersistentVolumeClaim) GetAccessModes() string {
	return r.Spec.AccessModes
}

func (r *ResPersistentVolumeClaim) GetStorage() string {
	return r.Spec.Resources.Request.Storage
}

func (r *ResPersistentVolumeClaim) GetStorageClassName() string {
	return r.Spec.StorageClassName
}

func (r *ResPersistentVolumeClaim) SetMetaDataName(name string) bool {
	r.Metadata.Name = name
	return true
}

func (r *ResPersistentVolumeClaim) SetNamespace(ns string) bool {
	r.Metadata.Namespace = ns
	return true
}

func (r *ResPersistentVolumeClaim) SetAccessMode(am string) bool {
	r.Spec.AccessModes = am
	return true
}

func (r *ResPersistentVolumeClaim) SetStorage(storage string) bool {
	r.Spec.Resources.Request.Storage = storage
	return true
}

func (r *ResPersistentVolumeClaim) SetVolumeName(vName string) bool {
	r.Spec.VolumeName = vName
	return true
}

func (r *ResPersistentVolumeClaim) SetStorageClassName(scName string) bool {
	r.Spec.StorageClassName = scName
	return true
}

func (r *ResPersistentVolumeClaim) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
