package db

import (
	"fmt"
	"kboard/config"
	"kboard/internal"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestNewRedisCluster(t *testing.T) {
	var redisCluster *internal.RedisCluster
	config := config.NewConfig().LoadConfigFile("../config/conf.toml")
	redisCluster = internal.NewRedisCluster(config)

	v, err := redisCluster.Singleton.Do("SET", "name", "red")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	v, err = redis.String(redisCluster.Singleton.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	defer redisCluster.Singleton.Close()
}
