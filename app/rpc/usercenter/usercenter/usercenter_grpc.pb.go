// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: usercenter.proto

package usercenter

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

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error)
	DelOldCaptcha(ctx context.Context, in *DelOldCaptchaReq, opts ...grpc.CallOption) (*SucResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*SucResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	UpdatePassword(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*SucResp, error)
	DeleteBySelf(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*SucResp, error)
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	IsExitUser(ctx context.Context, in *IsExitUserReq, opts ...grpc.CallOption) (*SucResp, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error) {
	out := new(GetCaptchaResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/getCaptcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DelOldCaptcha(ctx context.Context, in *DelOldCaptchaReq, opts ...grpc.CallOption) (*SucResp, error) {
	out := new(SucResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/delOldCaptcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*SucResp, error) {
	out := new(SucResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/getUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdatePassword(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*SucResp, error) {
	out := new(SucResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/updatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DeleteBySelf(ctx context.Context, in *DelReq, opts ...grpc.CallOption) (*SucResp, error) {
	out := new(SucResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/deleteBySelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/generateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) IsExitUser(ctx context.Context, in *IsExitUserReq, opts ...grpc.CallOption) (*SucResp, error) {
	out := new(SucResp)
	err := c.cc.Invoke(ctx, "/usercenter.usercenter/isExitUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	GetCaptcha(context.Context, *GetCaptchaReq) (*GetCaptchaResp, error)
	DelOldCaptcha(context.Context, *DelOldCaptchaReq) (*SucResp, error)
	Register(context.Context, *RegisterReq) (*SucResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	GetUserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
	UpdatePassword(context.Context, *LoginReq) (*SucResp, error)
	DeleteBySelf(context.Context, *DelReq) (*SucResp, error)
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	IsExitUser(context.Context, *IsExitUserReq) (*SucResp, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) GetCaptcha(context.Context, *GetCaptchaReq) (*GetCaptchaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcha not implemented")
}
func (UnimplementedUsercenterServer) DelOldCaptcha(context.Context, *DelOldCaptchaReq) (*SucResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelOldCaptcha not implemented")
}
func (UnimplementedUsercenterServer) Register(context.Context, *RegisterReq) (*SucResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUsercenterServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsercenterServer) GetUserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUsercenterServer) UpdatePassword(context.Context, *LoginReq) (*SucResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUsercenterServer) DeleteBySelf(context.Context, *DelReq) (*SucResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBySelf not implemented")
}
func (UnimplementedUsercenterServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUsercenterServer) IsExitUser(context.Context, *IsExitUserReq) (*SucResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsExitUser not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_GetCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/getCaptcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetCaptcha(ctx, req.(*GetCaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DelOldCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelOldCaptchaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DelOldCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/delOldCaptcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DelOldCaptcha(ctx, req.(*DelOldCaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/getUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/updatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdatePassword(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DeleteBySelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DeleteBySelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/deleteBySelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DeleteBySelf(ctx, req.(*DelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/generateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_IsExitUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsExitUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).IsExitUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usercenter.usercenter/isExitUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).IsExitUser(ctx, req.(*IsExitUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usercenter.usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCaptcha",
			Handler:    _Usercenter_GetCaptcha_Handler,
		},
		{
			MethodName: "delOldCaptcha",
			Handler:    _Usercenter_DelOldCaptcha_Handler,
		},
		{
			MethodName: "register",
			Handler:    _Usercenter_Register_Handler,
		},
		{
			MethodName: "login",
			Handler:    _Usercenter_Login_Handler,
		},
		{
			MethodName: "getUserInfo",
			Handler:    _Usercenter_GetUserInfo_Handler,
		},
		{
			MethodName: "updatePassword",
			Handler:    _Usercenter_UpdatePassword_Handler,
		},
		{
			MethodName: "deleteBySelf",
			Handler:    _Usercenter_DeleteBySelf_Handler,
		},
		{
			MethodName: "generateToken",
			Handler:    _Usercenter_GenerateToken_Handler,
		},
		{
			MethodName: "isExitUser",
			Handler:    _Usercenter_IsExitUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}
