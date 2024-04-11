package routers

import (
	"tiny-sec-kill/api/v1"
	"tiny-sec-kill/middlewares"
	"tiny-sec-kill/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) //生产采用debug模式
	r := gin.Default()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())
	auth := r.Group("api/v1")
	auth.Use(middlewares.JwtToken())
	{

	}
	router := r.Group("api/v1") //公共的，用来登录用的
	{
		router.POST("users/add", v1.AddUser)
		router.POST("login", v1.Login)
		router.GET("getgood/:id", v1.GetGood)
	}
	r.Run(utils.HttpPort)
}
