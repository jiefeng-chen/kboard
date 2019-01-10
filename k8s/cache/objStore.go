package cache

// 本地对象存储操作

type IObjStore interface {
}

type ObjStore struct {
	cache map[interface{}]IThreadSafeMap
}

func NewObjStore() IObjStore {
	return &ObjStore{
		cache: make(map[interface{}]IThreadSafeMap),
	}
}
