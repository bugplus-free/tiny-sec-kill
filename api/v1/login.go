package v1

import (
	"net/http"
	"tiny-sec-kill/middlewares"
	"tiny-sec-kill/utils/errMsg"
	"tiny-sec-kill/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data models.User
	var token string
	var code int
	c.ShouldBindJSON(&data)
	code = models.CheckLogin(data.Username, data.Password)
	if code == errMsg.SUCCESS {
		token, code = middlewares.SetToken(data.Username, data.Password) //服务端生成一个token，一般来说就使用了base64
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"token":  token,
	})
}
