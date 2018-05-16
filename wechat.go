package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markbest/wechat/conf"
	"github.com/markbest/wechat/middleware"
	"github.com/markbest/wechat/model"
)

func main() {
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}

	// set mode
	gin.SetMode(gin.ReleaseMode)

	// start server
	r := gin.Default()

	// add middleware
	r.Use(middleware.Logger())

	// router
	r.GET("/", model.HandleCheckSignature)
	r.POST("/", model.HandleRequest)
	r.Run(conf.Conf.App.Port)
}
