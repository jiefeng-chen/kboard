package resource

import "github.com/golang/go/src/pkg/errors"

// pod结构体
type Pod struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind string
	Metadata struct{
		Name string
		Namespace string
		Labels []map[string]string
		Annotations []map[string]string
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

func NewPod(name string) *Pod {
	return &Pod{
		ApiVersion: "v1",
		Kind: RESOURCE_POD,
		Metadata: struct {
			Name        string
			Namespace   string
			Labels      []map[string]string
			Annotations []map[string]string
		}{Name: name, Namespace: "", Labels: []map[string]string{}, Annotations: []map[string]string{}},
	}
}

func (r *Pod) SetMetadataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Name = name
	return nil
}


func (r *Pod) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *Pod) SetRestartPolicy(policy string) error {
	if policy == "" {
		return errors.New("policy is empty")
	}
	r.Spec.RestartPolicy = policy
	return nil
}

func (r *Pod) AddContainer(container *Container) error {
	if container == nil {
		return errors.New("container is nil")
	}
	r.Spec.Containers = append(r.Spec.Containers, container)
	return nil
}

func (r *Pod) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return errors.New("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return errors.New("labels key or val is empty")
		}
		anno := map[string]string{
			"key": k,
			"name": v,
		}
		r.Metadata.Labels = append(r.Metadata.Labels, anno)
	}
	return nil
}

func (r *Pod) AddVolume(vol *Volume) error {
	if vol == nil {
		return errors.New("volume is nil")
	}
	r.Spec.Volumes = append(r.Spec.Volumes, vol)
	return nil
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

func (r *Pod) SetAnnotations(annos map[string]string) error {
	if len(annos) <= 0 {
		return errors.New("annotations is empty")
	}
	for k, v := range annos {
		if k == "" || v == "" {
			return errors.New("annotation key or val is empty")
		}
		anno := map[string]string{
			"key": k,
			"name": v,
		}
		r.Metadata.Annotations = append(r.Metadata.Annotations, anno)
	}
	return nil
}




