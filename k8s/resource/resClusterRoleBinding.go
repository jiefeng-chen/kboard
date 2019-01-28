package resource


type IResClusterRoleBinding interface {
	IResource
	SetMetadataName(string) error
}

type ResClusterRoleBinding struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   struct {
		Name      string
	}
	Subjects []RoleBindingSubject
	RoleRef RoleRef
}


func NewResClusterRoleBinding() *ResClusterRoleBinding {
	return &ResClusterRoleBinding{
		ApiVersion: "rbac.authorization.k8s.io/v1",
		Kind:       RESOURCE_CLUSTER_ROLE_BINDING,
		Metadata: struct {
			Name      string
		}{Name: ""},
		Subjects: []RoleBindingSubject{},
		RoleRef: RoleRef{},
	}
}



