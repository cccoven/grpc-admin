package repo

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"grpc-admin/app/system/internal/conf"
	"grpc-admin/app/system/internal/dto"
	"grpc-admin/app/system/internal/model"
	"grpc-admin/app/system/internal/pkg/cache"
	"grpc-admin/app/system/internal/pkg/db"
	"grpc-admin/app/system/internal/pkg/errorx"
	"grpc-admin/app/system/internal/pkg/logger"
	"grpc-admin/common/util"
)

type ISystemRepo interface {
	FindRouteGroup(cond model.RouteGroup) (*model.RouteGroup, error)
	CreateRouteGroup(routeGroup *model.RouteGroup) error
	UpdateRouteGroup(routeGroup *model.RouteGroup) error
	DelRouteGroupByIDs(ids []uint32) error
	FindRouteGroupsPagination(page dto.Pagination, cond model.RouteGroup) ([]model.RouteGroup, int64, error)
	FindRoute(cond model.Route, withGroup bool) (*model.Route, error)
	CreateRoute(route *model.Route) error
	UpdateRoute(route *model.Route) error
	DelRoutesByIDs(ids []uint32) error
	FindRoutesPagination(page dto.Pagination, cond model.Route) ([]model.Route, int64, error)
	FindMenu(menu model.Menu, withChildren bool) (*model.Menu, error)
	CreateMenu(menu *model.Menu) error
	UpdateMenu(menu *model.Menu) error
	DelMenu(ids []uint32) error
	FindMenusPagination(page dto.Pagination, cond model.Menu) ([]model.Menu, int64, error)
}

type SystemRepo struct {
	config conf.Conf
	db     *gorm.DB
	rds    *redis.Client
	logger *zap.SugaredLogger
}

func (s *SystemRepo) buildChildren(parent []model.Menu) error {
	var children []model.Menu
	var parentIDs []uint
	parentMap := make(map[uint][]model.Menu)
	for _, p := range parent {
		parentIDs = append(parentIDs, p.ID)
	}

	if err := s.db.Model(&model.Menu{}).Where("parent_id IN ?", parentIDs).Debug().Find(&children).Error; err != nil {
		return errorx.Default()
	}

	for _, child := range children {
		parentMap[child.ParentID] = append(parentMap[child.ParentID], child)
	}

	for i := range parent {
		parent[i].Children = parentMap[parent[i].ID]
		if parent[i].Children != nil {
			if err := s.buildChildren(parent[i].Children); err != nil {
				return err
			}
		}
	}

	return nil
}

// FindMenusPagination 菜单分页
func (s *SystemRepo) FindMenusPagination(page dto.Pagination, cond model.Menu) ([]model.Menu, int64, error) {
	var menus []model.Menu
	var total int64
	builder := s.db.Model(&[]model.Menu{}).Where("parent_id", 0)

	if cond.Path != "" {
		builder = builder.Where("path LIKE ?", "%"+cond.Path+"%")
	}

	if cond.Name != "" {
		builder = builder.Where("name LIKE ?", "%"+cond.Name+"%")
	}

	err := builder.
		Scopes(util.Paginate(int(page.Page), int(page.PageSize))).
		Count(&total).
		Find(&menus).
		Error
	if err != nil {
		s.logger.Errorf("查询菜单列表失败：%s", err.Error())
		return nil, 0, errorx.Default()
	}

	if err = s.buildChildren(menus); err != nil {
		s.logger.Errorf("查询菜单列表失败：%s", err.Error())
		return nil, 0, err
	}

	return menus, total, nil
}

// DelMenu 批量删除菜单
func (s *SystemRepo) DelMenu(ids []uint32) error {
	if err := s.db.Where("id IN ?", ids).Delete(&[]model.Menu{}).Error; err != nil {
		s.logger.Errorf("删除菜单失败：%s", err.Error())
		return errorx.Default()
	}

	// 删除所有子菜单
	if err := s.db.Where("parent_id IN ?", ids).Delete(&[]model.Menu{}).Error; err != nil {
		s.logger.Errorf("删除菜单失败：%s", err.Error())
		return errorx.Default()
	}

	return nil
}

// UpdateMenu 更新菜单
func (s *SystemRepo) UpdateMenu(menu *model.Menu) error {
	if err := s.db.Updates(menu).Error; err != nil {
		s.logger.Errorf("更新菜单失败：%s", err.Error())
		return errorx.Default()
	}
	return nil
}

// CreateMenu 创建菜单
func (s *SystemRepo) CreateMenu(menu *model.Menu) error {
	if err := s.db.Create(menu).Error; err != nil {
		s.logger.Errorf("创建菜单失败：%s", err.Error())
		return errorx.Default()
	}
	return nil
}

// FindMenu 查找菜单
func (s *SystemRepo) FindMenu(cond model.Menu, withChildren bool) (*model.Menu, error) {
	var menu *model.Menu
	builder := s.db.Model(&menu)

	if cond.ID != 0 {
		builder = builder.Where("id", cond.ID)
	}

	if cond.Path != "" {
		builder = builder.Where("path", cond.Path)
	}

	if cond.Name != "" {
		builder = builder.Where("name", cond.Name)
	}

	if err := builder.First(&menu).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewFromCode(errorx.ErrResourceNoExist)
	}

	if withChildren {
		menus := []model.Menu{*menu}
		_ = s.buildChildren(menus)
		menu = &menus[0]
	}

	return menu, nil
}

// UpdateRoute 更新路由
func (s *SystemRepo) UpdateRoute(route *model.Route) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(&route).Error; err != nil {
			s.logger.Errorf("修改路由失败：%s", err.Error())
			return errorx.Default()
		}
		group, err := s.FindRouteGroup(model.RouteGroup{ID: route.GroupID})
		if err != nil {
			s.logger.Errorf("修改路由失败：%s", err.Error())
			return errorx.New(errorx.ErrResourceNoExist, "分组不存在")
		}
		if err := tx.Model(&route).Association("Group").Replace(group); err != nil {
			s.logger.Errorf("修改路由失败：%s", err.Error())
			return errorx.Default()
		}
		return nil
	})
}

// DelRoutesByIDs 批量删除路由
func (s *SystemRepo) DelRoutesByIDs(ids []uint32) error {
	if err := s.db.Where("id IN ?", ids).Unscoped().Delete(&[]model.Route{}).Error; err != nil {
		s.logger.Errorf("删除路由失败：%s", err.Error())
		return errorx.Default()
	}
	return nil
}

// FindRoutesPagination 分页获取路由
func (s *SystemRepo) FindRoutesPagination(page dto.Pagination, cond model.Route) ([]model.Route, int64, error) {
	var routes []model.Route
	var total int64
	builder := s.db.Model(&routes)

	if cond.Path != "" {
		builder = builder.Where("path LIKE ?", "%"+cond.Path+"%")
	}

	if cond.Method != "" {
		builder = builder.Where("method LIKE ?", "%"+cond.Method+"%")
	}

	if cond.GroupID != 0 {
		builder = builder.Where("group_id", cond.GroupID)
	}

	err := builder.
		Preload("Group").
		Scopes(util.Paginate(int(page.Page), int(page.PageSize))).
		Count(&total).
		Find(&routes).
		Error
	if err != nil {
		s.logger.Errorf("查询路由组列表失败：%s", err.Error())
		return nil, 0, errorx.Default()
	}
	return routes, total, nil
}

// CreateRoute 创建路由
func (s *SystemRepo) CreateRoute(route *model.Route) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&route).Error; err != nil {
			s.logger.Errorf("创建路由失败：%s", err.Error())
			return errorx.Default()
		}
		group, err := s.FindRouteGroup(model.RouteGroup{ID: route.GroupID})
		if err != nil {
			s.logger.Errorf("创建路由失败：%s", err.Error())
			return errorx.New(errorx.ErrResourceNoExist, "分组不存在")
		}
		if err := tx.Model(&route).Association("Group").Append(group); err != nil {
			s.logger.Errorf("创建路由失败：%s", err.Error())
			return errorx.Default()
		}
		return nil
	})
}

// FindRoute 查找路由
func (s *SystemRepo) FindRoute(cond model.Route, withGroup bool) (*model.Route, error) {
	var route *model.Route
	builder := s.db.Model(&route)

	if cond.ID != 0 {
		builder = builder.Where("id", cond.ID)
	}

	if cond.Path != "" {
		builder = builder.Where("path", cond.Path)
	}

	if cond.Method != "" {
		builder = builder.Where("method", cond.Method)
	}

	if withGroup {
		builder = builder.Preload("Group")
	}

	if err := builder.First(&route).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewFromCode(errorx.ErrResourceNoExist)
	}

	return route, nil
}

// FindRouteGroupsPagination 分页获取路由组
func (s *SystemRepo) FindRouteGroupsPagination(page dto.Pagination, cond model.RouteGroup) ([]model.RouteGroup, int64, error) {
	var groups []model.RouteGroup
	var total int64
	builder := s.db.Model(&groups)

	if cond.Name != "" {
		builder = builder.Where("name LIKE ?", "%"+cond.Name+"%")
	}

	err := builder.
		Scopes(util.Paginate(int(page.Page), int(page.PageSize))).
		Count(&total).
		Find(&groups).
		Error
	if err != nil {
		s.logger.Errorf("查询路由组列表失败：%s", err.Error())
		return nil, 0, errorx.Default()
	}
	return groups, total, nil
}

// DelRouteGroupByIDs 批量删除路由组
func (s *SystemRepo) DelRouteGroupByIDs(ids []uint32) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		// 先删除该组下的所有路由
		if err := tx.Unscoped().Delete(&model.Route{}, "group_id IN ?", ids).Error; err != nil {
			s.logger.Errorf("删除路由组失败：%s", err.Error())
			return errorx.Default()
		}

		if err := tx.Where("id IN ?", ids).Unscoped().Delete(&[]model.RouteGroup{}).Error; err != nil {
			s.logger.Errorf("删除路由组失败：%s", err.Error())
			return errorx.Default()
		}
		return nil
	})
}

// UpdateRouteGroup 更新路由组
func (s *SystemRepo) UpdateRouteGroup(routeGroup *model.RouteGroup) error {
	if err := s.db.Updates(&routeGroup).Error; err != nil {
		s.logger.Errorf("创建路由组失败：%s", err.Error())
		return errorx.Default()
	}
	return nil
}

// CreateRouteGroup 创建路由组
func (s *SystemRepo) CreateRouteGroup(routeGroup *model.RouteGroup) error {
	if err := s.db.Create(&routeGroup).Error; err != nil {
		s.logger.Errorf("创建路由组失败：%s", err.Error())
		return errorx.Default()
	}
	return nil
}

// FindRouteGroup 查找路由组
func (s *SystemRepo) FindRouteGroup(cond model.RouteGroup) (*model.RouteGroup, error) {
	var routerGroup *model.RouteGroup
	builder := s.db.Model(&routerGroup)

	if cond.ID != 0 {
		builder = builder.Where("id", cond.ID)
	}

	if cond.Name != "" {
		builder = builder.Where("name", cond.Name)
	}

	if err := builder.First(&routerGroup).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewFromCode(errorx.ErrResourceNoExist)
	}

	return routerGroup, nil
}

func NewSystemRepo() ISystemRepo {
	return &SystemRepo{
		db:     db.NewGormDB(),
		rds:    cache.NewRedisCache(),
		logger: logger.NewZapLogger(),
	}
}
