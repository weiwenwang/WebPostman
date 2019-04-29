package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Urlinfo(c *gin.Context) {
	//method := c.Query("method")
	//url := c.Query("url")
	//resp, err := http.Get(url)
	//if err != nil {
	//	// handle error
	//}
	//
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//
	//header := make(map[string]string)
	//vs := ""
	//for k, v := range resp.Header {
	//	for _, vv := range v {
	//		vs = vv + ","
	//	}
	//	header[k] = vs
	//}
	//header_str, err := json.Marshal(header)
	//
	//if err != nil {
	//	// handle error
	//}
	url := "http://iphonealpha.wochacha.com/exposure/list?tabid=9999&page=1&class=&city_id=1&udid=e2dcaffb7e6e58a4d205bda4dc89dcbcd9908d7e&openid=e2dcaffb7e6e58a4d205bda4dc89dcbcd9908d7e&v=9.5.10&urid=1376108936&devicetoken=959751087fcfd5fafd4e8cbeeccc9bcf&token=074261ae07f02d21df1f8c615e867da0&source=iphone&connectnet=wifi&lng=121.621579&lat=31.210546&os=iphone&his=1555651773.1556007398&device_model=iPhone&device_brand=iPhone7,1&wcctoken=3d281f9c54b2aebfb4d59068b84b5bb2&osversion=8.0"
	method := 2
	c.JSON(http.StatusOK, gin.H{
		"url":    url,
		"method": method,
	})
}
