## wechat
Golang开发微信公众号

## usage
- 进入conf文件夹，复制conf.toml.example为conf.toml，并完成配置
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

## feature
- 你问我答  
粉丝给公众号一条文本消息，公众号立马回复相同文本消息给粉丝
- “图”尚往来  
粉丝给公众号一条图片消息，公众号立马回复相同图片消息给粉丝