package Urlinfo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UrlinfoOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, gin.H{})
}
