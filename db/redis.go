package db

import (
	"kboard/exception"
	"kboard/config"
	"github.com/garyburd/redigo/redis"
	"fmt"
)


type RedisCluster struct {
	Singleton redis.Conn
}


func NewRedisCluster(config *config.Config) *RedisCluster {
	addr := fmt.Sprintf("%s:%d", config.Data.Mysql.Host, config.Data.Mysql.Port)
	clusterConn, err := redis.Dial("tcp", addr)
	exception.CheckError(err, -1)

	return &RedisCluster{
		Singleton: clusterConn,
	}
}






