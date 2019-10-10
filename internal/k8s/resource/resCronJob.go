package resource

import (
	"kboard/internal"

	"gopkg.in/yaml.v2"
)

type ICronJob interface {
	internal.IResource
	SetMetadataName(name string) error
	GetMetaDataName() string
	SetNamespace(string) error
}

type ResCronJob struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	Metadata   struct {
		Name      string
		Namespace string
	}
}

func NewCronJob() *ResCronJob {
	return &ResCronJob{
		Kind:       internal.RESOURCE_CRON_JOB,
		ApiVersion: "extensions/v1beta1",
		Metadata: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
	}
}

func (r *ResCronJob) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

func (r *ResCronJob) SetMetadataName(name string) error {
	if name == "" {
		return internal.NewError("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResCronJob) GetMetaDataName() string {
	return r.Metadata.Name
}

func (r *ResCronJob) SetNamespace(ns string) error {
	if ns == "" {
		return internal.NewError("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}
