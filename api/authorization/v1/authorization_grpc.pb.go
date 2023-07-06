// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: authorization/v1/authorization.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthorizationService_CheckAuthorization_FullMethodName = "/authorization.v1.AuthorizationService/CheckAuthorization"
	AuthorizationService_GetApiListAll_FullMethodName      = "/authorization.v1.AuthorizationService/GetApiListAll"
	AuthorizationService_GetApiList_FullMethodName         = "/authorization.v1.AuthorizationService/GetApiList"
	AuthorizationService_CreateApi_FullMethodName          = "/authorization.v1.AuthorizationService/CreateApi"
	AuthorizationService_UpdateApi_FullMethodName          = "/authorization.v1.AuthorizationService/UpdateApi"
	AuthorizationService_DeleteApi_FullMethodName          = "/authorization.v1.AuthorizationService/DeleteApi"
	AuthorizationService_GetRoleAll_FullMethodName         = "/authorization.v1.AuthorizationService/GetRoleAll"
	AuthorizationService_GetRoleList_FullMethodName        = "/authorization.v1.AuthorizationService/GetRoleList"
	AuthorizationService_CreateRole_FullMethodName         = "/authorization.v1.AuthorizationService/CreateRole"
	AuthorizationService_UpdateRole_FullMethodName         = "/authorization.v1.AuthorizationService/UpdateRole"
	AuthorizationService_DeleteRole_FullMethodName         = "/authorization.v1.AuthorizationService/DeleteRole"
)

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	CheckAuthorization(ctx context.Context, in *CheckAuthorizationReq, opts ...grpc.CallOption) (*CheckAuthorizationReply, error)
	// ================== Api ==================
	// Api列表 - 所有
	GetApiListAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetApiListAllReply, error)
	// Api列表 - 分页
	GetApiList(ctx context.Context, in *GetApiListReq, opts ...grpc.CallOption) (*GetApiListReply, error)
	// Api创建
	CreateApi(ctx context.Context, in *CreateApiReq, opts ...grpc.CallOption) (*Api, error)
	// Api更新
	UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...grpc.CallOption) (*CheckReply, error)
	// Api删除
	DeleteApi(ctx context.Context, in *DeleteApiReq, opts ...grpc.CallOption) (*CheckReply, error)
	// ================== 角色 ==================
	// 全部角色
	GetRoleAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRoleAllReply, error)
	GetRoleList(ctx context.Context, in *GetRoleListReq, opts ...grpc.CallOption) (*GetRoleListReply, error)
	// 角色创建
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*Role, error)
	// 角色更新
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*CheckReply, error)
	// 角色删除
	DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*CheckReply, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) CheckAuthorization(ctx context.Context, in *CheckAuthorizationReq, opts ...grpc.CallOption) (*CheckAuthorizationReply, error) {
	out := new(CheckAuthorizationReply)
	err := c.cc.Invoke(ctx, AuthorizationService_CheckAuthorization_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetApiListAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetApiListAllReply, error) {
	out := new(GetApiListAllReply)
	err := c.cc.Invoke(ctx, AuthorizationService_GetApiListAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetApiList(ctx context.Context, in *GetApiListReq, opts ...grpc.CallOption) (*GetApiListReply, error) {
	out := new(GetApiListReply)
	err := c.cc.Invoke(ctx, AuthorizationService_GetApiList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CreateApi(ctx context.Context, in *CreateApiReq, opts ...grpc.CallOption) (*Api, error) {
	out := new(Api)
	err := c.cc.Invoke(ctx, AuthorizationService_CreateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) UpdateApi(ctx context.Context, in *UpdateApiReq, opts ...grpc.CallOption) (*CheckReply, error) {
	out := new(CheckReply)
	err := c.cc.Invoke(ctx, AuthorizationService_UpdateApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) DeleteApi(ctx context.Context, in *DeleteApiReq, opts ...grpc.CallOption) (*CheckReply, error) {
	out := new(CheckReply)
	err := c.cc.Invoke(ctx, AuthorizationService_DeleteApi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetRoleAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRoleAllReply, error) {
	out := new(GetRoleAllReply)
	err := c.cc.Invoke(ctx, AuthorizationService_GetRoleAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetRoleList(ctx context.Context, in *GetRoleListReq, opts ...grpc.CallOption) (*GetRoleListReply, error) {
	out := new(GetRoleListReply)
	err := c.cc.Invoke(ctx, AuthorizationService_GetRoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, AuthorizationService_CreateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*CheckReply, error) {
	out := new(CheckReply)
	err := c.cc.Invoke(ctx, AuthorizationService_UpdateRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*CheckReply, error) {
	out := new(CheckReply)
	err := c.cc.Invoke(ctx, AuthorizationService_DeleteRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
// All implementations must embed UnimplementedAuthorizationServiceServer
// for forward compatibility
type AuthorizationServiceServer interface {
	CheckAuthorization(context.Context, *CheckAuthorizationReq) (*CheckAuthorizationReply, error)
	// ================== Api ==================
	// Api列表 - 所有
	GetApiListAll(context.Context, *emptypb.Empty) (*GetApiListAllReply, error)
	// Api列表 - 分页
	GetApiList(context.Context, *GetApiListReq) (*GetApiListReply, error)
	// Api创建
	CreateApi(context.Context, *CreateApiReq) (*Api, error)
	// Api更新
	UpdateApi(context.Context, *UpdateApiReq) (*CheckReply, error)
	// Api删除
	DeleteApi(context.Context, *DeleteApiReq) (*CheckReply, error)
	// ================== 角色 ==================
	// 全部角色
	GetRoleAll(context.Context, *emptypb.Empty) (*GetRoleAllReply, error)
	GetRoleList(context.Context, *GetRoleListReq) (*GetRoleListReply, error)
	// 角色创建
	CreateRole(context.Context, *CreateRoleReq) (*Role, error)
	// 角色更新
	UpdateRole(context.Context, *UpdateRoleReq) (*CheckReply, error)
	// 角色删除
	DeleteRole(context.Context, *DeleteRoleReq) (*CheckReply, error)
	mustEmbedUnimplementedAuthorizationServiceServer()
}

// UnimplementedAuthorizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (UnimplementedAuthorizationServiceServer) CheckAuthorization(context.Context, *CheckAuthorizationReq) (*CheckAuthorizationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthorization not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetApiListAll(context.Context, *emptypb.Empty) (*GetApiListAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiListAll not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetApiList(context.Context, *GetApiListReq) (*GetApiListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiList not implemented")
}
func (UnimplementedAuthorizationServiceServer) CreateApi(context.Context, *CreateApiReq) (*Api, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApi not implemented")
}
func (UnimplementedAuthorizationServiceServer) UpdateApi(context.Context, *UpdateApiReq) (*CheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApi not implemented")
}
func (UnimplementedAuthorizationServiceServer) DeleteApi(context.Context, *DeleteApiReq) (*CheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApi not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetRoleAll(context.Context, *emptypb.Empty) (*GetRoleAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleAll not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetRoleList(context.Context, *GetRoleListReq) (*GetRoleListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleList not implemented")
}
func (UnimplementedAuthorizationServiceServer) CreateRole(context.Context, *CreateRoleReq) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedAuthorizationServiceServer) UpdateRole(context.Context, *UpdateRoleReq) (*CheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedAuthorizationServiceServer) DeleteRole(context.Context, *DeleteRoleReq) (*CheckReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedAuthorizationServiceServer) mustEmbedUnimplementedAuthorizationServiceServer() {}

// UnsafeAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServiceServer will
// result in compilation errors.
type UnsafeAuthorizationServiceServer interface {
	mustEmbedUnimplementedAuthorizationServiceServer()
}

func RegisterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv AuthorizationServiceServer) {
	s.RegisterService(&AuthorizationService_ServiceDesc, srv)
}

func _AuthorizationService_CheckAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthorizationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CheckAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_CheckAuthorization_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CheckAuthorization(ctx, req.(*CheckAuthorizationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetApiListAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetApiListAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetApiListAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetApiListAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetApiList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetApiList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetApiList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetApiList(ctx, req.(*GetApiListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CreateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_CreateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateApi(ctx, req.(*CreateApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_UpdateApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).UpdateApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_UpdateApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).UpdateApi(ctx, req.(*UpdateApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_DeleteApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).DeleteApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_DeleteApi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).DeleteApi(ctx, req.(*DeleteApiReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetRoleAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetRoleAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetRoleAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetRoleAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_GetRoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetRoleList(ctx, req.(*GetRoleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_CreateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_UpdateRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthorizationService_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).DeleteRole(ctx, req.(*DeleteRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationService_ServiceDesc is the grpc.ServiceDesc for AuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.v1.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuthorization",
			Handler:    _AuthorizationService_CheckAuthorization_Handler,
		},
		{
			MethodName: "GetApiListAll",
			Handler:    _AuthorizationService_GetApiListAll_Handler,
		},
		{
			MethodName: "GetApiList",
			Handler:    _AuthorizationService_GetApiList_Handler,
		},
		{
			MethodName: "CreateApi",
			Handler:    _AuthorizationService_CreateApi_Handler,
		},
		{
			MethodName: "UpdateApi",
			Handler:    _AuthorizationService_UpdateApi_Handler,
		},
		{
			MethodName: "DeleteApi",
			Handler:    _AuthorizationService_DeleteApi_Handler,
		},
		{
			MethodName: "GetRoleAll",
			Handler:    _AuthorizationService_GetRoleAll_Handler,
		},
		{
			MethodName: "GetRoleList",
			Handler:    _AuthorizationService_GetRoleList_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _AuthorizationService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _AuthorizationService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _AuthorizationService_DeleteRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization/v1/authorization.proto",
}
