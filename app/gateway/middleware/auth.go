package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"grpc-admin/app/gateway/conf"
	"grpc-admin/app/gateway/helper"
	"grpc-admin/app/gateway/response"
	"grpc-admin/app/gateway/rpc"
	"grpc-admin/app/user/user"
	"net/http"
)

func Auth() gin.HandlerFunc {
	userRpc := user.NewUserClient(rpc.Discovery(conf.AppConf.UserRpc.Name, conf.AppConf.UserRpc.LoadBalanceMode))

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := helper.ParseJwt(token)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, err.Error())
			return
		}

		// 调用用户服务查询用户所属角色是否有该接口权限
		authorityReply, err := userRpc.CheckRolesAuthority(context.Background(), &user.CheckRolesAuthorityRequest{
			RoleIDs: claims.RoleIDs,
			Authority: &user.Authority{
				Path:   c.Request.URL.Path,
				Method: c.Request.Method,
			},
		})
		if err != nil {
			response.RpcError(c, err)
			c.Abort()
			return
		}

		if !authorityReply.Ok {
			response.Error(c, http.StatusForbidden, "权限不足")
			c.Abort()
			return
		}

		c.Next()
	}
}
