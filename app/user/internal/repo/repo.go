package repo

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"grpc-admin/app/user/internal/dto"
	"grpc-admin/app/user/internal/model"
	"grpc-admin/app/user/internal/pkg/cache"
	"grpc-admin/app/user/internal/pkg/db"
	"grpc-admin/app/user/internal/pkg/errorx"
	"grpc-admin/app/user/internal/pkg/logger"
	"grpc-admin/app/user/internal/pkg/rbac"
	"grpc-admin/common/pkg"
	"grpc-admin/common/util"
	"strconv"
)

type IUserRepo interface {
	FindUser(cond model.User, withRoles bool) (*model.User, error)
	CreateSignInLog(signInLog *model.UserSignInLog) error
	CreateRole(role *model.Role) error
	UpdateRole(role *model.Role) error
	AuthorizeRole(roleID uint32, authorities []dto.Authority) error
	AssignRoles(user *model.User) error
	FindDefaultRole() (model.Role, error)
	CheckRolesAuthority(roleIDs []uint32, authority dto.Authority) (bool, error)
	DelRole(id uint32) error
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DelUser(user *model.User) error
	FindUsersPagination(page dto.Pagination, cond model.User) ([]model.User, int64, error)
	FindRolesPagination(page dto.Pagination) ([]model.Role, int64, error)
	FindRolesByID(roleIDs []uint32) ([]model.Role, error)
	FindRole(cond model.Role, withUsers bool) (*model.Role, error)
	SetRoleMenus(role *model.Role, menuIDs []uint32) error
	GetRoutesByRoleID(roleID uint32) ([]model.Route, error)
	GetMenusByRoleID(roleID uint32) ([]model.Menu, error)
}

type UserRepo struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *zap.SugaredLogger

	// casbin
	casbinEnforcer *casbin.SyncedEnforcer
}

// GetMenusByRoleID 根据角色 ID 获取菜单
func (u *UserRepo) GetMenusByRoleID(roleID uint32) ([]model.Menu, error) {
	var menus []model.Menu
	role, err := u.FindRole(model.Role{GormModel: pkg.GormModel{ID: uint(roleID)}}, false)
	if err != nil {
		return nil, err
	}

	if err = u.db.Model(&role).Association("Menus").Find(&menus); err != nil {
		return nil, errorx.Default()
	}

	return menus, nil
}

// GetRoutesByRoleID 根据角色 ID 获取路由
func (u *UserRepo) GetRoutesByRoleID(roleID uint32) ([]model.Route, error) {
	var routes []model.Route
	_, err := u.FindRole(model.Role{GormModel: pkg.GormModel{ID: uint(roleID)}}, false)
	if err != nil {
		return nil, err
	}

	// 从 casbin 表中找出该角色的路由权限
	filteredPolicy := u.casbinEnforcer.GetFilteredPolicy(0, strconv.Itoa(int(roleID)))

	builder := u.db.Model(&[]model.Route{})
	for _, policy := range filteredPolicy {
		path, method := policy[1], policy[2]
		builder = builder.Or("`path` = ? AND `method` = ?", path, method)
	}
	if err := builder.Find(&routes).Error; err != nil {
		return nil, errorx.Default()
	}

	return routes, nil
}

// SetRoleMenus 角色分配菜单
func (u *UserRepo) SetRoleMenus(role *model.Role, menuIDs []uint32) error {
	var menus []model.Menu
	for _, id := range menuIDs {
		menus = append(menus, model.Menu{ID: uint(id)})
	}
	if err := u.db.Model(&role).Association("Menus").Replace(menus); err != nil {
		return errorx.Default()
	}
	return nil
}

// FindRole 根据条件查找角色
func (u *UserRepo) FindRole(cond model.Role, withUsers bool) (*model.Role, error) {
	var role *model.Role
	builder := u.db.Model(&role)

	if cond.ID != 0 {
		builder = builder.Where("id", cond.ID)
	}

	if cond.Name != "" {
		builder = builder.Where("name", cond.Name)
	}

	if withUsers {
		builder = builder.Preload("Users")
	}

	if err := builder.First(&role).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewFromCode(errorx.ErrResourceNoExist)
	}

	return role, nil
}

// FindRolesByID 根据多个 ID 找出对应角色
func (u *UserRepo) FindRolesByID(roleIDs []uint32) ([]model.Role, error) {
	var roles []model.Role
	if err := u.db.Model(&[]model.Role{}).Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
		u.logger.Errorf("查找角色列表失败：%s", err.Error())
		return nil, errorx.Default()
	}
	return roles, nil
}

func (u *UserRepo) buildChildren(parent []model.Role) error {
	var children []model.Role
	var parentIDs []uint
	parentMap := make(map[uint][]model.Role)
	for _, p := range parent {
		parentIDs = append(parentIDs, p.ID)
	}

	if err := u.db.Model(&model.Role{}).Where("parent_id IN ?", parentIDs).Find(&children).Error; err != nil {
		return errorx.Default()
	}

	for _, child := range children {
		parentMap[child.ParentID] = append(parentMap[child.ParentID], child)
	}

	for i := range parent {
		parent[i].Children = parentMap[parent[i].ID]
		if parent[i].Children != nil {
			if err := u.buildChildren(parent[i].Children); err != nil {
				return err
			}
		}
	}

	return nil
}

// FindRolesPagination 分页角色列表
func (u *UserRepo) FindRolesPagination(page dto.Pagination) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64
	builder := u.db.Model(&roles).Where("parent_id", 0)

	err := builder.
		Scopes(util.Paginate(int(page.Page), int(page.PageSize))).
		Count(&total).
		Find(&roles).
		Error
	if err != nil {
		u.logger.Errorf("查询角色列表失败：%s", err.Error())
		return nil, 0, errorx.Default()
	}

	if err = u.buildChildren(roles); err != nil {
		u.logger.Errorf("查询角色列表失败：%s", err.Error())
		return nil, 0, err
	}

	return roles, total, nil
}

// FindUsersPagination 分页用户列表
func (u *UserRepo) FindUsersPagination(page dto.Pagination, cond model.User) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	builder := u.db.Model(&users)

	if cond.Username != "" {
		builder = builder.Where("username LIKE ?", "%"+cond.Username+"%")
	}

	if cond.Phone != "" {
		builder = builder.Where("phone LIKE ?", "%"+cond.Phone+"%")
	}

	err := builder.
		Scopes(util.Paginate(int(page.Page), int(page.PageSize))).
		Preload("Roles").
		Count(&total).
		Find(&users).
		Error
	if err != nil {
		u.logger.Errorf("查询用户列表失败：%s", err.Error())
		return nil, 0, errorx.Default()
	}

	return users, total, nil
}

// DelUser 删除用户
func (u *UserRepo) DelUser(user *model.User) error {
	if err := u.db.Delete(&user).Error; err != nil {
		u.logger.Errorf("删除用户失败：%s", err.Error())
		return errorx.Default()
	}

	// 删除角色关联
	err := u.db.Model(&user).Association("Roles").Delete(user.Roles)
	if err != nil {
		u.logger.Errorf("删除用户关联失败：%s", err.Error())
		return errorx.Default()
	}

	// 更新 casbin
	ok, err := u.casbinEnforcer.RemoveFilteredGroupingPolicy(0, strconv.Itoa(int(user.ID)))
	if !ok || err != nil {
		return errorx.Default()
	}

	return nil
}

// UpdateUser 编辑用户
func (u *UserRepo) UpdateUser(user *model.User) error {
	if err := u.db.Updates(&user).Error; err != nil {
		u.logger.Errorf("编辑用户失败：%s", err.Error())
		return errorx.Default()
	}
	// 给用户分配角色
	if err := u.AssignRoles(user); err != nil {
		return err
	}
	return nil
}

// CreateUser 创建用户
func (u *UserRepo) CreateUser(user *model.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		u.logger.Errorf("创建用户失败：%s", err.Error())
		return errorx.Default()
	}
	// 给用户分配角色
	if err = u.AssignRoles(user); err != nil {
		return err
	}
	return nil
}

// DelRole 删除角色
func (u *UserRepo) DelRole(id uint32) error {
	// role, err := u.FindRoleWithUsers(id)
	role, err := u.FindRole(model.Role{GormModel: pkg.GormModel{ID: uint(id)}}, true)
	if err != nil { // 角色不存在（中途被别人删了）
		return err
	}
	if len(role.Users) != 0 { // 角色正在被使用
		return errorx.NewFromCode(errorx.ErrResourceBeingUsed)
	}

	// 从角色表中查出 parent_id = 当前角色 id 的第一条记录
	// 存在则表示该角色拥有子角色
	err = u.db.Model(&model.Role{}).Where("parent_id", id).First(&model.Role{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errorx.NewFromCode(errorx.ErrResourceBeingUsed)
	}

	err = u.db.Transaction(func(tx *gorm.DB) error {
		// 删除角色（物理删除）
		if err = tx.Unscoped().Delete(&role).Error; err != nil {
			return err
		}

		// 更新角色与用户关联表
		// 关联模式删除只会删除引用，不会从数据库中删除这些对象，而这里是要去中间表删除用户和角色的关联，需要使用物理删除
		// err = u.db.Model(&role).Association("Users").Delete(role.Users)
		// TODO 上面判断了如果存在用户则无法删除，这里暂时没有意义
		if err = u.db.Where("role_id", id).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}

		// 删除角色权限
		_, err := u.casbinEnforcer.RemoveFilteredPolicy(0, strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		u.logger.Errorf("删除角色失败：%s", err.Error())
		return errorx.Default()
	}

	return nil
}

// CheckRolesAuthority 检查角色权限
func (u *UserRepo) CheckRolesAuthority(roleIDs []uint32, authority dto.Authority) (bool, error) {
	var requests [][]any
	for _, roleID := range roleIDs {
		requests = append(requests, []any{strconv.Itoa(int(roleID)), authority.Path, authority.Method})
	}
	// 检查用户拥有角色的所有权限
	bools, err := u.casbinEnforcer.BatchEnforce(requests)
	// ok, err := u.casbinEnforcer.Enforce(roleID, authority.Path, authority.Method)
	if err != nil {
		u.logger.Errorf("查询角色权限失败：%s", err.Error())
		return false, errorx.Default()
	}

	// 如果有一个角色的权限匹配上则直接返回持有权限
	for _, ok := range bools {
		if ok {
			return ok, nil
		}
	}

	return false, nil
}

// FindDefaultRole 查找默认角色
func (u *UserRepo) FindDefaultRole() (model.Role, error) {
	var role model.Role
	err := u.db.Model(&role).Where("is_default", 1).First(&role).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return role, errorx.NewFromCode(errorx.ErrResourceNoExist)
	}
	return role, nil
}

// AssignRoles 分配角色到用户
func (u *UserRepo) AssignRoles(user *model.User) error {
	var rules [][]string
	for _, role := range user.Roles {
		rules = append(rules, []string{strconv.Itoa(int(user.ID)), strconv.Itoa(int(role.ID))})
	}

	// 在 casbin 表中定义 g 表示用户对应的角色
	// 如果在 casbin 表中用户已经存在角色，那么就需要使用修改 api
	// 这里为了简单起见，每次给用户分配角色时都先清空 casbin 表中这个用户的角色
	// RemoveFilteredGroupingPolicy(0, strconv.Itoa(int(user.ID))) 表示删除所有 V0 列等于 userID 的行
	_, err := u.casbinEnforcer.RemoveFilteredGroupingPolicy(0, strconv.Itoa(int(user.ID)))
	if err != nil {
		u.logger.Errorf("清空用户角色失败：%s", err.Error())
		return errorx.Default()
	}

	_, err = u.casbinEnforcer.AddGroupingPolicies(rules)
	if err != nil {
		u.logger.Errorf("用户分配角色失败：%s", err.Error())
		return errorx.Default()
	}

	// User 跟 Role 是多对多的关联关系（一个用户有多个角色，一个角色下有多个用户）
	// 可以使用 Association 更新关联（前提是在 User Model 中显示关联了 Role
	// Association 会自动更新 User 和 Role 的中间表
	if err := u.db.Model(&user).Association("Roles").Replace(user.Roles); err != nil {
		u.logger.Errorf("用户分配角色失败：%s", err.Error())
		return errorx.Default()
	}

	return nil
}

// AuthorizeRole 给角色赋予权限
func (u *UserRepo) AuthorizeRole(roleID uint32, authorities []dto.Authority) error {
	rid := strconv.Itoa(int(roleID))
	var rules [][]string
	for _, authority := range authorities {
		rules = append(rules, []string{rid, authority.Path, authority.Method})
	}

	_, err := u.casbinEnforcer.RemoveFilteredPolicy(0, strconv.Itoa(int(roleID)))
	if err != nil {
		u.logger.Errorf("清空角色权限失败：%s", err.Error())
		return errorx.Default()
	}
	_, err = u.casbinEnforcer.AddPolicies(rules)
	if err != nil {
		u.logger.Errorf("角色添加权限失败：%s", err.Error())
		return errorx.Default()
	}

	return nil
}

// CreateRole 创建角色
func (u *UserRepo) CreateRole(role *model.Role) error {
	if err := u.db.Create(&role).Error; err != nil {
		u.logger.Errorf("创建角色失败: %s", err.Error())
		return errorx.Default()
	}
	return nil
}

// UpdateRole 更新角色
func (u *UserRepo) UpdateRole(role *model.Role) error {
	if role.IsDefault == 1 {
		// 如果将这个角色设为默认角色，那么先修改之前的角色
		if err := u.db.Model(&model.Role{}).
			Where("id != ?", role.ID).
			Where("is_default", 1).
			Update("is_default", 0).
			Error; err != nil {
			u.logger.Errorf("修改默认角色失败：:%s", err.Error())
			return errorx.Default()
		}
	}
	// 默认情况下，gorm 的 Update 只会更新非零值
	if err := u.db.Model(&model.Role{}).Where("id", role.ID).Updates(map[string]any{
		"name":       role.Name,
		"desc":       role.Desc,
		"parent_id":  role.ParentID,
		"is_default": role.IsDefault,
	}).Error; err != nil {
		u.logger.Errorf("编辑角色失败: %s", err.Error())
		return errorx.Default()
	}
	return nil
}

// CreateSignInLog 创建登录日志
func (u *UserRepo) CreateSignInLog(signInLog *model.UserSignInLog) error {
	if err := u.db.Create(signInLog).Error; err != nil {
		u.logger.Errorf("记录登录日志失败: %s", err.Error())
		return errorx.Default()
	}
	return nil
}

// FindUser 查找满足条件的用户
func (u *UserRepo) FindUser(cond model.User, withRoles bool) (*model.User, error) {
	var user *model.User
	builder := u.db.Model(user)

	if cond.ID != 0 {
		builder = builder.Where("id", cond.ID)
	}

	if cond.Username != "" {
		builder = builder.Where("username", cond.Username)
	}

	if cond.Phone != "" {
		builder = builder.Where("phone", cond.Phone)
	}

	if withRoles {
		builder = builder.Preload("Roles")
	}

	if err := builder.First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.NewFromCode(errorx.ErrResourceNoExist)
	}

	return user, nil
}

func NewUserRepo() IUserRepo {
	return &UserRepo{
		db:     db.NewGormDB(),
		rds:    cache.NewRedisCache(),
		logger: logger.NewZapLogger(),

		casbinEnforcer: rbac.NewCasbinEnforcer(db.NewGormDB(), logger.NewZapLogger()),
	}
}
