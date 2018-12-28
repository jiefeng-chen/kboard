package db

import (
	"kboard/exception"
	"kboard/config"
	"time"
	"github.com/chasex/redis-go-cluster"
)


type RedisCluster struct {
	Singleton *redis.Cluster
}


func NewRedisCluster(config *config.Config) *RedisCluster {
	clusterConn, err := redis.NewCluster(&redis.Options{
		StartNodes: []string{
			"192.168.37.150:8001",
			"192.168.37.150:8002",
			"192.168.37.150:8003",
			"192.168.37.150:8004",
			"192.168.37.150:8005",
			"192.168.37.150:8006"},
		ConnTimeout: 50 * time.Millisecond,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
	})
	exception.CheckError(err, -1)

	return &RedisCluster{
		Singleton: clusterConn,
	}
}






