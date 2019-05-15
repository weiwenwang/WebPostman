package main

import (
	"github.com/gin-gonic/gin"
	"github.com/weiwenwang/WebPostman/controllers"
	"github.com/weiwenwang/WebPostman/controllers/Nav"
	"net/http"
	"github.com/weiwenwang/WebPostman/models"
)

func main() {
	db, err := models.InitDB()
	//db.LogMode(true) // 可以用来调试
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
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
