package v1

import (
	"fmt"
	"net/http"
	"tiny-sec-kill/models"
	"tiny-sec-kill/utils/errMsg"
	"github.com/gin-gonic/gin"
)

var code int

func AddUser(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	fmt.Println(data)
	code := models.CheckUser(data.Username)
	if code == errMsg.SUCCESS {
		models.CreateUser(&data)
	} else if code == errMsg.ERROR {
		code = errMsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
	})
}
