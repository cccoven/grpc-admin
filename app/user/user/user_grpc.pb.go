// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 后台管理登录
	AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*UserSchema, error)
	// 创建/更新用户
	ModifyUser(ctx context.Context, in *UserSchema, opts ...grpc.CallOption) (*UserSchema, error)
	// 用户详情
	GetUser(ctx context.Context, in *UserSchema, opts ...grpc.CallOption) (*UserSchema, error)
	// 删除用户
	DeleteUser(ctx context.Context, in *UserSchema, opts ...grpc.CallOption) (*UserProtoEmpty, error)
	// 用户列表
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
	// 创建/更新角色
	ModifyRole(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*RoleSchema, error)
	// 删除角色
	DeleteRole(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*UserProtoEmpty, error)
	// 角色授权
	AuthorizeRole(ctx context.Context, in *AuthorizeRoleRequest, opts ...grpc.CallOption) (*UserProtoEmpty, error)
	// 角色设置菜单
	SetRoleMenus(ctx context.Context, in *SetRoleMenusRequest, opts ...grpc.CallOption) (*UserProtoEmpty, error)
	// 给用户赋予角色
	AssignRoles(ctx context.Context, in *AssignRolesRequest, opts ...grpc.CallOption) (*UserProtoEmpty, error)
	// 检查角色权限
	CheckRolesAuthority(ctx context.Context, in *CheckRolesAuthorityRequest, opts ...grpc.CallOption) (*CheckRolesAuthorityReply, error)
	// 角色列表
	ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleReply, error)
	// 获取角色权限
	RoleAuthorities(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*RoleAuthoritiesReply, error)
	// 获取角色菜单
	RoleMenus(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*RoleMenusReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) AdminLogin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*UserSchema, error) {
	out := new(UserSchema)
	err := c.cc.Invoke(ctx, "/pb.User/AdminLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ModifyUser(ctx context.Context, in *UserSchema, opts ...grpc.CallOption) (*UserSchema, error) {
	out := new(UserSchema)
	err := c.cc.Invoke(ctx, "/pb.User/ModifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *UserSchema, opts ...grpc.CallOption) (*UserSchema, error) {
	out := new(UserSchema)
	err := c.cc.Invoke(ctx, "/pb.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *UserSchema, opts ...grpc.CallOption) (*UserProtoEmpty, error) {
	out := new(UserProtoEmpty)
	err := c.cc.Invoke(ctx, "/pb.User/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, "/pb.User/ListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ModifyRole(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*RoleSchema, error) {
	out := new(RoleSchema)
	err := c.cc.Invoke(ctx, "/pb.User/ModifyRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteRole(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*UserProtoEmpty, error) {
	out := new(UserProtoEmpty)
	err := c.cc.Invoke(ctx, "/pb.User/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AuthorizeRole(ctx context.Context, in *AuthorizeRoleRequest, opts ...grpc.CallOption) (*UserProtoEmpty, error) {
	out := new(UserProtoEmpty)
	err := c.cc.Invoke(ctx, "/pb.User/AuthorizeRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetRoleMenus(ctx context.Context, in *SetRoleMenusRequest, opts ...grpc.CallOption) (*UserProtoEmpty, error) {
	out := new(UserProtoEmpty)
	err := c.cc.Invoke(ctx, "/pb.User/SetRoleMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AssignRoles(ctx context.Context, in *AssignRolesRequest, opts ...grpc.CallOption) (*UserProtoEmpty, error) {
	out := new(UserProtoEmpty)
	err := c.cc.Invoke(ctx, "/pb.User/AssignRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckRolesAuthority(ctx context.Context, in *CheckRolesAuthorityRequest, opts ...grpc.CallOption) (*CheckRolesAuthorityReply, error) {
	out := new(CheckRolesAuthorityReply)
	err := c.cc.Invoke(ctx, "/pb.User/CheckRolesAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListRole(ctx context.Context, in *ListRoleRequest, opts ...grpc.CallOption) (*ListRoleReply, error) {
	out := new(ListRoleReply)
	err := c.cc.Invoke(ctx, "/pb.User/ListRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RoleAuthorities(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*RoleAuthoritiesReply, error) {
	out := new(RoleAuthoritiesReply)
	err := c.cc.Invoke(ctx, "/pb.User/RoleAuthorities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RoleMenus(ctx context.Context, in *RoleSchema, opts ...grpc.CallOption) (*RoleMenusReply, error) {
	out := new(RoleMenusReply)
	err := c.cc.Invoke(ctx, "/pb.User/RoleMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 后台管理登录
	AdminLogin(context.Context, *AdminLoginRequest) (*UserSchema, error)
	// 创建/更新用户
	ModifyUser(context.Context, *UserSchema) (*UserSchema, error)
	// 用户详情
	GetUser(context.Context, *UserSchema) (*UserSchema, error)
	// 删除用户
	DeleteUser(context.Context, *UserSchema) (*UserProtoEmpty, error)
	// 用户列表
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	// 创建/更新角色
	ModifyRole(context.Context, *RoleSchema) (*RoleSchema, error)
	// 删除角色
	DeleteRole(context.Context, *RoleSchema) (*UserProtoEmpty, error)
	// 角色授权
	AuthorizeRole(context.Context, *AuthorizeRoleRequest) (*UserProtoEmpty, error)
	// 角色设置菜单
	SetRoleMenus(context.Context, *SetRoleMenusRequest) (*UserProtoEmpty, error)
	// 给用户赋予角色
	AssignRoles(context.Context, *AssignRolesRequest) (*UserProtoEmpty, error)
	// 检查角色权限
	CheckRolesAuthority(context.Context, *CheckRolesAuthorityRequest) (*CheckRolesAuthorityReply, error)
	// 角色列表
	ListRole(context.Context, *ListRoleRequest) (*ListRoleReply, error)
	// 获取角色权限
	RoleAuthorities(context.Context, *RoleSchema) (*RoleAuthoritiesReply, error)
	// 获取角色菜单
	RoleMenus(context.Context, *RoleSchema) (*RoleMenusReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) AdminLogin(context.Context, *AdminLoginRequest) (*UserSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedUserServer) ModifyUser(context.Context, *UserSchema) (*UserSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUser not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *UserSchema) (*UserSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *UserSchema) (*UserProtoEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) ListUser(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServer) ModifyRole(context.Context, *RoleSchema) (*RoleSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyRole not implemented")
}
func (UnimplementedUserServer) DeleteRole(context.Context, *RoleSchema) (*UserProtoEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedUserServer) AuthorizeRole(context.Context, *AuthorizeRoleRequest) (*UserProtoEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeRole not implemented")
}
func (UnimplementedUserServer) SetRoleMenus(context.Context, *SetRoleMenusRequest) (*UserProtoEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoleMenus not implemented")
}
func (UnimplementedUserServer) AssignRoles(context.Context, *AssignRolesRequest) (*UserProtoEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignRoles not implemented")
}
func (UnimplementedUserServer) CheckRolesAuthority(context.Context, *CheckRolesAuthorityRequest) (*CheckRolesAuthorityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRolesAuthority not implemented")
}
func (UnimplementedUserServer) ListRole(context.Context, *ListRoleRequest) (*ListRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRole not implemented")
}
func (UnimplementedUserServer) RoleAuthorities(context.Context, *RoleSchema) (*RoleAuthoritiesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAuthorities not implemented")
}
func (UnimplementedUserServer) RoleMenus(context.Context, *RoleSchema) (*RoleMenusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleMenus not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/AdminLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AdminLogin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ModifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ModifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/ModifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ModifyUser(ctx, req.(*UserSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*UserSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*UserSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/ListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ModifyRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ModifyRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/ModifyRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ModifyRole(ctx, req.(*RoleSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteRole(ctx, req.(*RoleSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AuthorizeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AuthorizeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/AuthorizeRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AuthorizeRole(ctx, req.(*AuthorizeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetRoleMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoleMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetRoleMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/SetRoleMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetRoleMenus(ctx, req.(*SetRoleMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AssignRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AssignRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/AssignRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AssignRoles(ctx, req.(*AssignRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckRolesAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRolesAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckRolesAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/CheckRolesAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckRolesAuthority(ctx, req.(*CheckRolesAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/ListRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListRole(ctx, req.(*ListRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RoleAuthorities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RoleAuthorities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/RoleAuthorities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RoleAuthorities(ctx, req.(*RoleSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RoleMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RoleMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/RoleMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RoleMenus(ctx, req.(*RoleSchema))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLogin",
			Handler:    _User_AdminLogin_Handler,
		},
		{
			MethodName: "ModifyUser",
			Handler:    _User_ModifyUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _User_ListUser_Handler,
		},
		{
			MethodName: "ModifyRole",
			Handler:    _User_ModifyRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _User_DeleteRole_Handler,
		},
		{
			MethodName: "AuthorizeRole",
			Handler:    _User_AuthorizeRole_Handler,
		},
		{
			MethodName: "SetRoleMenus",
			Handler:    _User_SetRoleMenus_Handler,
		},
		{
			MethodName: "AssignRoles",
			Handler:    _User_AssignRoles_Handler,
		},
		{
			MethodName: "CheckRolesAuthority",
			Handler:    _User_CheckRolesAuthority_Handler,
		},
		{
			MethodName: "ListRole",
			Handler:    _User_ListRole_Handler,
		},
		{
			MethodName: "RoleAuthorities",
			Handler:    _User_RoleAuthorities_Handler,
		},
		{
			MethodName: "RoleMenus",
			Handler:    _User_RoleMenus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
