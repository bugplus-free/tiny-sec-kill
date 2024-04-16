package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"tiny-sec-kill/models"
	"tiny-sec-kill/utils/errMsg"

	"github.com/gin-gonic/gin"
)

var code int
// @Summary 添加用户
// @Accept application/json
// @Produce json
// @Router /api/v1/users/add [post]
// 添加用户
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

// @Summary 查看用户信息
// @Accept application/json
// @Produce json
// @Param id query string false "查询的用户ID"
// @Router /api/v1/getuser/{id} [get]
// 查看用户信息
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	var maps = make(map[string]interface{})
	data, code := models.GetUser(id)
	maps["username"] = data.Username
	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
			"data":   maps,
			"total":  1,
		},
	)

}

// @Summary 更新用户
// @Accept application/json
// @Produce json
// @Param id query string false "用户ID"
// @Router /api/v1/users/{id} [put]
func EditUser(c *gin.Context) {
	var data models.User
	tokenHeader := c.Request.Header.Get("Authorization")
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.CheckUser(data.Username)
	
	if code == errMsg.SUCCESS {
		models.EditUser(id, &data, tokenHeader)
	} else if code == errMsg.ERROR {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
	})
}
// @Summary 获取用户组
// @Accept application/json
// @Produce json
// @Param pageSize query string false "页面大小"
// @Param pageNum query string false "页面数"
// @Router /api/v1/getusers [get]
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageNum == 0 {
		pageNum = -1
	}
	tokenHeader := c.Request.Header.Get("Authorization")
	data, total := models.GetUsers(pageSize, pageNum, tokenHeader)
	if total == -1 {
		c.JSON(http.StatusOK, gin.H{
			"status": errMsg.ERROR_USER_NO_PERMISSION,
		})
	} else {
		code := errMsg.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   data,
			"total":  total,
		})
	}
}
//这两个功能并不希望对外开放，可能会出现token泛滥，不断地有超级用户生成
//并且没有进行用户删除以后token的清理工作，token的消失只能依靠时间来消失
//这里存在着一些问题
func ChangeUserPassword(c *gin.Context) {
	var data models.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := models.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status": code,
		},
	)
}
//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tokenHeader := c.Request.Header.Get("Authorization")
	code := models.DeleteUser(id, tokenHeader)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
	})
}
