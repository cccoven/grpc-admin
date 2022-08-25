package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"grpc-admin/app/gateway/conf"
	"grpc-admin/app/gateway/helper"
	"grpc-admin/app/gateway/pkg/logger"
	"grpc-admin/app/gateway/request"
	"grpc-admin/app/gateway/response"
	"grpc-admin/app/gateway/rpc"
	"grpc-admin/app/system/system"
	"grpc-admin/common/util"
	"strconv"
	"time"
)

type SystemApi struct {
	logger    *zap.SugaredLogger
	systemRpc system.SystemClient
}

// CreateRouteGroup 创建路由组
func (s *SystemApi) CreateRouteGroup(c *gin.Context) {
	var req request.ModifyRouteGroup
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	routeGroupSchema, err := s.systemRpc.ModifyRouteGroup(context.Background(), &system.RouteGroupSchema{
		Name: req.Name,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.RouteGroup{
		ID:        routeGroupSchema.ID,
		CreatedAt: time.UnixMilli(routeGroupSchema.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.UnixMilli(routeGroupSchema.UpdatedAt).Format("2006-01-02 15:04:05"),
		Name:      routeGroupSchema.Name,
	}

	response.Success(c, resp)
}

// UpdateRouteGroup 编辑路由组
func (s *SystemApi) UpdateRouteGroup(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	var req request.ModifyRouteGroup
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	routeGroupSchema, err := s.systemRpc.ModifyRouteGroup(context.Background(), &system.RouteGroupSchema{
		ID:   uint32(id),
		Name: req.Name,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.RouteGroup{
		ID:        routeGroupSchema.ID,
		CreatedAt: time.UnixMilli(routeGroupSchema.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.UnixMilli(routeGroupSchema.UpdatedAt).Format("2006-01-02 15:04:05"),
		Name:      routeGroupSchema.Name,
	}

	response.Success(c, resp)
}

// GetRouteGroup 路由组详情
func (s *SystemApi) GetRouteGroup(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	group, err := s.systemRpc.GetRouteGroup(context.Background(), &system.RouteGroupSchema{ID: uint32(id)})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	response.Success(c, response.RouteGroup{
		ID:        group.ID,
		CreatedAt: time.UnixMilli(group.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.UnixMilli(group.UpdatedAt).Format("2006-01-02 15:04:05"),
		Name:      group.Name,
	})
}

// DelRouteGroups 批量删除路由组
func (s *SystemApi) DelRouteGroups(c *gin.Context) {
	var req request.MultipleIDs
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	_, err := s.systemRpc.DeleteRouteGroup(context.Background(), &system.MultipleID{IDs: req.IDs})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	response.Success(c, nil)
}

// ListRouteGroup 路由组列表分页
func (s *SystemApi) ListRouteGroup(c *gin.Context) {
	var req request.ListRouteGroup
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	listReply, err := s.systemRpc.ListRouteGroup(context.Background(), &system.ListRouteGroupRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		RouteGroup: &system.RouteGroupSchema{
			Name: req.RouteGroup.Name,
		},
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	var items []response.RouteGroup
	for _, schema := range listReply.List {
		items = append(items, response.RouteGroup{
			ID:        schema.ID,
			CreatedAt: time.UnixMilli(schema.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.UnixMilli(schema.UpdatedAt).Format("2006-01-02 15:04:05"),
			Name:      schema.Name,
		})
	}
	response.Success(c, response.Pagination{
		Page:  req.Page,
		List:  items,
		Total: listReply.Total,
	})
}

// CreateRoute 创建路由
func (s *SystemApi) CreateRoute(c *gin.Context) {
	var req request.ModifyRoute
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	schema, err := s.systemRpc.ModifyRoute(context.Background(), &system.RouteSchema{
		Path:    req.Path,
		Method:  req.Method,
		Desc:    req.Desc,
		GroupID: req.GroupID,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	response.Success(c, response.Route{
		ID:        schema.ID,
		CreatedAt: time.UnixMilli(schema.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.UnixMilli(schema.UpdatedAt).Format("2006-01-02 15:04:05"),
		Path:      schema.Path,
		Method:    schema.Method,
		Desc:      schema.Desc,
		GroupID:   schema.GroupID,
		Group: response.RouteGroup{
			ID:        schema.Group.ID,
			CreatedAt: time.UnixMilli(schema.Group.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.UnixMilli(schema.Group.UpdatedAt).Format("2006-01-02 15:04:05"),
			Name:      schema.Group.Name,
		},
	})
}

// UpdateRoute 编辑路由
func (s *SystemApi) UpdateRoute(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	var req request.ModifyRoute
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	schema, err := s.systemRpc.ModifyRoute(context.Background(), &system.RouteSchema{
		ID:      uint32(id),
		Path:    req.Path,
		Method:  req.Method,
		Desc:    req.Desc,
		GroupID: req.GroupID,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	response.Success(c, response.Route{
		ID:        schema.ID,
		CreatedAt: time.UnixMilli(schema.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.UnixMilli(schema.UpdatedAt).Format("2006-01-02 15:04:05"),
		Path:      schema.Path,
		Method:    schema.Method,
		Desc:      schema.Desc,
		GroupID:   schema.GroupID,
		Group: response.RouteGroup{
			ID:        schema.Group.ID,
			CreatedAt: time.UnixMilli(schema.Group.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.UnixMilli(schema.Group.UpdatedAt).Format("2006-01-02 15:04:05"),
			Name:      schema.Group.Name,
		},
	})
}

// GetRoute 路由详情
func (s *SystemApi) GetRoute(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	schema, err := s.systemRpc.GetRoute(context.Background(), &system.RouteSchema{ID: uint32(id)})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	response.Success(c, response.Route{
		ID:        schema.ID,
		CreatedAt: time.UnixMilli(schema.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.UnixMilli(schema.UpdatedAt).Format("2006-01-02 15:04:05"),
		Path:      schema.Path,
		Method:    schema.Method,
		Desc:      schema.Desc,
		GroupID:   schema.GroupID,
		Group: response.RouteGroup{
			ID:        schema.Group.ID,
			CreatedAt: time.UnixMilli(schema.Group.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.UnixMilli(schema.Group.UpdatedAt).Format("2006-01-02 15:04:05"),
			Name:      schema.Group.Name,
		},
	})
}

// DelRoutes 批量删除路由
func (s *SystemApi) DelRoutes(c *gin.Context) {
	var req request.MultipleIDs
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	_, err := s.systemRpc.DeleteRoute(context.Background(), &system.MultipleID{IDs: req.IDs})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	response.Success(c, nil)
}

// ListRoute 路由列表分页
func (s *SystemApi) ListRoute(c *gin.Context) {
	var req request.ListRoute
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	listReply, err := s.systemRpc.ListRoute(context.Background(), &system.ListRouteRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Route: &system.RouteSchema{
			Path:    req.Route.Path,
			Method:  req.Route.Method,
			GroupID: req.Route.GroupID,
		},
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}
	var items []response.Route
	for _, schema := range listReply.List {
		items = append(items, response.Route{
			ID:        schema.ID,
			CreatedAt: time.UnixMilli(schema.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.UnixMilli(schema.UpdatedAt).Format("2006-01-02 15:04:05"),
			Path:      schema.Path,
			Method:    schema.Method,
			Desc:      schema.Desc,
			GroupID:   schema.GroupID,
			Group: response.RouteGroup{
				ID:        schema.Group.ID,
				CreatedAt: time.UnixMilli(schema.Group.CreatedAt).Format("2006-01-02 15:04:05"),
				UpdatedAt: time.UnixMilli(schema.Group.UpdatedAt).Format("2006-01-02 15:04:05"),
				Name:      schema.Group.Name,
			},
		})
	}
	response.Success(c, response.Pagination{
		Page:  req.Page,
		List:  items,
		Total: listReply.Total,
	})
}

// CreateMenu 创建菜单
func (s *SystemApi) CreateMenu(c *gin.Context) {
	var req request.ModifyMenu
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	schema, err := s.systemRpc.ModifyMenu(context.Background(), &system.MenuSchema{
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Sort:      req.Sort,
		ParentID:  req.ParentID,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := &response.Menu{Children: []response.Menu{}}
	_ = util.CopyWithTimeConverter(resp, schema)

	response.Success(c, resp)
}

// UpdateMenu 编辑菜单
func (s *SystemApi) UpdateMenu(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	var req request.ModifyMenu
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	schema, err := s.systemRpc.ModifyMenu(context.Background(), &system.MenuSchema{
		ID:        uint32(id),
		Path:      req.Path,
		Name:      req.Name,
		Component: req.Component,
		Sort:      req.Sort,
		ParentID:  req.ParentID,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := &response.Menu{Children: []response.Menu{}}
	_ = util.CopyWithTimeConverter(resp, schema)

	response.Success(c, resp)
}

// DelMenus 批量删除菜单
func (s *SystemApi) DelMenus(c *gin.Context) {
	var req request.MultipleIDs
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	_, err := s.systemRpc.DeleteMenu(context.Background(), &system.MultipleID{IDs: req.IDs})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	response.Success(c, nil)
}

// ListMenu 菜单列表
func (s *SystemApi) ListMenu(c *gin.Context) {
	var req request.ListMenu
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	listReply, err := s.systemRpc.ListMenu(context.Background(), &system.ListMenuRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Menu: &system.MenuSchema{
			Path: req.Path,
			Name: req.Name,
		},
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.Pagination{
		Page:  req.Page,
		List:  []response.Menu{},
		Total: listReply.Total,
	}
	_ = util.CopyWithTimeConverter(&resp.List, listReply.List)

	response.Success(c, resp)
}

// GetMenu 菜单详情
func (s *SystemApi) GetMenu(c *gin.Context) {
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	schema, err := s.systemRpc.GetMenu(context.Background(), &system.MenuSchema{ID: uint32(id)})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := &response.Menu{Children: []response.Menu{}}
	_ = util.CopyWithTimeConverter(resp, schema)

	response.Success(c, resp)
}

func NewSystemApi() *SystemApi {
	return &SystemApi{
		logger:    logger.NewZapLogger(),
		systemRpc: system.NewSystemClient(rpc.Discovery(conf.AppConf.SystemRpc.Name, conf.AppConf.SystemRpc.LoadBalanceMode)),
	}
}
