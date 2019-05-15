package Nav

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/weiwenwang/WebPostman/models/Url"
	"strconv"
)

func Info(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	nav_id := c.Query("nav_id")
	b, _ := strconv.Atoi(nav_id)
	c.JSON(http.StatusOK, gin.H{
		"data":  Url.Info(b),
		"errno": 0,
	})
}
