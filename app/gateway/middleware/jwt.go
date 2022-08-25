package middleware

import (
	"github.com/gin-gonic/gin"
	"grpc-admin/app/gateway/helper"
	"grpc-admin/app/gateway/response"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.Error(c, http.StatusUnauthorized, "授权失败")
			c.Abort()
			return
		}

		claims, err := helper.ParseJwt(token)
		if err != nil {
			if err == helper.TokenExpired {
				response.Error(c, http.StatusUnauthorized, "授权已过期")
				c.Abort()
				return
			}
			response.Error(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		
		// TODO 刷新 token
		
		c.Set("claims", claims)
		c.Next()
	}
}
