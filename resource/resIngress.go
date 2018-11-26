package resource

import (
	"dashboard/core"
	"errors"
	"gopkg.in/yaml.v2"
)

type IResIngress interface {
	IResource
	SetMetaDataName(string) error
	GetMetaDataName() string
	SetNamespace(string) error
	SetRules(string, []map[string]string) error
	SetAnnotations(string) error
	SetLabels([]map[string]string) error
}

type ResIngress struct {
	Kind       string
	ApiVersion string `yaml:"apiVersion"`
	MetaData   struct {
		Name        string
		Namespace   string
		Annotations map[string]string
		Labels      map[string]string
	}
	Spec struct {
		Rules []*IngressRule
	}
}

type IngressRule struct {
	Host string
	Http struct {
		Paths []IngressPath
	}
}

type IngressPath struct {
	Path    string
	Backend IngressBackend
}

type IngressBackend struct {
	ServiceName string `yaml:"serviceName"`
	ServicePort int    `yaml:"servicePort"`
}

const ANNOTATIONS_INGRESS_CLASS = "kubernetes.io/ingress.class"

func NewIngress() *ResIngress {
	return &ResIngress{
		Kind:       RESOURCE_INGRESS,
		ApiVersion: "extensions/v1beta1",
		MetaData: struct {
			Name        string
			Namespace   string
			Annotations map[string]string
			Labels      map[string]string
		}{Name: "", Namespace: "", Annotations: map[string]string{}, Labels: map[string]string{}},
	}
}

func (r *ResIngress) SetAnnotations(annot map[string]string) error {
	if len(annot) <= 0 {
		return errors.New("Annotations is empty")
	}
	for _, v := range annot {
		r.MetaData.Annotations[ANNOTATIONS_INGRESS_CLASS] = v
	}
	return nil
}

func (r *ResIngress) SetMetaDataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.MetaData.Name = name
	return nil
}

func (r *ResIngress) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.MetaData.Namespace = ns
	return nil
}

func (r *ResIngress) GetMetaDataName() string {
	return r.MetaData.Name
}

func (r *ResIngress) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}

func (r *ResIngress) SetRules(host string, rules []map[string]string) error {
	if len(rules) > 0 {
		rule := new(IngressRule)
		rule.Host = host
		for _, v := range rules {
			path := new(IngressPath)
			if v["serviceName"] == "" || v["servicePort"] == "" {
				return errors.New("服务名称或服务端口为空")
			}
			if v["path"] != "" {
				// 这里允许访问路径为空，因为可以直接通过域名访问
				path.Path = v["path"]
			}
			path.Backend = IngressBackend{
				ServiceName: v["serviceName"],
				ServicePort: core.ToInt(v["servicePort"]),
			}
			rule.Http.Paths = append(rule.Http.Paths, *path)
		}
		r.Spec.Rules = append(r.Spec.Rules, rule)
	}
	return nil
}

func (r *ResIngress) SetLabels(labels []map[string]string) error {
	if len(labels) > 0 {
		for _, v := range labels {
			for key, val := range v {
				r.MetaData.Labels[key] = val
			}
		}
	}
	return nil
}
