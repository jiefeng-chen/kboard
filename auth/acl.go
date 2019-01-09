package auth

const (
	ACL_USER_R = 1
	ACL_USER_W = 2
	ACL_USER_X = 4

	ACL_GROUP_R = 1
	ACL_GROUP_W = 2
	ACL_GROUP_X = 4

	ACL_OTEHR_R = 1
	ACL_OTHER_W = 2
	ACL_OTHER_X = 4
)

type acl interface {
	SetAcl()
	GetAcl()
}

type Acl struct {
}

func NewAcl() *Acl {
	return &Acl{}
}

func (a *Acl) SetAcl() {

}

func (a *Acl) GetAcl() {

}
