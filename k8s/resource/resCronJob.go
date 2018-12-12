package resource

import "gopkg.in/yaml.v2"

type ICronJob interface {
	IResource
	SetMetaDataName(name string) bool
	GetMetaDataName() string
	SetNamespace(string) bool
}

type ResCronJob struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	MetaData   struct {
		Name      string
		Namespace string
	}
	Spec struct {
		Rules []IngressRule
	}
}

func NewCronJob() *ResCronJob {
	return &ResCronJob{
		Kind:       RESOURCE_CRON_JOB,
		ApiVersion: "extensions/v1beta1",
		MetaData: struct {
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
