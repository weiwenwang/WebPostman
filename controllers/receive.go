package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/weiwenwang/WebRedis/common"
	"github.com/weiwenwang/WebRedis/redis"
	"net/http"
	"strconv"
)

func Receive(c *gin.Context) {
	method := c.Query("method")
	url := c.Query("url")
	//resp, err := http.Get(url)
	fmt.Println(method)
	ttl, _ := redis2.Int64(redis.Redis_fd.Do("TTL", url))
	ttl_string := strconv.FormatInt(ttl, 10)
	b, _ := redis2.String(redis.Redis_fd.Do("TYPE", url))
	fmt.Println(b)
	if (b == "string") {
		value, _ := redis2.String(redis.Redis_fd.Do("GET", url))
		ttl, _ := redis2.String(redis.Redis_fd.Do("TTL", url))
		fmt.Println("ttl", ttl)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.STRING,
			"VALUE": value,
			"TTL":   ttl_string,
		})
	}
	if (b == "set") {
		value, _ := redis2.Strings(redis.Redis_fd.Do("SMEMBERS", url))
		ttl, _ := redis2.String(redis.Redis_fd.Do("TTL", url))
		fmt.Println(value)
		fmt.Println("ttl", ttl)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.SET,
			"VALUE": value,
			"TTL":   ttl_string,
		})
	}
	if (b == "list") {
		value, _ := redis2.Strings(redis.Redis_fd.Do("LRANGE", url, 0, -1))
		len, _ := redis2.String(redis.Redis_fd.Do("LLEN", url))
		ttl, _ := redis2.String(redis.Redis_fd.Do("TTL", url))
		fmt.Println(value)
		fmt.Println("ttl", ttl)
		fmt.Println("len", len)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.LIST,
			"VALUE": value,
			"TTL":   ttl_string,
			"LEN":   len,
		})
	}
}
