package Nav

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"github.com/weiwenwang/WebPostman/models/Nav"
)

//type Nav struct {
//	id   int32
//	name string
//
//	sub  *Nav
//}

func List(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{
		"navlist": Nav.NavList(),
	})
}

func info() {

}
