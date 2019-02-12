package cache

import (
	"testing"
	"strconv"
)

func TestNewThreadSafeMap(t *testing.T) {
	var threadMap IThreadSafeMap

	threadMap = NewThreadSafeMap(100)

	for i:= 0; i <= 1000; i++ {
		threadMap.Add(strconv.Itoa(i), i)
	}

	threadMap.Update("2", "10001")

	data, err := threadMap.Get("2")
	if err != nil {
		t.Fatal(err.Error())
	}else{
		t.Fatal(data)
	}

	t.Fatal(threadMap.List(), threadMap.Len())
}


