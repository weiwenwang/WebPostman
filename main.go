package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	redis2 "github.com/gomodule/redigo/redis"
	"github.com/weiwenwang/WebRedis/controllers"
	"github.com/weiwenwang/WebRedis/controllers/Nav"
	"github.com/weiwenwang/WebRedis/models"
	"github.com/weiwenwang/WebRedis/redis"
	"net/http"
	"strings"
)

func main() {
	db, err := models.InitDB()
	redis.InitRedis()
	str, _ := redis2.String(redis.Redis_fd.Do("CLIENT", "LIST"))
	arr := strings.Split(str, "\n")

	//fmt.Println(arr)
	fmt.Println(len(arr))

	db.LogMode(true) // 可以用来调试
	if err != nil {
		panic(err)
	}

	//ret, err := Clients.Info()
	//fmt.Println(ret)
	//for _, v := range (ret) {
	//	fmt.Println(v.Count)
	//}

	//defer db.Close()
	//db.LogMode(true)
	//
	//a := Nav2.NavList()
	//fmt.Println(a)
	//os.Exit(1)
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/static", "./static")
	r.GET("/", controllers.Index)
	r.GET("/index", controllers.Index)
	r.GET("/receive", controllers.Receive)

	r.GET("/urlinfo", controllers.Urlinfo)
	r.GET("/navlist", Nav.List)
	r.GET("/navinfo", Nav.Info)

	http.ListenAndServe(":8005", r)
	r.Run()
}
