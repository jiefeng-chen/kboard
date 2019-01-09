package cache

import "sync"

// thread safe map
type IThreadSafeMap interface {
	Set(key string, value interface{}) error
	Get(key string) (item interface{}, exist bool)
	List() []interface{}
	Clear()
	Exist()
	Size() int
}

type ThreadSafeMap struct {
	lock sync.RWMutex

	items map[string]interface{}


}

func (t *ThreadSafeMap) Set(key string, value interface{}) error {
	// LRU实现
}

func (t *ThreadSafeMap) Get(key string) (item interface{}, exist bool) {

}

func (t *ThreadSafeMap) List() []interface{} {

}

// 清理
func (t *ThreadSafeMap) Clear() {

}

func (t *ThreadSafeMap) Exist()  {

}

func (t *ThreadSafeMap) Size() int {

}

func NewThreadSafeMap() IThreadSafeMap {
	return &ThreadSafeMap{

	}
}