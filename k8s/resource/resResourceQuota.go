package resource


import (
	"gopkg.in/yaml.v2"
	"github.com/golang/go/src/pkg/errors"
)

// 需要启动时配置ResourceQuota adminssion control
type IResResourceQuota interface {
	IResource
	SetMetaDataName(name string) error
	GetNamespace() string
	SetNamespace(string) error
}

type ResResourceQuota struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	MetaData   struct {
		Name      string
		Namespace string
	}
	Spec struct{
		Hard *Hard
	}
}

type Hard struct {
	Configmaps string
	Persistentvolumeclaims string
	Replicationcontrollers string
	Secrets string
	Services string
	Pods string
	Requests *Request
	Limits *Limits
}

type Limits struct {
	Cpu string
	Memory string
}

type Request struct {
	Cpu string
	Memory string
}

func NewResResourceQuota() *ResResourceQuota {
	return &ResResourceQuota{
		Kind:       RESOURCE_RESOURCE_QUOTA,
		ApiVersion: "v1",
		MetaData: struct {
			Name      string
			Namespace string
		}{Name: "", Namespace: ""},
		Spec: struct{ Hard *Hard }{Hard: &Hard{
			Requests: &Request{
				Cpu: "",
				Memory: "",
			},
			Limits: &Limits{
				Cpu: "",
				Memory: "",
			},
		}},
	}
}

func (r *ResResourceQuota) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.MetaData.Namespace = ns
	return nil
}

func (r *ResResourceQuota) SetMetaDataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.MetaData.Name = name
	return nil
}

func (r *ResResourceQuota) GetNamespace() string {
	return r.MetaData.Namespace
}

func (r *ResResourceQuota) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
