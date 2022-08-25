package server

import (
	context "context"
	"grpc-admin/app/user/internal/dto"
	"grpc-admin/app/user/internal/model"
	"grpc-admin/app/user/internal/model/migrate"
	"grpc-admin/app/user/internal/pkg/db"
	"grpc-admin/app/user/internal/pkg/errorx"
	"grpc-admin/app/user/internal/repo"
	"grpc-admin/app/user/user"
	"grpc-admin/common/pkg"
	"grpc-admin/common/util"
)

type server struct {
	user.UnimplementedUserServer
	repo repo.IUserRepo
}

func NewUserServer() *server {
	migrate.Do(db.NewGormDB())

	return &server{
		repo: repo.NewUserRepo(),
	}
}

// SetRoleMenus 角色分配菜单
func (s *server) SetRoleMenus(ctx context.Context, in *user.SetRoleMenusRequest) (*user.UserProtoEmpty, error) {
	role, err := s.repo.FindRole(model.Role{GormModel: pkg.GormModel{ID: uint(in.RoleID)}}, false)
	if err != nil {
		return nil, errorx.StatusError(err)
	}
	if err := s.repo.SetRoleMenus(role, in.Menus); err != nil {
		return nil, errorx.StatusError(err)
	}
	return &user.UserProtoEmpty{}, nil
}

// ListRole 角色列表
func (s *server) ListRole(ctx context.Context, in *user.ListRoleRequest) (*user.ListRoleReply, error) {
	page := dto.Pagination{Page: in.Page, PageSize: in.PageSize}
	list, total, err := s.repo.FindRolesPagination(page)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	listReply := &user.ListRoleReply{
		List:  []*user.RoleSchema{},
		Total: total,
	}

	_ = util.CopyWithTimeConverter(&listReply.List, list)

	return listReply, nil
}

// ListUser 用户列表
func (s *server) ListUser(ctx context.Context, in *user.ListUserRequest) (*user.ListUserReply, error) {
	page := dto.Pagination{Page: in.Page, PageSize: in.PageSize}
	cond := model.User{Username: in.User.Username, Phone: in.User.Phone}
	list, total, err := s.repo.FindUsersPagination(page, cond)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	listReply := &user.ListUserReply{
		List:  []*user.UserSchema{},
		Total: total,
	}
	_ = util.CopyWithTimeConverter(&listReply.List, list)

	return listReply, nil
}

// DeleteUser 删除用户
func (s *server) DeleteUser(ctx context.Context, in *user.UserSchema) (*user.UserProtoEmpty, error) {
	u, err := s.repo.FindUser(model.User{GormModel: pkg.GormModel{ID: uint(in.ID)}}, true)
	if err != nil {
		return nil, err
	}

	if err := s.repo.DelUser(u); err != nil {
		return nil, errorx.StatusError(err)
	}

	return &user.UserProtoEmpty{}, nil
}

// GetUser 用户详情
func (s *server) GetUser(ctx context.Context, in *user.UserSchema) (*user.UserSchema, error) {
	userSchema := &user.UserSchema{}
	matchedUser, err := s.repo.FindUser(model.User{GormModel: pkg.GormModel{ID: uint(in.ID)}}, true)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	_ = util.CopyWithTimeConverter(userSchema, matchedUser)
	return userSchema, nil
}

// DeleteRole 删除角色
func (s *server) DeleteRole(ctx context.Context, in *user.RoleSchema) (*user.UserProtoEmpty, error) {
	err := s.repo.DelRole(in.ID)
	if err != nil {
		return nil, errorx.StatusError(err)
	}
	return &user.UserProtoEmpty{}, nil
}

// CheckRolesAuthority 检查角色权限
func (s *server) CheckRolesAuthority(ctx context.Context, in *user.CheckRolesAuthorityRequest) (*user.CheckRolesAuthorityReply, error) {
	authority := dto.Authority{
		Path:   in.Authority.Path,
		Method: in.Authority.Method,
	}
	ok, err := s.repo.CheckRolesAuthority(in.RoleIDs, authority)
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	return &user.CheckRolesAuthorityReply{Ok: ok}, nil
}

// AssignRoles 赋予用户角色
func (s *server) AssignRoles(ctx context.Context, in *user.AssignRolesRequest) (*user.UserProtoEmpty, error) {
	// u := &model.User{
	// 	GormModel: pkg.GormModel{ID: uint(in.UserID)},
	// }
	// if err := s.repo.AssignRoles(u, in.Roles); err != nil {
	// 	return nil, errorx.StatusError(err)
	// }
	
	return &user.UserProtoEmpty{}, nil
}

// AuthorizeRole 角色授权
func (s *server) AuthorizeRole(ctx context.Context, in *user.AuthorizeRoleRequest) (*user.UserProtoEmpty, error) {
	var authorities []dto.Authority
	for _, authority := range in.Authorities {
		authorities = append(authorities, dto.Authority{Path: authority.Path, Method: authority.Method})
	}
	if err := s.repo.AuthorizeRole(in.RoleID, authorities); err != nil {
		return nil, errorx.StatusError(err)
	}

	return &user.UserProtoEmpty{}, nil
}

// ModifyRole 创建/更新角色
func (s *server) ModifyRole(ctx context.Context, in *user.RoleSchema) (*user.RoleSchema, error) {
	var err error
	role := &model.Role{}

	if in.ID == 0 {
		// 存在同名角色
		if found, _ := s.repo.FindRole(model.Role{Name: in.Name}, false); found != nil {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在同名角色"))
		}
		role.Name = in.Name
		role.Desc = in.Desc
		role.ParentID = uint(in.ParentID)
		if err = s.repo.CreateRole(role); err != nil {
			return nil, errorx.StatusError(err)
		}
	} else {
		// 存在同名角色
		if found, _ := s.repo.FindRole(model.Role{Name: in.Name}, false); found != nil && found.ID != uint(in.ID) {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在同名角色"))
		}
		if role, err = s.repo.FindRole(model.Role{GormModel: pkg.GormModel{ID: uint(in.ID)}}, false); err != nil {
			return nil, errorx.StatusError(err)
		}
		role.Name = in.Name
		role.Desc = in.Desc
		role.ParentID = uint(in.ParentID)
		if err = s.repo.UpdateRole(role); err != nil {
			return nil, errorx.StatusError(err)
		}
	}
	if err != nil {
		return nil, errorx.StatusError(err)
	}

	roleSchema := &user.RoleSchema{}
	_ = util.CopyWithTimeConverter(roleSchema, role)

	return roleSchema, nil
}

// AdminLogin 管理端登录
func (s *server) AdminLogin(ctx context.Context, in *user.AdminLoginRequest) (*user.UserSchema, error) {
	userSchema := &user.UserSchema{}
	var u *model.User
	var err error

	if in.AccountType == 2 { // 手机号验证码登录
		u, err = s.repo.FindUser(model.User{Phone: in.Account}, true)
		if err != nil {
			return nil, errorx.StatusError(err)
		}
	}

	// 如果用户在系统中没有角色，表示该用户没有使用管理端的权限
	if len(u.Roles) == 0 {
		return nil, errorx.StatusError(errorx.NewFromCode(errorx.ErrResourceNoPermission))
	}

	// 如果用户没有权限，则赋予一个默认的权限
	// if len(u.Roles) == 0 {
	// 	defaultRole, err := s.repo.FindDefaultRole()
	// 	if err != nil { // 没有默认角色直接返回错误
	// 		return nil, errorx.Default()
	// 	}
	//
	// 	err = s.repo.AssignRoles(uint32(u.ID), []uint32{uint32(defaultRole.ID)})
	// 	if err != nil { // 赋予角色失败直接返回错误
	// 		return nil, errorx.Default()
	// 	}
	//
	// 	roles = append(roles, defaultRole)
	// }
	// _ = util.CopyWithTimeConverter(&userSchema.Roles, roles)

	// 记录登录日志
	signInLog := &model.UserSignInLog{
		UserID: u.ID,
		Ip:     in.Ip,
		Agent:  in.Agent,
	}
	_ = s.repo.CreateSignInLog(signInLog)
	_ = util.CopyWithTimeConverter(userSchema, u)

	return userSchema, nil
}

// ModifyUser 创建/编辑用户
func (s *server) ModifyUser(ctx context.Context, in *user.UserSchema) (*user.UserSchema, error) {
	var err error
	var roleIDs []uint32
	for _, role := range in.Roles {
		roleIDs = append(roleIDs, role.ID)
	}
	u := &model.User{}

	if in.ID == 0 {
		// 存在相同手机号的用户
		if found, _ := s.repo.FindUser(model.User{Phone: in.Phone}, false); found != nil {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在相同手机号用户"))
		}
		u.Username = in.Username
		u.Avatar = in.Avatar
		u.Phone = in.Phone
		u.Roles, _ = s.repo.FindRolesByID(roleIDs)
		if err = s.repo.CreateUser(u); err != nil {
			return nil, errorx.StatusError(err)
		}
	} else {
		// 存在相同手机号的用户
		if found, _ := s.repo.FindUser(model.User{Phone: in.Phone}, false); found != nil && found.ID != uint(in.ID) {
			return nil, errorx.StatusError(errorx.New(errorx.ErrResourceAlreadyExist, "已存在相同手机号用户"))
		}
		// 这里先查一次用户，是为了返回前端完整的用户信息
		if u, err = s.repo.FindUser(model.User{GormModel: pkg.GormModel{ID: uint(in.ID)}}, false); err != nil {
			return nil, errorx.StatusError(err)
		}
		u.Username = in.Username
		u.Avatar = in.Avatar
		u.Phone = in.Phone
		u.Roles, _ = s.repo.FindRolesByID(roleIDs)
		if err = s.repo.UpdateUser(u); err != nil {
			return nil, errorx.StatusError(err)
		}
	}

	userSchema := &user.UserSchema{}
	_ = util.CopyWithTimeConverter(userSchema, u)

	return userSchema, nil
}
