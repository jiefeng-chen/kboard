package resource

import (
	"gopkg.in/yaml.v2"
	"kboard/internal"
)

type IResPersistentVolume interface {
	internal.IResource
	SetMetaDataName(string) error
	SetCapacityStorage(string) error
	SetNamespace(string) error
	SetVolumeMode(string) error
	SetAccessModes([]string) error
	SetPersistentVolumeReclaimPolicy(string) error
	SetStorageClassName(string) error
	SetRbd(*Rbd) error
	SetClaimRef(*PersistentVolumeClaimRef) error
}

type ResPersistentVolume struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	}
	Spec struct {
		Capacity struct {
			Storage string
		}
		VolumeMode                    string   `yaml:"volumeMode"`
		AccessModes                   []string `yaml:"accessModes"`
		PersistentVolumeReclaimPolicy string   `yaml:"persistentVolumeReclaimPolicy"`
		StorageClassName              string   `yaml:"storageClassName"`
		Rbd                           *Rbd
		ClaimRef                      *PersistentVolumeClaimRef `yaml:"claimRef"`
	}
}

type PersistentVolumeClaimRef struct {
	Kind       string
	Namespace  string
	Name       string
	Uid        string
	ApiVersion string `yaml:"apiVersion"`
}

type Rbd struct {
	Monitors  []string
	Pool      string
	Image     string
	User      string
	SecretRef struct {
		Name string
	}
	FsType   string `yaml:"fsType"`
	ReadOnly bool   `yaml:"readOnly"`
	Keyring  string `yaml:"keyring"`
}

func NewPersistentVolume() *ResPersistentVolume {
	return &ResPersistentVolume{
		ApiVersion: "v1",
		Kind:       internal.RESOURCE_PERSISTENT_VOLUME,
	}
}

func (r *ResPersistentVolume) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

func (r *ResPersistentVolume) SetMetaDataName(name string) error {
	r.Metadata.Name = name
	return nil
}

func (r *ResPersistentVolume) SetNamespace(ns string) error {
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResPersistentVolume) SetCapacityStorage(s string) error {
	r.Spec.Capacity.Storage = s
	return nil
}

func (r *ResPersistentVolume) SetVolumeMode(vMode string) error {
	r.Spec.VolumeMode = vMode
	return nil
}

func (r *ResPersistentVolume) SetAccessModes(aModes []string) error {
	r.Spec.AccessModes = aModes
	return nil
}

func (r *ResPersistentVolume) SetPersistentVolumeReclaimPolicy(pvRP string) error {
	r.Spec.PersistentVolumeReclaimPolicy = pvRP
	return nil
}

func (r *ResPersistentVolume) SetStorageClassName(scName string) error {
	r.Spec.StorageClassName = scName
	return nil
}

func (r *ResPersistentVolume) SetRbd(rbd *Rbd) error {
	r.Spec.Rbd = rbd
	return nil
}

func (r *ResPersistentVolume) SetClaimRef(ref *PersistentVolumeClaimRef) error {
	r.Spec.ClaimRef = ref
	return nil
}
