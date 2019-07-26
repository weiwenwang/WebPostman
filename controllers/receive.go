package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/weiwenwang/WebRedis/common"
	"github.com/weiwenwang/WebRedis/models"
	"github.com/weiwenwang/WebRedis/redis"
	"net/http"
	"strconv"
)

func Receive(c *gin.Context) {
	select_host := c.Query("select_host")
	key := c.Query("key")
	//resp, err := http.Get(url)
	fmt.Println(select_host)
	ttl, _ := redis2.Int64(redis.Redis_fd.Do("TTL", key))
	ttl_string := strconv.FormatInt(ttl, 10)
	b, _ := redis2.String(redis.Redis_fd.Do("TYPE", key))
	fmt.Println(b)
	if (b == "string") {
		value, _ := redis2.String(redis.Redis_fd.Do("GET", key))
		fmt.Println("ttl", ttl)
		models.AddKeys(key)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.STRING,
			"VALUE": value,
			"TTL":   ttl_string,
		})
	}
	if (b == "set") {
		value, _ := redis2.Strings(redis.Redis_fd.Do("SMEMBERS", key))
		fmt.Println(value)
		fmt.Println("ttl", ttl)
		models.AddKeys(key)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.SET,
			"VALUE": value,
			"TTL":   ttl_string,
		})
	}
	if (b == "list") {
		value, _ := redis2.Strings(redis.Redis_fd.Do("LRANGE", key, 0, -1))
		len, _ := redis2.String(redis.Redis_fd.Do("LLEN", key))
		fmt.Println(value)
		fmt.Println("ttl", ttl)
		fmt.Println("len", len)
		models.AddKeys(key)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.LIST,
			"VALUE": value,
			"TTL":   ttl_string,
			"LEN":   len,
		})
	}
	if (b == "hash") {
		value, _ := redis2.Strings(redis.Redis_fd.Do("HGETALL", key))
		fmt.Println(value)
		var ret_str string
		for k, v := range value {
			if (k%2 == 1) {
				ret_str += v + "<br>"
			} else {
				ret_str += v + ":"
			}
		}
		fmt.Println("ttl", ttl)
		models.AddKeys(key)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.HASH,
			"VALUE": ret_str,
			"TTL":   ttl_string,
		})
	}

	if (b == "zset") {
		value, _ := redis2.Strings(redis.Redis_fd.Do("ZRANGE", key, 0, 100, "WITHSCORES"))
		fmt.Println(value)
		var ret_str string
		for k, v := range value {
			if (k%2 == 1) {
				ret_str += v + "<br>"
			} else {
				ret_str += v + ":"
			}
		}
		models.AddKeys(key)
		c.JSON(http.StatusOK, gin.H{
			"TYPE":  common.ZSET,
			"VALUE": ret_str,
			"TTL":   ttl_string,
		})
	}

}
