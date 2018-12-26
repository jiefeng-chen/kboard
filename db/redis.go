package db

import (
	"kboard/exception"
	"kboard/config"
	"github.com/mediocregopher/radix.v2/cluster"
	"fmt"
)


type RedisCluster struct {
	singleton *cluster.Cluster
}


func NewRedisCluster(config *config.Config) *RedisCluster {
	addr := fmt.Sprintf("%s:%d", config.Data.Mysql.Host, config.Data.Mysql.Port)
	clusterConn, err := cluster.New(addr)
	exception.CheckError(err, -1)

	return &RedisCluster{
		singleton: clusterConn,
	}
}




