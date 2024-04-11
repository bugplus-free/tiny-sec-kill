package routers

import (
	middleware "tiny-sec-kill/middlewares"
	"tiny-sec-kill/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) //生产采用debug模式
	r := gin.Default()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	auth:=r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

	}
	router := r.Group("api/v1")//公共的，用来登录用的
	{

	}
	r.Run(utils.HttpPort)
}
