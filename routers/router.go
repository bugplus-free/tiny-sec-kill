package routers

import (
	"github.com/gin-gonic/gin"
	"tiny-sec-kill/utils"
	v1 "tiny-sec-kill/api/v1"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) //生产采用debug模式
	r := gin.Default()

	router := r.Group("api/v1")//公共的，用来登录用的
	
	r.Run(utils.HttpPort)
}
