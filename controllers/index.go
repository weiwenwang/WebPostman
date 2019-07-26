package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Node struct {
	Title string
	Url   string
	Sub   []*Node
}

func Index(c *gin.Context) {
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

	fathers = append(fathers, &father)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"title": "WebRedis",
		"nav":   fathers,
	})
}
