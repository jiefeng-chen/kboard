package cache

// 本地对象存储操作

type IObjStore interface {

}


type ObjStore struct {

}

func NewObjStore() IObjStore {
	return &ObjStore{}
}
