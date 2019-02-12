package auth

// ACLs typically associate an object with a list of users and groups
// for an object is a set of operations which may be performed on that object

//

const (
	ACL_R = 1
	ACL_W = 2
	ACL_X = 4
)

type acl interface {
	SetAcl()
	GetAcl()
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
