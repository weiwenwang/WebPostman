package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var Redis_fd redis.Conn

func InitRedis() () {
	c, err := redis.Dial("tcp", "192.168.1.9:6379")
	if err != nil {
		fmt.Println(err)
	}
	Redis_fd = c
}
