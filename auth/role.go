package auth


type role interface {
	Create()
	Delete()
	Grant()
}

const (
	// 角色身份
	ROLE_IDENTIDY_SUPER = 1	 // 系统管理员
	ROLE_IDENTIDY_ADMIN = 2  // 项目管理员
	ROLE_IDENTIDY_OPS = 3 	 // 运维工程师
	ROLE_IDENTIDY_DEV = 4	 // 研发工程师
	ROLE_IDENTIDY_TEST = 5   // 测试工程师
)

type Identity int


type Role struct {
	Identity Identity // 身份
}

func NewRole(identity Identity) *Role {
	return &Role{
		Identity: identity,
	}
}


func (r *Role) Create() {

}

func (r *Role) Delete() {

}

func (r *Role) Grant() {

}

