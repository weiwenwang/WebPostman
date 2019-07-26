package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/weiwenwang/WebRedis/models"
	"net/http"
)

type Node struct {
	Title string
	Url   string
	Sub   []*Node
}

func Index(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var fathers []*Node
	sub_1 := Node{
		"性能",
		"https://www.baidu.com",
		nil,
	}
	sub_2 := Node{
		"查询",
		"https://www.baidu.com",
		nil,
	}
	self := Node{
		"redis",
		"https://www.baidu.com",
		nil,
	}
	father := Node{
		"redis管理",
		"https://www.baidu.com",
		nil,
	}

	self.Sub = append(self.Sub, &sub_1)
	self.Sub = append(self.Sub, &sub_2)
	father.Sub = append(father.Sub, &self)
	keys := reverse(models.GetKeys())
	fmt.Println(keys)
	//[]string{"tms_string", "tms_list", "tms_zset"}
	fathers = append(fathers, &father)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"title": "WebRedis",
		"nav":   fathers,
		"keys":  keys,
	})
}

func reverse(s []string) []string {
	count := len(s)
	fmt.Println("count:", count)
	var ret []string
	for i := 0; i < count; i++ {
		ret = append(ret, s[count-i-1])
	}
	return ret
}
