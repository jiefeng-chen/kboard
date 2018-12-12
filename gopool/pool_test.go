package gopool

import (
	"testing"
	"fmt"
)

func TestNewPool(t *testing.T) {
	pool := NewPool(1000)

	go func() {
		for i := 0; i < 100000; i++{
			pool.AddTask(NewTask(taskFunc, callbackFunc, i))
		}
	}()

	pool.Run()

	pool.GetRunTime()
}

func taskFunc(args interface{}) (error, interface{}) {
	fmt.Printf("task %d completed", args)
	return nil, args
}

func callbackFunc(result interface{}) (error, interface{}) {
	// 处理
	fmt.Printf("callback completed [%d]", result)
	return nil, result
}
