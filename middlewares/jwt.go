package middlewares

import (
	"net/http"
	"strings"
	"time"
	"tiny-sec-kill/utils"
	"tiny-sec-kill/utils/errMsg"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	Username string `json:"username" `
	Password string `json:"password"`

	jwt.StandardClaims
}

// 生成token
func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "tiny-sec-kill",
		},
	}
	requestClaim := jwt.NewWithClaims(jwt.SigningMethodES256, SetClaims)
	token, err := requestClaim.SignedString(JwtKey)
	if err != nil {
		return "", errMsg.ERROR
	}
	return token, errMsg.SUCCESS
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errMsg.SUCCESS
	} else {
		return nil, errMsg.ERROR
	}
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code = errMsg.SUCCESS
		if tokenHeader == "" {
			code = errMsg.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errMsg.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
			})
			c.Abort()
			return
		}
		key, err := CheckToken(checkToken[1])
		if err == errMsg.ERROR {
			code = errMsg.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
