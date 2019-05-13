package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	url2 "net/url"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func Urlinfo(c *gin.Context) {
	method := c.Query("method")
	url := c.Query("url")
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println()

	fmt.Println(resp.Header);
	var header []map[string]string
	for k, v := range resp.Header {
		mp1 := make(map[string]string)
		mp1["key"] = k
		mp1["value"] = v[0]
		header = append(header, mp1)
	}

	fmt.Println(header)

	if err != nil {

	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url2.Parse(url)
	//fmt.Println(u.RawQuery)
	m, _ := url2.ParseQuery(u.RawQuery)
	content := make(map[string]interface{})
	json.Unmarshal(body, &content)

	var param []map[string]string
	for k, v := range m {
		mp1 := make(map[string]string)
		mp1["key"] = k
		mp1["value"] = v[0]
		param = append(param, mp1)
	}

	c.JSON(http.StatusOK, gin.H{
		"url":        url,
		"parameters": param,
		"method":     method,
		"header":     header,
		"content":    content,
	})
}
