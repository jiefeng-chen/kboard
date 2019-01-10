package cache

import (
	"kboard/exception"
	"sync"
)

// thread safe map
type IThreadSafeMap interface {
	Add(string, interface{}) error
	Get(string) (interface{}, bool)

	CleanAll() error
	Exist(string) bool
	Size() int

	Delete(string) error
	Update(string, interface{}) error
	Replace(string, interface{}) error
	List() []interface{}
}

type ThreadSafeMap struct {
	lock sync.RWMutex

	items map[string]interface{} // 资源链表

	len int // 当前长度
}

func (t *ThreadSafeMap) Add(key string, item interface{}) error {
	if key == "" || item == nil {
		return exception.NewError("set value error: key or value is empty")
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	// 判断是否存在
	t.items[key] = item

	return nil
}

func (t *ThreadSafeMap) Get(key string) (interface{}, bool) {
	if key == "" {
		return nil, false
	}

	if t.Size() <= 0 {
		return nil, false
	}

	t.lock.RLock()
	defer t.lock.RUnlock()

	ele, ok := t.items[key]
	return ele, ok
}

func (t *ThreadSafeMap) List() []interface{} {
	t.lock.RLock()
	defer t.lock.RUnlock()

	data := make([]interface{}, 0, len(t.items))

	for _, v := range t.items {
		data = append(data, v)
	}

	return data
}

// 清理
func (t *ThreadSafeMap) CleanAll() error {
	if t.Size() <= 0 {
		return nil
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	for k, _ := range t.items {
		delete(t.items, k)
		t.len--
	}

	return nil
}

func (t *ThreadSafeMap) Exist(key string) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if _, ok := t.items[key]; ok {
		return true
	}

	return false
}

func (t *ThreadSafeMap) Delete(key string) error {
	if key == "" {
		return exception.NewError("key is empty")
	}
	if t.Size() <= 0 {
		return exception.NewError("list is empty")
	}
	t.lock.Lock()
	defer t.lock.Unlock()

	delete(t.items, key)
	t.len--

	return nil
}

func (t *ThreadSafeMap) Update(key string, item interface{}) error {
	if key == "" {
		return exception.NewError("key is empty")
	}
	t.lock.Lock()
	defer t.lock.Unlock()

	if !t.Exist(key) {
		return exception.NewError("key[%s] not exist", key)
	}

	t.items[key] = item

	return nil
}


func (t *ThreadSafeMap) Replace(key string, item interface{}) error {
	if key == "" {
		return exception.NewError("key is empty")
	}
	t.lock.Lock()
	defer t.lock.Unlock()

	if !t.Exist(key) {
		return exception.NewError("key[%s] not exist", key)
	}

	t.items[key] = item

	return nil
}

func (t *ThreadSafeMap) Size() int {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.len
}

func NewThreadSafeMap() IThreadSafeMap {
	return &ThreadSafeMap{
		len:   0,
		lock:  sync.RWMutex{},
		items: make(map[string]interface{}),
	}
}
