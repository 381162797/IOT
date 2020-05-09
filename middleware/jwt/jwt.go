package jwt

import (
	"mserver/def"
	"mserver/tool"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		var code int
		code = def.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = def.INVALID_PARAMS
		} else {
			claims, err := tool.ParseToken(token)
			if err != nil {
				code = def.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = def.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != def.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  def.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
