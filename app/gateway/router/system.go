package router

import (
	"github.com/gin-gonic/gin"
	v1 "grpc-admin/app/gateway/api/v1"
	"grpc-admin/app/gateway/middleware"
)

func RegisterSystemRouterV1(version *gin.RouterGroup) {
	api := v1.NewSystemApi()

	// 后台管理授权路由组
	adminGroup := version.Group("/admin/system")
	adminGroup.Use(middleware.Jwt())
	{
		// 创建路由组
		adminGroup.POST("/route/group", api.CreateRouteGroup)
		// 编辑路由组
		adminGroup.PUT("/route/group/:id", api.UpdateRouteGroup)
		// 路由组详情
		adminGroup.GET("/route/group/:id", api.GetRouteGroup)
		// 批量删除路由组
		adminGroup.DELETE("/route/group", api.DelRouteGroups)
		// 路由组列表
		adminGroup.GET("/route/group/list", api.ListRouteGroup)

		// 创建路由
		adminGroup.POST("/route", api.CreateRoute)
		// 编辑路由
		adminGroup.PUT("/route/:id", api.UpdateRoute)
		// 路由详情
		adminGroup.GET("/route/:id", api.GetRoute)
		// 批量删除路由
		adminGroup.DELETE("/route", api.DelRoutes)
		// 路由列表
		adminGroup.GET("/route/list", api.ListRoute)
		
		// 创建菜单
		adminGroup.POST("/menu", api.CreateMenu)
		// 编辑菜单
		adminGroup.PUT("/menu/:id", api.UpdateMenu)
		// 菜单详情
		adminGroup.GET("/menu/:id", api.GetMenu)
		// 批量删除菜单
		adminGroup.DELETE("/menu", api.DelMenus)
		// 菜单列表
		adminGroup.GET("/menu/list", api.ListMenu)
	}
}
