package resource

import "gopkg.in/yaml.v2"

type IResPersistentVolumeClaim interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetAccessMode(string) error
	SetStorage(string) error
	SetVolumeName(string) error
	SetVolumeMode(string) error
	SetStorageClassName(string) error
}

type ResPersistentVolumeClaim struct {
	ApiVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name      string
		Namespace string
	}
	Spec struct {
		AccessModes []string `yaml:"accessModes" json:"accessModes"`
		Resources   struct {
			Requests struct {
				Storage string
			}
		}
		VolumeMode       string `yaml:"volumeMode" json:"volumeMode"`
		StorageClassName string `yaml:"storageClassName" json:"storageClassName"`
		VolumeName       string `yaml:"volumeName" json:"VolumeName"`
	}
	Kind string
}

func NewPersistentVolumeClaim() *ResPersistentVolumeClaim {
	return &ResPersistentVolumeClaim{
		ApiVersion: "v1",
		Kind:       RESOURCE_PERSISTENT_VOLUME_CLAIM,
	}
}

func (r *ResPersistentVolumeClaim) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

func (r *ResPersistentVolumeClaim) GetAccessModes() string {
	ac := r.Spec.AccessModes
	if len(ac) > 0 {
		return ac[0]
	}
	return ""
}

func (r *ResPersistentVolumeClaim) GetStorage() string {
	return r.Spec.Resources.Requests.Storage
}

func (r *ResPersistentVolumeClaim) GetStorageClassName() string {
	return r.Spec.StorageClassName
}

func (r *ResPersistentVolumeClaim) SetMetadataName(name string) error {
	r.Metadata.Name = name
	return nil
}

func (r *ResPersistentVolumeClaim) SetNamespace(ns string) error {
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResPersistentVolumeClaim) SetAccessMode(am string) error {
	r.Spec.AccessModes = append(r.Spec.AccessModes, am)
	return nil
}

func (r *ResPersistentVolumeClaim) SetStorage(storage string) error {
	r.Spec.Resources.Requests.Storage = storage
	return nil
}

func (r *ResPersistentVolumeClaim) SetVolumeName(vName string) error {
	r.Spec.VolumeName = vName
	return nil
}

func (r *ResPersistentVolumeClaim) SetVolumeMode(vName string) error {
	r.Spec.VolumeMode = vName
	return nil
}

func (r *ResPersistentVolumeClaim) SetStorageClassName(scName string) error {
	r.Spec.StorageClassName = scName
	return nil
}
