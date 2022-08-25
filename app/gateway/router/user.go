package router

import (
	"github.com/gin-gonic/gin"
	v1 "grpc-admin/app/gateway/api/v1"
	"grpc-admin/app/gateway/middleware"
)

func RegisterUserRouterV1(version *gin.RouterGroup) {
	api := v1.NewUserApi()

	// C 端公共路由组
	// publicUserGroup := version.Group("/user")
	// {
	// publicUserGroup.POST("/login", api.SignIn)
	// }

	// C 端授权路由组
	userGroup := version.Group("/user")
	userGroup.Use(middleware.Jwt())
	{

	}

	// 后台管理公共路由组
	publicAdminGroup := version.Group("/admin/user")
	{
		publicAdminGroup.POST("/login", api.AdminLogin)
	}

	// 后台管理授权路由组
	adminGroup := version.Group("/admin/user")
	adminGroup.Use(middleware.Jwt() /*middleware.Auth()*/)
	{
		/* 用户相关 */
		// 创建用户
		adminGroup.POST("/", api.CreateUser)
		// 用户详情
		adminGroup.GET("/:id", api.UserDetail)
		// 编辑用户
		adminGroup.PUT("/:id", api.UpdateUser)
		// 删除用户
		adminGroup.DELETE("/:id", api.DelUser)
		// 用户列表
		adminGroup.GET("/list", api.UserList)

		/* RBAC 相关 */
		// 创建角色
		adminGroup.POST("/role", api.CreateRole)
		// 编辑角色
		adminGroup.PATCH("/role/:id", api.UpdateRole)
		// 给角色分配权限
		adminGroup.POST("/role/:id/authorize", api.AuthorizeRole)
		// 给角色分配菜单
		adminGroup.POST("/role/:id/menus", api.SetRoleMenus)
		// 给用户分配角色
		adminGroup.PUT("/:id/role", api.AssignRolesToUser)
		// 删除角色
		adminGroup.DELETE("/role/:id", api.DelRole)
		// 角色列表
		adminGroup.GET("/role/list", api.ListRole)
		// TODO 获取角色权限

		// TODO 获取角色菜单
	}
}

func RegisterUserRouterV2(v *gin.RouterGroup) {}
