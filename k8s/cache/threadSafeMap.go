package cache

import (
	"sync"
	"container/list"
)

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

	items *list.List // 链表

	len int // 当前长度

	cap int // 容量
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

func NewThreadSafeMap(cap int) IThreadSafeMap {
	return &ThreadSafeMap{
		cap: cap,
		len: 0,
		lock:  sync.RWMutex{},
		items:  list.New(),
	}
}