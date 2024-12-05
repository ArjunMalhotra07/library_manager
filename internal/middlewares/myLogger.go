package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.FullPath()
		// headers := c.Request.Header
		queryParams := c.Request.URL.Query()
		clientIP := c.ClientIP()
		bodyBytes, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		fmt.Println("Method", method)
		fmt.Println("path", path)
		// fmt.Println("headers", headers)
		fmt.Println("queryParams", queryParams)
		fmt.Println("clientIP", clientIP)
		fmt.Println("Body", string(bodyBytes))
		c.Next()
	}
}
