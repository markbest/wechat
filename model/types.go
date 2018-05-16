package model

import (
	"encoding/xml"
)

type BaseMsg struct {
	XMLName xml.Name `xml:"xml"`
	MsgType string   `xml:"MsgType"`
}

type TextMsg struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cDataString `xml:"ToUserName"`
	FromUserName cDataString `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      cDataString `xml:"MsgType"`
	Content      cDataString `xml:"Content"`
}

type ReceiveImageMsg struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cDataString `xml:"ToUserName"`
	FromUserName cDataString `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      cDataString `xml:"MsgType"`
	PicUrl       cDataString `xml:"PicUrl"`
	MediaId      cDataString `xml:"MediaId"`
}

type ReplyImageMsg struct {
	XMLName      xml.Name    `xml:"xml"`
	ToUserName   cDataString `xml:"ToUserName"`
	FromUserName cDataString `xml:"FromUserName"`
	CreateTime   int64       `xml:"CreateTime"`
	MsgType      cDataString `xml:"MsgType"`
	Image        image       `xml:"Image"`
}

type image struct {
	MediaId cDataString `xml:"MediaId"`
}

type cDataString struct {
	Value string `xml:",cdata"`
}
