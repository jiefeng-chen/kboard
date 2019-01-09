package resource

import (
	"errors"

	"gopkg.in/yaml.v2"
)

// 对pod所包含的所有container所需的cpu和内存资源进行管理
type IResLimitRange interface {
	IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	SetLabels(map[string]string) error
	AddLimits(*Limit) error
}

type ResLimitRange struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string
	}
	Spec struct {
		Limits []*Limit
	}
}

type Limit LimitContainer

type LimitPod struct {
	Max                  Limits
	Min                  Limits
	MaxLimitRequestRatio Limits `yaml:"maxLimitRequestRatio"`
	Type                 string
}

type LimitContainer struct {
	Default        Limits
	DefaultRequest Request `yaml:"defaultRequest"`
	LimitPod
}

func NewResLimitRange() *ResLimitRange {
	return &ResLimitRange{
		ApiVersion: "v1",
		Kind:       RESOURCE_LIMIT_RANGE,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{
			Name:      "",
			Namespace: "",
			Labels:    map[string]string{}},
		Spec: struct{ Limits []*Limit }{Limits: []*Limit{}},
	}
}

func (r *ResLimitRange) SetMetadataName(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResLimitRange) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResLimitRange) SetLabels(labels map[string]string) error {
	if len(labels) <= 0 {
		return errors.New("labels is empty")
	}
	for k, v := range labels {
		if k == "" || v == "" {
			return errors.New("label's key or value is empty")
		}
		if r.Metadata.Labels[k] != "" {
			return errors.New("label [" + k + "] is exist")
		}
		r.Metadata.Labels[k] = v
	}
	return nil
}

func (r *ResLimitRange) AddLimits(limit *Limit) error {
	if limit == nil {
		return errors.New("limit is nil")
	}
	r.Spec.Limits = append(r.Spec.Limits, limit)
	return nil
}

func (r *ResLimitRange) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
