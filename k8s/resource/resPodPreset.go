package resource

import "gopkg.in/yaml.v2"

type IResPodPreset interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetRestartPolicy(string) error
	SetLabels(map[string]string) error
	AddContainer(IContainer) error
	AddVolume(*Volume) error
	SetAnnotations(map[string]string) error
}

// pod结构体
type ResPodPreset struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name        string
		Namespace   string
		Labels      map[string]string
		Annotations map[string]string
	}
	Spec *ResPodPresetSpec
}

type ResPodPresetSpec struct {
	Containers       []IContainer
	RestartPolicy    string              `yaml:"restartPolicy"` // [Always | Never | OnFailure]
	NodeSelector     struct{}            `yaml:"nodeSelector"`
	ImagePullSecrets []map[string]string `yaml:"imagePullSecrets"`
	HostNetwork      bool                `yaml:"hostNetwork"`
	Volumes          []*Volume
}

func NewResPodPreset(name string) IResPodPreset {
	return &ResPodPreset{
		ApiVersion: "settings.k8s.io/v1alpha1",
		Kind:       RESOURCE_POD_PRESET,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      map[string]string
			Annotations map[string]string
		}{
			Name:        name,
			Namespace:   "",
			Labels:      map[string]string{},
			Annotations: map[string]string{}},
		Spec: &ResPodPresetSpec{
			Containers:       []IContainer{},
			RestartPolicy:    "",
			NodeSelector:     struct{}{},
			ImagePullSecrets: []map[string]string{},
			HostNetwork:      false,
			Volumes:          []*Volume{},
		},
	}
}

func (r *ResPodPreset) SetMetadataName(string) error {
	return nil
}

func (r *ResPodPreset) SetNamespace(string) error {
	return nil
}


func (r *ResPodPreset) SetRestartPolicy(string) error {
	return nil
}

func (r *ResPodPreset) SetLabels(map[string]string) error {
	return nil
}

func (r *ResPodPreset) AddContainer(IContainer) error {
	return nil
}

func (r *ResPodPreset) AddVolume(*Volume) error {
	return nil
}

func (r *ResPodPreset) SetAnnotations(map[string]string) error {
	return nil
}

func (r *ResPodPreset) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
