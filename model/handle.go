package model

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/markbest/wechat/conf"
	"github.com/markbest/wechat/utils"
	"net/http"
	"sort"
	"time"
)

// Validator url is wechat send
func HandleCheckSignature(c *gin.Context) {
	signature, _ := c.GetQuery("signature")
	timeStamp, _ := c.GetQuery("timestamp")
	nonce, _ := c.GetQuery("nonce")
	echoStr, _ := c.GetQuery("echostr")

	tmpStrings := []string{conf.Conf.App.Token, timeStamp, nonce}
	sort.Strings(tmpStrings)
	tmpStr := tmpStrings[0] + tmpStrings[1] + tmpStrings[2]
	tmp := utils.StrToSha1(tmpStr)
	if tmp == signature {
		c.String(200, echoStr)
	} else {
		c.String(401, "Unauthorized")
	}
}

// Handle request
func HandleRequest(c *gin.Context) {
	var base BaseMsg
	contentType := c.Request.Header.Get("Content-Type")
	switch contentType {
	case "text/xml":
		body, _ := c.GetRawData()
		err := xml.Unmarshal(body, &base)
		if err != nil {
			c.XML(http.StatusBadGateway, gin.H{
				"status":  false,
				"message": "Unmarshal body err",
			})
			return
		}
		result, err := handleReply(base, body)
		if err != nil {
			c.XML(http.StatusBadGateway, gin.H{
				"status":  false,
				"message": "Unmarshal body err",
			})
			return
		}
		result = utils.ReplaceEscapeStr(result)
		c.String(http.StatusOK, string(result))
	}
}

// Handle reply
func handleReply(base BaseMsg, body []byte) ([]byte, error) {
	// text msg
	if base.MsgType == "text" {
		var xmlContent TextMsg
		err := xml.Unmarshal(body, &xmlContent)
		if err != nil {
			return nil, err
		}
		result := TextMsg{}
		result.ToUserName = cDataString{Value: xmlContent.FromUserName.Value}
		result.FromUserName = cDataString{Value: xmlContent.ToUserName.Value}
		result.CreateTime = time.Now().Unix()
		result.MsgType = cDataString{Value: xmlContent.MsgType.Value}
		result.Content = cDataString{Value: xmlContent.Content.Value}
		reply, _ := xml.Marshal(result)
		return reply, nil
	}

	// image msg
	if base.MsgType == "image" {
		var xmlContent ReceiveImageMsg
		err := xml.Unmarshal(body, &xmlContent)
		if err != nil {
			return nil, err
		}
		result := ReplyImageMsg{}
		result.ToUserName = cDataString{Value: xmlContent.FromUserName.Value}
		result.FromUserName = cDataString{Value: xmlContent.ToUserName.Value}
		result.CreateTime = time.Now().Unix()
		result.MsgType = cDataString{Value: xmlContent.MsgType.Value}
		result.Image = image{MediaId: cDataString{Value: xmlContent.MediaId.Value}}
		reply, _ := xml.Marshal(result)
		return reply, nil
	}
	return nil, nil
}
