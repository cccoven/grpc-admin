package bootstrap

import (
	"github.com/gin-gonic/gin"
	"grpc-admin/app/gateway/middleware"
	"grpc-admin/app/gateway/router"
)

func registerRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.Cors())
	{
		router.RegisterUserRouterV1(apiv1)
		router.RegisterThirdPartyRouterV1(apiv1)
		router.RegisterSystemRouterV1(apiv1)
	}
}
