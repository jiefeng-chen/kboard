package auth

// ACLs typically associate an object with a list of users and groups
// for an object is a set of operations which may be performed on that object

//

const (
	ACL_R = 1 // 可读，可以查看，可以列出目录内容
	ACL_W = 2 // 可写，可以修改文件的内容
	ACL_X = 4 // 可执行，可以执行文件，可以进入目录
)

type acl interface {
	SetAcl()
	GetAcl()
}

type Acl struct {
	ResourceName string
	Owner string
	Group string
	Other string
}

func NewAcl(resName string) *Acl {
	return &Acl{
		ResourceName: resName,
	}
}

func (a *Acl) Init() {
	// 加载初始化参数

}

func (a *Acl) SetAcl() {

}

func (a *Acl) GetAcl() {

}
