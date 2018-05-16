## wechat
Golang开发微信公众号

## 安装使用
- 进入conf文件夹，复制app.conf.example为app.conf，并完成配置
```
[app]
port = "端口号"
token = "公众号验证token"
```
- 编译可执行文件：
```
go build -o bin/wechat wechat.go
```
- 启动服务bin/wechat

## 功能介绍
- 你问我答  
粉丝给公众号一条文本消息，公众号立马回复相同文本消息给粉丝
- “图”尚往来  
粉丝给公众号一条图片消息，公众号立马回复相同图片消息给粉丝