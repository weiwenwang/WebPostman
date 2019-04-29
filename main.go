package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weiwenwang/WebPostman/controllers"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/static", "./static")
	r.GET("/", controllers.Index)
	r.GET("/index", controllers.Index)
	r.GET("/receive", controllers.Receive)
	r.GET("/urlinfo", controllers.Urlinfo)
	r.Run()
}
