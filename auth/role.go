package auth


type role interface {
	Create()
	Delete()
	Grant()
}

const (
	// 角色身份
	ROLE_IDENTIDY_SUPER = 1	// 超管
	ROLE_IDENTIDY_ADMIN = 2  // 管理员
	ROLE_IDENTIDY_MEMBER = 3 // 普通成员
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

