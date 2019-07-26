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
	select_host := c.Query("select_host")
	key := c.Query("key")
	resp, err := http.Get(key)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Header);
	var response_header [][]string
	for k, v := range resp.Header {
		var mp1 []string
		mp1 = append(mp1, k)
		mp1 = append(mp1, v[0])
		response_header = append(response_header, mp1)
	}

	if err != nil {

	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url2.Parse(key)
	//fmt.Println(u.RawQuery)
	m, _ := url2.ParseQuery(u.RawQuery)
	content := make(map[string]interface{})
	json.Unmarshal(body, &content)

	var param [][]string
	for k, v := range m {
		var mp1 []string
		mp1 = append(mp1, k)
		mp1 = append(mp1, v[0])
		param = append(param, mp1)
	}

	c.JSON(http.StatusOK, gin.H{
		"url":             url,
		"method":          select_host,
		"parameters":      param,
		"response_header": response_header, // 响应头
		"request_header": response_header, // 响应头
		"content":         content,
	})
}
