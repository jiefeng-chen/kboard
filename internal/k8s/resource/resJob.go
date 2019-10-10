package resource

import (
	"kboard/internal"

	"gopkg.in/yaml.v2"
)

type IJob interface {
	internal.IResource
	SetMetadataName(string) error
	GetNamespace() string
	SetNamespace(string) error
	AddContainer(*internal.Container) error
	SetCompletion(int) error
	SetTemplateName(string) error
	SetRestartPolicy(string) error
}

type ResJob struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	MetaData   struct {
		Name      string
		Namespace string
	}
	Spec *JobSpec
}

type JobSpec struct {
	Completions int // 固定结束次数
	Template    *JobTemplate
}

type JobTemplate struct {
	Metadata struct {
		Name string
	}
	Spec *JobTemplateSpec
}

type JobTemplateSpec struct {
	Container     []*internal.Container
	RestartPolicy string `yaml:"restartPolicy"`
}

func NewJob() *ResJob {
	return &ResJob{
		Kind:       internal.RESOURCE_JOB,
		ApiVersion: "batch/v1",
		MetaData: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Spec: &JobSpec{
			Completions: 0,
			Template: &JobTemplate{
				Metadata: struct{ Name string }{Name: ""},
				Spec: &JobTemplateSpec{
					Container:     nil,
					RestartPolicy: "",
				},
			},
		},
	}
}

func (r *ResJob) SetMetadataName(name string) error {
	if name == "" {
		return internal.NewError("name is empty")
	}
	r.MetaData.Name = name
	return nil
}

func (r *ResJob) GetNamespace() string {
	return r.MetaData.Namespace
}

func (r *ResJob) SetNamespace(ns string) error {
	if ns == "" {
		return internal.NewError("namespace is empty")
	}
	r.MetaData.Namespace = ns
	return nil
}

func (r *ResJob) AddContainer(container *internal.Container) error {
	if container == nil {
		return internal.NewError("container is nil")
	}
	r.Spec.Template.Spec.Container = append(r.Spec.Template.Spec.Container, container)
	return nil
}

func (r *ResJob) SetCompletion(t int) error {
	if t <= 0 {
		return internal.NewError("completion is zero")
	}
	r.Spec.Completions = t
	return nil
}

func (r *ResJob) SetTemplateName(name string) error {
	if name == "" {
		return internal.NewError("name is empty")
	}
	r.Spec.Template.Metadata.Name = name
	return nil
}

func (r *ResJob) SetRestartPolicy(policy string) error {
	if policy == "" {
		return internal.NewError("restart policy is empty")
	}
	r.Spec.Template.Spec.RestartPolicy = policy
	return nil
}

func (r *ResJob) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
