package routers

import (
	"tiny-sec-kill/api/v1"
	"tiny-sec-kill/middlewares"
	"tiny-sec-kill/utils"

	"github.com/gin-gonic/gin"

	_ "tiny-sec-kill/docs"    // 不要忘了导入把你上一步生成的docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) //生产采用debug模式
	r := gin.Default()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())
	auth := r.Group("api/v1")
	auth.Use(middlewares.JwtToken())
	{
		auth.PUT("users/:id", v1.EditUser)
		auth.GET("getuser/:id", v1.GetUserInfo)
		auth.GET("getusers", v1.GetUsers)
		auth.POST("goods/add", v1.AddGood)
		auth.PUT("goods/:id", v1.EditGood)
		auth.PUT("preseckill/:id", v1.PreSecKill)
		auth.GET("seckill/:id", v1.SecKill)
	}
	router := r.Group("api/v1") //公共的，用来登录用的
	{
		router.POST("users/add", v1.AddUser)
		router.POST("login", v1.Login)
		router.GET("getgood/:id", v1.GetGood)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	r.Run(utils.HttpPort)
}
