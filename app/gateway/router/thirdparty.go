package router

import (
	"github.com/gin-gonic/gin"
	v1 "grpc-admin/app/gateway/api/v1"
)

func RegisterThirdPartyRouterV1(v *gin.RouterGroup) {
	g := v.Group("/thirdparty")
	api := v1.NewThirdPartyApi()
	{
		// 发送短信
		g.POST("/sms/send", api.SendSMS)
	}
}
