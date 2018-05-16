package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/markbest/wechat/utils"
	"io/ioutil"
	"os"
	"time"
)

type logDebug struct {
	CreatedAt string
	Request   request
}

type request struct {
	Url         string
	Method      string
	ContentType []string
	Body        string
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _ := os.OpenFile("wechat.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer file.Close()
		logData := logDebug{}

		// created at
		t := time.Now()
		logData.CreatedAt = t.Format("2006-01-02 15:04:05")

		// request url
		logData.Request.Url = c.Request.URL.RequestURI()

		// request method
		logData.Request.Method = c.Request.Method

		// request content type
		logData.Request.ContentType = c.Request.Header["Content-Type"]

		// request body
		body, _ := ioutil.ReadAll(c.Request.Body)
		logData.Request.Body = string(body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// write log
		data, _ := json.Marshal(logData)
		data = utils.ReplaceEscapeStr(data)
		file.WriteString(string(data) + "\n")

		c.Next()
	}
}
