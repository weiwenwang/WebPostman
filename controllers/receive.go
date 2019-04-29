package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func Receive(c *gin.Context) {
	method := c.Query("method")
	url := c.Query("url")
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	header := make(map[string]string)
	vs := ""
	for k, v := range resp.Header {
		for _, vv := range v {
			vs = vv + ","
		}
		header[k] = vs
	}
	header_str, err := json.Marshal(header)

	if err != nil {
		// handle error
	}

	c.JSON(http.StatusOK, gin.H{
		"url":    url,
		"method": method,
		"body":   string(body),
		"header": string(header_str),
	})
}
