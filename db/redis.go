package db

import (
	"github.com/chasex/redis-go-cluster"
	"kboard/exception"
	"time"
)

type redisCluster interface {
	Set(string, string) error
}

type RedisCluster struct {
	singleton *redis.Cluster
}


func NewRedisCluster() *RedisCluster {
	cluster, err := redis.NewCluster(&redis.Options{
		StartNodes: []string{}, // 集群节点
		ConnTimeout: 50 * time.Microsecond,
		ReadTimeout: 50 * time.Microsecond,
		WriteTimeout: 50 * time.Microsecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
	})
	exception.CheckError(err, -1)
	return &RedisCluster{
		singleton: cluster,
	}
}

func (rc *RedisCluster) Set() error {

}

func (rc *RedisCluster) Set() error {

}

