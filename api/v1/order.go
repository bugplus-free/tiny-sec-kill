package v1

import (
	"net/http"
	"strconv"
	"tiny-sec-kill/models"
	"github.com/gin-gonic/gin"
)
// @Summary 抢购商品
// @Accept application/json
// @Produce json
// @Param id query string false "抢购商品ID"
// @Router /api/v1/seckill/{id} [get]
func SecKill(c *gin.Context) {
	var data models.Order
	id, _ := strconv.Atoi(c.Param("id"))
	data.Ip = c.ClientIP()
	data.Gid = id
	tokenHeader := c.Request.Header.Get("Authorization")
	dd := models.TransToken(tokenHeader)

	models.SecKill(id, &data, dd.Username)
	c.JSON(http.StatusOK, gin.H{
		"status": data.Msg,
	})

}
