package resource

type IResPodPreset interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetRestartPolicy(string) error
	SetLabels(map[string]string) error
	AddContainer(*Container) error
	AddVolume(*Volume) error
	SetAnnotations(map[string]string) error
}

// pod结构体
type ResPodPreset struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Labels map[string]string
		Annotations map[string]string
	}
	Spec *ResPodPresetSpec
}

type ResPodPresetSpec struct {
	Containers []*Container
	RestartPolicy string `yaml:"restartPolicy"`  // [Always | Never | OnFailure]
	NodeSelector struct{} `yaml:"nodeSelector"`
	ImagePullSecrets []map[string]string `yaml:"imagePullSecrets"`
	HostNetwork bool `yaml:"hostNetwork"`
	Volumes []*Volume
}

func NewResPodPreset(name string) *ResPodPreset {
	return &ResPodPreset{
		ApiVersion: "settings.k8s.io/v1alpha1",
		Kind: RESOURCE_POD_PRESET,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      map[string]string
			Annotations map[string]string
		}{
			Name: name,
			Namespace: "",
			Labels: map[string]string{},
			Annotations: map[string]string{}},
		Spec: &ResPodPresetSpec{
			Containers: []*Container{},
			RestartPolicy: "",
			NodeSelector: struct{}{},
			ImagePullSecrets: []map[string]string{},
			HostNetwork: false,
			Volumes: []*Volume{},
		},
	}
}



