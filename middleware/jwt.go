package middleware

import (
	"github.com/gin-gonic/gin"
	"test_mysql/pkg/e"
	"test_mysql/pkg/util"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = 200
		//从请求头中获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 过期了
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.Success {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				//"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
