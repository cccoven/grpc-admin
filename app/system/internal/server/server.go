package server

import (
	"context"
	"grpc-admin/app/system/internal/dto"
	"grpc-admin/app/system/internal/model"
	"grpc-admin/app/system/internal/model/migrate"
	"grpc-admin/app/system/internal/pkg/db"
	"grpc-admin/app/system/internal/pkg/errorx"
	"grpc-admin/app/system/internal/repo"
	"grpc-admin/app/system/system"
	"grpc-admin/common/util"
)

type server struct {
	system.UnimplementedSystemServer
	repo repo.ISystemRepo
}

// ModifyMenu 创建/更新菜单
func (s *server) ModifyMenu(ctx context.Context, in *system.MenuSchema) (*system.MenuSchema, error) {
	var err error
	menu := new(model.Menu)

	if found, _ := s.repo.FindMenu(model.Menu{Path: in.Path}, false); found != nil && found.ID != uint(in.ID) {
		return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在相同菜单"))
	}
	if in.ID == 0 {
		// if found, _ := s.repo.FindMenu(model.Menu{Path: in.Path}, false); found != nil {
		// 	return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在相同菜单"))
		// }
		menu.Path = in.Path
		menu.Name = in.Name
		menu.Component = in.Component
		menu.Sort = int(in.Sort)
		menu.ParentID = uint(in.ParentID)
		if err = s.repo.CreateMenu(menu); err != nil {
			return nil, errorx.StatusError(err)
		}
	} else {
		// 获取完整数据
		if menu, err = s.repo.FindMenu(model.Menu{ID: uint(in.ID)}, true); err != nil {
			return nil, errorx.StatusError(err)
		}
		menu.Path = in.Path
		menu.Name = in.Name
		menu.Component = in.Component
		menu.Sort = int(in.Sort)
		menu.ParentID = uint(in.ParentID)
		if err = s.repo.UpdateMenu(menu); err != nil {
			return nil, errorx.StatusError(err)
		}
	}

	reply := new(system.MenuSchema)
	_ = util.CopyWithTimeConverter(reply, menu)

	return reply, nil
}

// GetMenu 获取菜单详情
func (s *server) GetMenu(ctx context.Context, in *system.MenuSchema) (*system.MenuSchema, error) {
	menu, err := s.repo.FindMenu(model.Menu{ID: uint(in.ID)}, true)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	reply := new(system.MenuSchema)
	_ = util.CopyWithTimeConverter(reply, menu)

	return reply, nil
}

// DeleteMenu 删除菜单
func (s *server) DeleteMenu(ctx context.Context, in *system.MultipleID) (*system.SystemProtoEmpty, error) {
	if err := s.repo.DelMenu(in.IDs); err != nil {
		return nil, errorx.StatusError(err)
	}
	return &system.SystemProtoEmpty{}, nil
}

// ListMenu 菜单列表
func (s *server) ListMenu(ctx context.Context, in *system.ListMenuRequest) (*system.ListMenuReply, error) {
	page := dto.Pagination{Page: in.Page, PageSize: in.PageSize}
	cond := model.Menu{Path: in.Menu.Path, Name: in.Menu.Name}
	list, total, err := s.repo.FindMenusPagination(page, cond)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	reply := &system.ListMenuReply{List: []*system.MenuSchema{}, Total: total}
	_ = util.CopyWithTimeConverter(&reply.List, list)
	return reply, nil
}

// ModifyRouteGroup 创建/更新路由组
func (s *server) ModifyRouteGroup(ctx context.Context, in *system.RouteGroupSchema) (*system.RouteGroupSchema, error) {
	var err error
	group := &model.RouteGroup{}

	if in.ID == 0 {
		// 已存在同名路由组
		if found, _ := s.repo.FindRouteGroup(model.RouteGroup{Name: in.Name}); found != nil {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在同名路由组"))
		}
		group.Name = in.Name
		if err = s.repo.CreateRouteGroup(group); err != nil {
			return nil, errorx.StatusError(err)
		}
	} else {
		// 已存在同名路由组
		if found, _ := s.repo.FindRouteGroup(model.RouteGroup{Name: in.Name}); found != nil && found.ID != uint(in.ID) {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在同名路由组"))
		}
		// 获取完整数据
		if group, err = s.repo.FindRouteGroup(model.RouteGroup{ID: uint(in.ID)}); err != nil {
			return nil, errorx.StatusError(err)
		}
		group.Name = in.Name
		if err := s.repo.UpdateRouteGroup(group); err != nil {
			return nil, err
		}
	}

	routeGroupSchema := &system.RouteGroupSchema{
		ID:        uint32(group.ID),
		CreatedAt: group.CreatedAt.UnixMilli(),
		UpdatedAt: group.UpdatedAt.UnixMilli(),
		Name:      group.Name,
	}

	return routeGroupSchema, nil
}

// GetRouteGroup 获取路由组详情
func (s *server) GetRouteGroup(ctx context.Context, in *system.RouteGroupSchema) (*system.RouteGroupSchema, error) {
	group, err := s.repo.FindRouteGroup(model.RouteGroup{ID: uint(in.ID)})
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	return &system.RouteGroupSchema{
		ID:        uint32(group.ID),
		CreatedAt: group.CreatedAt.UnixMilli(),
		UpdatedAt: group.UpdatedAt.UnixMilli(),
		Name:      group.Name,
	}, nil
}

// DeleteRouteGroup 批量删除路由组
func (s *server) DeleteRouteGroup(ctx context.Context, in *system.MultipleID) (*system.SystemProtoEmpty, error) {
	if err := s.repo.DelRouteGroupByIDs(in.IDs); err != nil {
		return nil, errorx.StatusError(err)
	}
	return &system.SystemProtoEmpty{}, nil
}

// ListRouteGroup 路由组列表
func (s *server) ListRouteGroup(ctx context.Context, in *system.ListRouteGroupRequest) (*system.ListRouteGroupReply, error) {
	page := dto.Pagination{Page: in.Page, PageSize: in.PageSize}
	cond := model.RouteGroup{Name: in.RouteGroup.Name}
	list, total, err := s.repo.FindRouteGroupsPagination(page, cond)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	reply := &system.ListRouteGroupReply{Total: total}
	for _, group := range list {
		reply.List = append(reply.List, &system.RouteGroupSchema{
			ID:        uint32(group.ID),
			CreatedAt: group.CreatedAt.UnixMilli(),
			UpdatedAt: group.UpdatedAt.UnixMilli(),
			Name:      group.Name,
		})
	}
	return reply, nil
}

// ModifyRoute 创建/更新路由
func (s *server) ModifyRoute(ctx context.Context, in *system.RouteSchema) (*system.RouteSchema, error) {
	var err error
	route := &model.Route{}

	if in.ID == 0 {
		// 已存在相同路由
		if found, _ := s.repo.FindRoute(model.Route{Path: in.Path, Method: in.Method}, false); found != nil {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在相同路由"))
		}
		route.Path = in.Path
		route.Method = in.Method
		route.Desc = in.Desc
		route.GroupID = uint(in.GroupID)
		if err = s.repo.CreateRoute(route); err != nil {
			return nil, errorx.StatusError(err)
		}
	} else {
		// 已存在相同路由
		if found, _ := s.repo.FindRoute(model.Route{Path: in.Path, Method: in.Method}, false); found != nil && found.ID != uint(in.ID) {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在相同路由"))
		}
		// 获取完整数据
		if route, err = s.repo.FindRoute(model.Route{ID: uint(in.ID)}, true); err != nil {
			return nil, errorx.StatusError(err)
		}
		route.Path = in.Path
		route.Method = in.Method
		route.Desc = in.Desc
		route.GroupID = uint(in.GroupID)
		if err := s.repo.UpdateRoute(route); err != nil {
			return nil, errorx.StatusError(err)
		}
	}

	routeSchema := &system.RouteSchema{
		ID:        uint32(route.ID),
		CreatedAt: route.CreatedAt.UnixMilli(),
		UpdatedAt: route.UpdatedAt.UnixMilli(),
		Path:      route.Path,
		Method:    route.Method,
		Desc:      route.Desc,
		GroupID:   uint32(route.GroupID),
		Group: &system.RouteGroupSchema{
			ID:        uint32(route.Group.ID),
			CreatedAt: route.Group.CreatedAt.UnixMilli(),
			UpdatedAt: route.Group.UpdatedAt.UnixMilli(),
			Name:      route.Group.Name,
		},
	}

	return routeSchema, nil
}

// GetRoute 获取路由
func (s *server) GetRoute(ctx context.Context, in *system.RouteSchema) (*system.RouteSchema, error) {
	route, err := s.repo.FindRoute(model.Route{ID: uint(in.ID)}, true)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	return &system.RouteSchema{
		ID:        uint32(route.ID),
		CreatedAt: route.CreatedAt.UnixMilli(),
		UpdatedAt: route.UpdatedAt.UnixMilli(),
		Path:      route.Path,
		Method:    route.Method,
		Desc:      route.Desc,
		GroupID:   uint32(route.GroupID),
		Group: &system.RouteGroupSchema{
			ID:        uint32(route.Group.ID),
			CreatedAt: route.Group.CreatedAt.UnixMilli(),
			UpdatedAt: route.Group.UpdatedAt.UnixMilli(),
			Name:      route.Group.Name,
		},
	}, nil
}

// DeleteRoute 批量删除路由
func (s *server) DeleteRoute(ctx context.Context, in *system.MultipleID) (*system.SystemProtoEmpty, error) {
	if err := s.repo.DelRoutesByIDs(in.IDs); err != nil {
		return nil, errorx.StatusError(err)
	}
	return &system.SystemProtoEmpty{}, nil
}

// ListRoute 路由列表
func (s *server) ListRoute(ctx context.Context, in *system.ListRouteRequest) (*system.ListRouteReply, error) {
	page := dto.Pagination{Page: in.Page, PageSize: in.PageSize}
	cond := model.Route{
		Path:    in.Route.Path,
		Method:  in.Route.Method,
		GroupID: uint(in.Route.GroupID),
	}
	list, total, err := s.repo.FindRoutesPagination(page, cond)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	reply := &system.ListRouteReply{Total: total}
	for _, route := range list {
		reply.List = append(reply.List, &system.RouteSchema{
			ID:        uint32(route.ID),
			CreatedAt: route.CreatedAt.UnixMilli(),
			UpdatedAt: route.UpdatedAt.UnixMilli(),
			Path:      route.Path,
			Method:    route.Method,
			Desc:      route.Desc,
			GroupID:   uint32(route.GroupID),
			Group: &system.RouteGroupSchema{
				ID:        uint32(route.Group.ID),
				CreatedAt: route.Group.CreatedAt.UnixMilli(),
				UpdatedAt: route.Group.UpdatedAt.UnixMilli(),
				Name:      route.Group.Name,
			},
		})
	}
	return reply, nil
}

func NewSystemServer() *server {
	migrate.Do(db.NewGormDB())

	return &server{
		repo: repo.NewSystemRepo(),
	}
}
