package cache

// 本地对象存储操作

type IObjStore interface {

}


type ObjStore struct {
	resourceType string // 资源类型

}

func NewObjStore() IObjStore {
	return &ObjStore{}
}
