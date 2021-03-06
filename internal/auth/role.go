package auth

import "kboard/internal"

type role interface {
	Create() error
	Delete() error
	Grant() error
	AssociateOp(internal.PermCode) error
	AssociateUser(string) error
}

// user, roles, operations, and sessions
// role/operation m:n
// user/role  m:n
// user/session 1:m
// session/role 1:1

const (
	// 角色身份
	ROLE_IDENTITY_SUPER = iota + 1 // SRE - 系统管理员
	ROLE_IDENTITY_PRO              // SRE - 项目管理员
	ROLE_IDENTITY_OPS              // SRE - 运维工程师
	ROLE_IDENTITY_DEV              // 研发工程师
	ROLE_IDENTITY_TEST             // 测试工程师

	ROLE_IDENTITY_PRO_ADMIN // 团队管理员
)

type Identity int

type Role struct {
	Identity Identity // 身份
	Name     string   // 名称
	SREFlag  bool     // 是否SRE团队成员
}

// 获取角色名称
func getRoleName(id Identity) string {
	roleNames := map[Identity]string{
		ROLE_IDENTITY_SUPER:     "系统管理员",
		ROLE_IDENTITY_PRO:       "项目管理员",
		ROLE_IDENTITY_OPS:       "运维工程师",
		ROLE_IDENTITY_DEV:       "研发工程师",
		ROLE_IDENTITY_TEST:      "测试工程师",
		ROLE_IDENTITY_PRO_ADMIN: "团队管理员",
	}

	if name, ok := roleNames[id]; ok {
		return name
	}

	return "未授权人员"
}

// SRE 标志
func getSreFlag(id Identity) bool {
	roleNames := map[Identity]bool{
		ROLE_IDENTITY_SUPER:     true,
		ROLE_IDENTITY_PRO:       true,
		ROLE_IDENTITY_OPS:       true,
		ROLE_IDENTITY_DEV:       false,
		ROLE_IDENTITY_TEST:      false,
		ROLE_IDENTITY_PRO_ADMIN: false,
	}

	if flag, ok := roleNames[id]; ok {
		return flag
	}

	return false
}

func NewRole(identity Identity) role {
	return &Role{
		Identity: identity,
		Name:     getRoleName(identity),
		SREFlag:  getSreFlag(identity),
	}
}

func (r *Role) Create() error {
	return nil
}

func (r *Role) Delete() error {
	return nil
}

func (r *Role) Grant() error {
	return nil
}

func (r *Role) AssociateOp(perm internal.PermCode) error {

	return nil
}

func (r *Role) AssociateUser(user string) error {

	return nil
}
