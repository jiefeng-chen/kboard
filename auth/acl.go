package auth

// acl 访问控制表
// ACLs typically associate an object with a list of users and groups
// for an object is a set of operations which may be performed on that object



const (
	ACL_USER_R = 1
	ACL_USER_W = 2
	ACL_USER_X = 4

	ACL_TEAM_R = 1
	ACL_TEAM_W = 2
	ACL_TEAM_X = 4

	ACL_OTEHR_R = 1
	ACL_OTHER_W = 2
	ACL_OTHER_X = 4
)

type acl interface {
	SetAcl() // 设置acl权限
	GetAcl() // 查看acl权限
}

type Acl struct {

}

func NewAcl( string) *Acl {
	return &Acl{

	}
}

func (a *Acl) SetAcl() {

}

func (a *Acl) GetAcl() {

}
