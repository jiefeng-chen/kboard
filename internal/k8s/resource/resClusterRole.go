package resource

import "kboard/internal"

type IResClusterRole interface {
	internal.IResource
	SetMetadataName(string) error
}

type ResClusterRole struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		// "namespace" omitted since ClusterRoles are not namespaced
		Name string
	}
	Rules []ClusterRoleRule
}

type ClusterRoleRule struct {
	ApiGroups []string `yaml:"apiGroups"` // "" indicates the core API group
	Resources []string
	Verbs     []string
}

func NewResClusterRole() *ResClusterRole {
	return &ResClusterRole{
		ApiVersion: "rbac.authorization.k8s.io/v1",
		Kind:       internal.RESOURCE_CLUSTER_ROLE,
		Metadata: struct {
			Name string
		}{Name: ""},
		Rules: []ClusterRoleRule{},
	}
}
