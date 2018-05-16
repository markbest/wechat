package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbest/wechat/utils"
	"io/ioutil"
	"os"
	"time"
)

type logDebug struct {
	CreatedAt   string   `json:"create_at"`
	Url         string   `json:"url"`
	Method      string   `json:"method"`
	ContentType []string `json:"content_type"`
	Body        string   `json:"body"`
	ClientIp    string   `json:"client_ip"`
	Status      int      `json:"status"`
	Latency     string   `json:"latency"`
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
		logData.Url = c.Request.URL.RequestURI()

		// request method
		logData.Method = c.Request.Method

		// request content type
		logData.ContentType = c.Request.Header["Content-Type"]

		// request body
		body, _ := ioutil.ReadAll(c.Request.Body)
		logData.Body = string(body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// request client ip
		logData.ClientIp = c.ClientIP()

		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		// request latency
		logData.Latency = fmt.Sprintf("%v", latency)

		// request status
		logData.Status = c.Writer.Status()

		// write log
		data, _ := json.Marshal(logData)
		data = utils.ReplaceEscapeStr(data)
		file.WriteString(string(data) + "\n")
	}
}
