package v1

import (
	"net/http"
	"strconv"

	"tiny-sec-kill/models"
	"tiny-sec-kill/utils/errMsg"

	"github.com/gin-gonic/gin"
)
// @Summary 查看商品
// @Accept application/json
// @Produce json
// @Param id query string false "查看的商品ID"
// @Router /api/v1/getgood/{id} [get]
func GetGood(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := models.GetGood(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
	})
}
// @Summary 添加商品
// @Accept application/json
// @Produce json
// @Router /api/v1/goods/add [post]
func AddGood(c *gin.Context) {
	var data models.Good
	_ = c.ShouldBindJSON(&data)
	tokenHeader := c.Request.Header.Get("Authorization")
	models.CreateGood(&data, tokenHeader)
	c.JSON(http.StatusOK, gin.H{
		"status": errMsg.SUCCESS,
		"data":   data,
	})
}
// @Summary 编辑更改商品
// @Accept application/json
// @Produce json
// @Param id query string false "编辑更改的商品ID"
// @Router /api/v1/goods/{id} [put]
func EditGood(c *gin.Context) {
	var data models.Good
	_ = c.ShouldBindJSON(&data)
	id, _ := strconv.Atoi(c.Param("id"))
	tokenHeader := c.Request.Header.Get("Authorization")
	code := models.EditGood(id, &data, tokenHeader)
	c.JSON(http.StatusOK, gin.H{
		"剩余数量": data,
		"code": code,
	})
}
// @Summary 预先抢购商品
// @Accept application/json
// @Produce json
// @Param id query string false "预先抢购商品ID"
// @Router /api/v1/preseckill/{id} [put]
func PreSecKill(c *gin.Context) {
	//var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	tokenHeader := c.Request.Header.Get("Authorization")
	code, data := models.PreSecKill(id, tokenHeader)
	//fmt.Println(data)
	if code == errMsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"秒杀数量": data.StockNum,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "您没有权限",
		})
	}

}