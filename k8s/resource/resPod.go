package resource

import (
	"github.com/golang/go/src/pkg/errors"
	"gopkg.in/yaml.v2"
	"kboard/exception"
)

type IResPod interface {
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
type ResPod struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Labels map[string]string
		Annotations map[string]string
	}
	Spec struct{
		Containers []*Container
		RestartPolicy string `yaml:"restartPolicy"`  // [Always | Never | OnFailure]
		NodeSelector struct{} `yaml:"nodeSelector"`
		ImagePullSecrets []map[string]string `yaml:"imagePullSecrets"`
		HostNetwork bool `yaml:"hostNetwork"`
		Volumes []*Volume
	}
}

func NewResPod(name string) *ResPod {
	return &ResPod{
		ApiVersion: "v1",
		Kind: RESOURCE_POD,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      map[string]string
			Annotations map[string]string
		}{Name: name, Namespace: "", Labels: map[string]string{}, Annotations: map[string]string{}},
		Spec: struct {
			Containers       []*Container
			RestartPolicy    string
			NodeSelector     struct{}
			ImagePullSecrets []map[string]string
			HostNetwork      bool
			Volumes          []*Volume
		}{
			Containers: []*Container{},
			RestartPolicy: "",
			NodeSelector: struct{}{},
			ImagePullSecrets: []map[string]string{},
			HostNetwork: false,
			Volumes: []*Volume{}},
	}
}

func (r *ResPod) SetMetadataName(name string) error {
	if name == "" {
		return exception.NewError("name is empty")
	}
	r.Metadata.Name = name
	return nil
}


func (r *ResPod) SetNamespace(ns string) error {
	if ns == "" {
		return exception.NewError("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResPod) SetRestartPolicy(policy string) error {
	if policy == "" {
		return exception.NewError("policy is empty")
	}
	r.Spec.RestartPolicy = policy
	return nil
}

func (r *ResPod) AddContainer(container *Container) error {
	if container == nil {
		return exception.NewError("container is nil")
	}
	r.Spec.Containers = append(r.Spec.Containers, container)
	return nil
}

func (r *ResPod) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return exception.NewError("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return exception.NewError("labels key or val is empty")
		}

		r.Metadata.Labels[k] = v
	}
	return nil
}

func (r *ResPod) AddVolume(vol *Volume) error {
	if vol == nil {
		return exception.NewError("volume is nil")
	}
	r.Spec.Volumes = append(r.Spec.Volumes, vol)
	return nil
}


func (r *ResPod) SetAnnotations(annos map[string]string) error {
	if len(annos) <= 0 {
		return errors.New("annotations is empty")
	}
	for k, v := range annos {
		if k == "" || v == "" {
			return exception.NewError("annotation key or val is empty")
		}
		r.Metadata.Annotations[k] = v
	}
	return nil
}


func (r *ResPod) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

func NewVolume() *Volume {
	return &Volume{
		Name:"",
		EmptyDir: struct {

		}{},
		HostPath: struct{ Path string }{Path: ""},
		Secret: &Secret{SecretName: "", Items: []map[string]string{}},
		ConfigMap: struct {
			Name  string
			Items []map[string]string
		}{Name: "", Items: []map[string]string{}},
	}
}

type Volume struct {
	Name string
	EmptyDir interface{} `yaml:"emptyDir"`
	HostPath struct{
		Path string
	} `yaml:"hostPath"`
	Secret *Secret
	ConfigMap struct{
		Name string
		Items []map[string]string // [key:string, path:string]
	} `yaml:"configMap"`
}


