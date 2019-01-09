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
	GetStorage() string
}

type ResPersistentVolumeClaim struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
	}
	Spec *ResPVCSpec
}

type PVCResource struct {
	Requests *PVCRequest
}

type PVCRequest struct {
	Storage string
}

func NewPersistentVolumeClaim() *ResPersistentVolumeClaim {
	return &ResPersistentVolumeClaim{
		ApiVersion: "v1",
		Kind:       RESOURCE_PERSISTENT_VOLUME_CLAIM,
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Spec: &ResPVCSpec{
			AccessModes: []string{},
			Resources: &PVCResource{
				Requests: &PVCRequest{
					Storage: "",
				},
			},
			VolumeMode:       "",
			StorageClassName: "",
			VolumeName:       "",
		},
	}
}

type ResPVCSpec struct {
	AccessModes      []string `yaml:"accessModes" json:"accessModes"`
	Resources        *PVCResource
	VolumeMode       string `yaml:"volumeMode" json:"volumeMode"`
	StorageClassName string `yaml:"storageClassName" json:"storageClassName"`
	VolumeName       string `yaml:"volumeName" json:"VolumeName"`
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
