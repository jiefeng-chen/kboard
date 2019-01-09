package test

import (
	"testing"
	"time"

	"github.com/chasex/redis-go-cluster"
)

func TestRedisNode(t *testing.T) {

}

func TestRedisConn(t *testing.T) {
	clusterConn, err := redis.NewCluster(&redis.Options{
		StartNodes: []string{
			"192.168.37.150:8001",
			"192.168.37.150:8002",
			"192.168.37.150:8003",
			"192.168.37.150:8004",
			"192.168.37.150:8005",
			"192.168.37.150:8006"},
		ConnTimeout:  50 * time.Millisecond,
		ReadTimeout:  50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive:    16,
		AliveTime:    50 * time.Second,
	})

	if err != nil {
		t.Fatalf("redis.New error: %s", err.Error())
	}

	_, err = clusterConn.Do("SET", "myname", "realjf")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	data, err := clusterConn.Do("GET", "myname")
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	defer clusterConn.Close()

	t.Errorf("%v", data)

}
