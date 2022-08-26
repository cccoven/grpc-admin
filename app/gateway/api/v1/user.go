package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"grpc-admin/app/gateway/conf"
	"grpc-admin/app/gateway/helper"
	"grpc-admin/app/gateway/pkg/cache"
	"grpc-admin/app/gateway/pkg/logger"
	"grpc-admin/app/gateway/request"
	"grpc-admin/app/gateway/response"
	"grpc-admin/app/gateway/rpc"
	"grpc-admin/app/user/user"
	"grpc-admin/common/util"
	"strconv"
)

type UserApi struct {
	logger  *zap.SugaredLogger
	rds     *redis.Client
	userRpc user.UserClient
}

// AdminLogin 管理端登录
func (u *UserApi) AdminLogin(c *gin.Context) {
	var req request.AdminLogin
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	var resp response.Login

	if req.Type == 1 {
		if req.Password == "" {
			response.ParameterError(c, "Password 为必填字段")
			return
		}
	}

	if req.Type == 2 {
		if req.Code == "" {
			response.ParameterError(c, "Code 为必填字段")
			return
		}

		// 从缓存取出验证码判断是否正确
		smsCode, _ := u.rds.Get(context.Background(), fmt.Sprintf("smscode_%s", req.Account)).Result()
		if smsCode != req.Code {
			response.Error(c, 5001, "验证码错误")
			return
		}
	}

	ip, _ := helper.GetIPFromRequest(c.Request)
	signInReply, err := u.userRpc.AdminLogin(context.Background(), &user.AdminLoginRequest{
		Account:     req.Account,
		AccountType: 2,
		Ip:          ip,
		Agent:       c.Request.UserAgent(),
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	// 生成 Token
	var roleIDs []uint32
	for _, role := range signInReply.Roles {
		roleIDs = append(roleIDs, role.ID)
	}
	claims := helper.NewClaims(helper.CustomClaims{
		UserID:   signInReply.ID,
		Username: signInReply.Username,
		RoleIDs:  roleIDs,
	})
	jwt, err := helper.GenJwt(claims)
	if err != nil {
		response.DefaultError(c)
		return
	}
	resp.Token = jwt
	resp.ExpiresAt = claims.ExpiresAt
	resp.UserID = claims.CustomClaims.UserID

	response.Success(c, resp)
}

// CreateRole 创建角色
func (u *UserApi) CreateRole(c *gin.Context) {
	var req request.ModifyRole
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	roleReply, err := u.userRpc.ModifyRole(context.Background(), &user.RoleSchema{
		Name:     req.Name,
		Desc:     req.Desc,
		ParentID: uint32(req.ParentID),
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.Role{}
	if err := util.CopyWithTimeConverter(&resp, roleReply); err != nil {
		response.DefaultError(c)
		return
	}

	response.Success(c, resp)
}

// UpdateRole 编辑角色
func (u *UserApi) UpdateRole(c *gin.Context) {
	paramID := c.Param("id")
	roleID, _ := strconv.Atoi(paramID)
	var req request.Role
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	roleReply, err := u.userRpc.ModifyRole(context.Background(), &user.RoleSchema{
		ID:        uint32(roleID),
		Name:      req.Name,
		Desc:      req.Desc,
		ParentID:  uint32(req.ParentID),
		IsDefault: int32(req.IsDefault),
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.Role{}
	if err := util.CopyWithTimeConverter(&resp, roleReply); err != nil {
		response.DefaultError(c)
		return
	}

	response.Success(c, resp)
}

// AuthorizeRole 给角色授权
func (u *UserApi) AuthorizeRole(c *gin.Context) {
	roleID := c.Param("id")
	var req request.AuthorizeRole
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	id, _ := strconv.Atoi(roleID)
	rpcRequest := &user.AuthorizeRoleRequest{
		RoleID:      uint32(id),
		Authorities: []*user.Authority{},
	}
	_ = copier.Copy(rpcRequest, req)
	_, err := u.userRpc.AuthorizeRole(context.Background(), rpcRequest)
	if err != nil {
		response.RpcError(c, err)
		return
	}

	response.Success(c, nil)
}

// AssignRolesToUser 给用户赋予角色
func (u *UserApi) AssignRolesToUser(c *gin.Context) {
	paramID := c.Param("id")
	userID, _ := strconv.Atoi(paramID)
	var req request.AssignRolesToUser
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	_, err := u.userRpc.AssignRoles(context.Background(), &user.AssignRolesRequest{
		UserID: uint32(userID),
		Roles:  req.Roles,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	response.Success(c, nil)
}

// DelRole 删除角色
func (u *UserApi) DelRole(c *gin.Context) {
	paramID := c.Param("id")
	roleID, _ := strconv.Atoi(paramID)

	_, err := u.userRpc.DeleteRole(context.Background(), &user.RoleSchema{ID: uint32(roleID)})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	response.Success(c, nil)
}

// CreateUser 创建用户
func (u *UserApi) CreateUser(c *gin.Context) {
	var req request.ModifyUser
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	var roles []*user.RoleSchema
	for _, id := range req.RoleIDs {
		roles = append(roles, &user.RoleSchema{ID: id})
	}
	reply, err := u.userRpc.ModifyUser(context.Background(), &user.UserSchema{
		ID:       req.ID,
		Username: req.Username,
		Avatar:   req.Avatar,
		Phone:    req.Phone,
		Roles:    roles,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.User{}
	_ = util.CopyWithTimeConverter(&resp, reply)

	response.Success(c, resp)
}

// UserDetail 用户详情
func (u *UserApi) UserDetail(c *gin.Context) {
	paramID := c.Param("id")
	userID, _ := strconv.Atoi(paramID)

	userSchema, err := u.userRpc.GetUser(context.Background(), &user.UserSchema{ID: uint32(userID)})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := &response.User{}
	_ = util.CopyWithTimeConverter(resp, userSchema)

	response.Success(c, resp)
}

// UpdateUser 编辑用户
func (u *UserApi) UpdateUser(c *gin.Context) {
	var req request.ModifyUser
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}
	paramID := c.Param("id")
	userID, _ := strconv.Atoi(paramID)
	var roles []*user.RoleSchema
	for _, id := range req.RoleIDs {
		roles = append(roles, &user.RoleSchema{ID: id})
	}

	userSchema, err := u.userRpc.ModifyUser(context.Background(), &user.UserSchema{
		ID:       uint32(userID),
		Username: req.Username,
		Avatar:   req.Avatar,
		Phone:    req.Phone,
		Roles:    roles,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := &response.User{}
	_ = util.CopyWithTimeConverter(resp, userSchema)

	response.Success(c, resp)
}

// DelUser 删除用户
func (u *UserApi) DelUser(c *gin.Context) {
	paramID := c.Param("id")
	userID, _ := strconv.Atoi(paramID)

	_, err := u.userRpc.DeleteUser(context.Background(), &user.UserSchema{ID: uint32(userID)})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	response.Success(c, nil)
}

// UserList 获取用户列表
func (u *UserApi) UserList(c *gin.Context) {
	var req request.ListUser
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	listReply, err := u.userRpc.ListUser(context.Background(), &user.ListUserRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		User: &user.UserSchema{
			Username: req.Username,
			Phone:    req.Phone,
		},
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.Pagination{
		Page:  req.Page,
		List:  []response.User{},
		Total: listReply.Total,
	}
	_ = util.CopyWithTimeConverter(&resp.List, listReply.List)

	response.Success(c, resp)
}

// ListRole 角色列表
func (u *UserApi) ListRole(c *gin.Context) {
	var req request.ListRole
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	listReply, err := u.userRpc.ListRole(context.Background(), &user.ListRoleRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := response.Pagination{
		Page:  req.Page,
		List:  []response.Role{},
		Total: listReply.Total,
	}
	_ = util.CopyWithTimeConverter(&resp.List, listReply.List)

	response.Success(c, resp)
}

// SetRoleMenus 给角色分配菜单
func (u *UserApi) SetRoleMenus(c *gin.Context) {
	paramID := c.Param("id")
	roleID, _ := strconv.Atoi(paramID)
	var req request.SetRoleMenus
	if err := helper.BindAndValidate(c, &req); err != nil {
		response.ParameterError(c, err.Error())
		return
	}

	_, err := u.userRpc.SetRoleMenus(context.Background(), &user.SetRoleMenusRequest{
		RoleID: uint32(roleID),
		Menus:  req.Menus,
	})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	response.Success(c, nil)
}

// GetRoleAuthorities 获取角色权限
func (u *UserApi) GetRoleAuthorities(c *gin.Context) {
	paramID := c.Param("id")
	roleID, _ := strconv.Atoi(paramID)

	reply, err := u.userRpc.RoleAuthorities(context.Background(), &user.RoleSchema{ID: uint32(roleID)})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := make([]response.Route, len(reply.Authorities))
	_ = util.CopyWithTimeConverter(&resp, reply.Authorities)

	response.Success(c, resp)
}

// GetRoleMenus 获取角色菜单
func (u *UserApi) GetRoleMenus(c *gin.Context) {
	paramID := c.Param("id")
	roleID, _ := strconv.Atoi(paramID)

	reply, err := u.userRpc.RoleMenus(context.Background(), &user.RoleSchema{ID: uint32(roleID)})
	if err != nil {
		response.RpcError(c, err)
		return
	}

	resp := make([]response.Menu, len(reply.Menus))
	_ = util.CopyWithTimeConverter(&resp, reply.Menus)

	response.Success(c, resp)
}

func NewUserApi() *UserApi {
	return &UserApi{
		logger:  logger.NewZapLogger(),
		rds:     cache.NewRedisCache(),
		userRpc: user.NewUserClient(rpc.Discovery(conf.AppConf.UserRpc.Name, conf.AppConf.UserRpc.LoadBalanceMode)),
	}
}
