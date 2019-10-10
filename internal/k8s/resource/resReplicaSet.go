package resource

import (
	"errors"
	"kboard/internal"

	"gopkg.in/yaml.v2"
)

type IResReplicaSet interface {
	internal.IResource
	SetMetadataName(string) error
	SetNamespace(string) error
	GetNamespace() string
	SetLabels(map[string]string) error
	GetLabel(string) string
	SetTemplateLabel(map[string]string) error
	AddContainer(internal.IContainer) error
	SetReplicas(int) error
	SetSelector(internal.Selector) error
}

type ResReplicaSet struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
		Namespace string
		Labels    map[string]string // 标签组
	}
	Spec struct {
		Replicas int
		Selector internal.Selector
		Template ReplicaSetTemplate
	}
}

type ReplicaSetTemplate struct {
	Metadata struct {
		Labels map[string]string
	}
	Spec struct {
		Containers []internal.IContainer
	}
}

func NewResReplicaSet() *ResReplicaSet {
	return &ResReplicaSet{
		ApiVersion: "apps/v1",
		Kind:       internal.RESOURCE_REPLICASET,
		Metadata: struct {
			Name      string
			Namespace string
			Labels    map[string]string
		}{
			Name:      "",
			Namespace: "",
			Labels:    map[string]string{},
		},
		Spec: struct {
			Replicas int
			Selector internal.Selector
			Template ReplicaSetTemplate
		}{
			Replicas: 0,
			Selector: internal.Selector{
				MatchLabels:      map[string]string{},
				MatchExpressions: nil,
			},
			Template: ReplicaSetTemplate{
				Metadata: struct{ Labels map[string]string }{
					Labels: map[string]string{}},
				Spec: struct{ Containers []internal.IContainer }{
					Containers: []internal.IContainer{}},
			}},
	}
}

func (r *ResReplicaSet) SetMetadataName(name string) error {
	if name == "" {
		return errors.New("metadata name is empty")
	}
	r.Metadata.Name = name
	return nil
}

func (r *ResReplicaSet) SetNamespace(ns string) error {
	if ns == "" {
		return errors.New("namespace is empty")
	}
	r.Metadata.Namespace = ns
	return nil
}

func (r *ResReplicaSet) GetNamespace() string {
	return r.Metadata.Namespace
}

func (r *ResReplicaSet) SetLabels(data map[string]string) error {
	if len(data) > 0 {
		for k, v := range data {
			if k == "" || v == "" {
				return errors.New("label key or value is empty")
			}
			r.Metadata.Labels[k] = v
		}

		return nil
	} else {
		return errors.New("no labels will be set")
	}
}

func (r *ResReplicaSet) SetSelector(selector internal.Selector) error {
	if len(selector.MatchLabels) <= 0 {
		return errors.New("selector's match labels is empty")
	}
	r.Spec.Selector = selector
	r.Spec.Template.Metadata.Labels = selector.MatchLabels
	return nil
}

func (r *ResReplicaSet) GetLabel(name string) string {
	return r.Metadata.Labels[name]
}

func (r *ResReplicaSet) SetTemplateLabel(labels map[string]string) error {
	if len(labels) <= 0 {
		return errors.New("labels is empty")
	}
	for k, v := range labels {
		r.Spec.Template.Metadata.Labels[k] = v
	}
	return nil
}

func (r *ResReplicaSet) AddContainer(container internal.IContainer) error {
	if container == nil {
		return errors.New("container is nil")
	}
	r.Spec.Template.Spec.Containers = append(r.Spec.Template.Spec.Containers, container)
	return nil
}

func (r *ResReplicaSet) SetReplicas(replica int) error {
	if replica <= 0 {
		return errors.New("replica is empty")
	}
	r.Spec.Replicas = replica
	return nil
}

func (r *ResReplicaSet) ToYamlFile() ([]byte, error) {
	yamlData, err := yaml.Marshal(*r)
	if err != nil {
		return []byte{}, err
	}
	return yamlData, nil
}
