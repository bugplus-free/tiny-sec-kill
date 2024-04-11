package v1

import (
	"net/http"
	"strconv"

	"tiny-sec-kill/models"

	"github.com/gin-gonic/gin"
)

func GetGood(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := models.GetGood(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
	})
}
