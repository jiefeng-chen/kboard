package resource

import "gopkg.in/yaml.v2"

type resPersistentVolume interface {
	SetMetaDataName(string) bool
	SetCapacityStorage(string) bool
	SetNamespace(string) bool
	SetVolumeMode(string) bool
	SetAccessModes([]string) bool
	SetPersistentVolumeReclaimPolicy(string) bool
	SetStorageClassName(string) bool
	SetRbd(*Rbd) bool
	SetClaimRef(*ClaimRef) bool
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
		ClaimRef                      *ClaimRef `yaml:"claimRef"`
	}
}

type ClaimRef struct {
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
		Kind:       "PersistentVolume",
	}
}

func (r *ResPersistentVolume) SetMetaDataName(name string) bool {
	r.Metadata.Name = name
	return true
}

func (r *ResPersistentVolume) SetNamespace(ns string) bool {
	r.Metadata.Namespace = ns
	return true
}

func (r *ResPersistentVolume) SetCapacityStorage(s string) bool {
	r.Spec.Capacity.Storage = s
	return true
}

func (r *ResPersistentVolume) SetVolumeMode(vMode string) bool {
	r.Spec.VolumeMode = vMode
	return true
}

func (r *ResPersistentVolume) SetAccessModes(aModes []string) bool {
	r.Spec.AccessModes = aModes
	return true
}

func (r *ResPersistentVolume) SetPersistentVolumeReclaimPolicy(pvRP string) bool {
	r.Spec.PersistentVolumeReclaimPolicy = pvRP
	return true
}

func (r *ResPersistentVolume) SetStorageClassName(scName string) bool {
	r.Spec.StorageClassName = scName
	return true
}

func (r *ResPersistentVolume) SetRbd(rbd *Rbd) bool {
	r.Spec.Rbd = rbd
	return true
}

func (r *ResPersistentVolume) SetClaimRef(ref *ClaimRef) bool {
	r.Spec.ClaimRef = ref
	return true
}

func (r *ResPersistentVolume) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
