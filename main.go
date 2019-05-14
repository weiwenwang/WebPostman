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
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/static", "./static")
	r.GET("/", controllers.Index)
	r.GET("/index", controllers.Index)
	r.GET("/receive", controllers.Receive)

	r.GET("/urlinfo", controllers.Urlinfo)
	r.GET("/navlist", Nav.List)

	http.ListenAndServe(":8005", r)
	r.Run()
}
