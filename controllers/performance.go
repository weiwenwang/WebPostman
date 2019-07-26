package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiwenwang/WebRedis/common"
	"github.com/weiwenwang/WebRedis/models/Clients"
	"net/http"
)

func Perform(c *gin.Context) {
	//str, _ := redis2.String(redis.Redis_fd.Do("CLIENT", "LIST"))
	//arr := strings.Split(str, "\n")
	ret, _ := Clients.Info()
	fmt.Println(ret)
	var Count_arr []int
	for _, v := range (ret) {
		Count_arr = append(Count_arr, v.Count)
	}
	//fmt.Println(arr)saQwtyuiop.[/

	//fmt.Println(len(arr))
	c.JSON(http.StatusOK, gin.H{
		"TYPE":  common.STRING,
		"VALUE": Count_arr,
		"TTL":   1,
	})

}
