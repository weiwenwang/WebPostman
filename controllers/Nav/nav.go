package Nav

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/weiwenwang/WebRedis/models/Nav"
	"net/http"
)

func List(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"navlist": Nav.NavList(),
	})
}

func info() {

}
