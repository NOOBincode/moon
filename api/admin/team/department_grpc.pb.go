// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: admin/team/department.proto

package team

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

const (
	Department_CreateDepartment_FullMethodName   = "/api.admin.team.Department/CreateDepartment"
	Department_UpdateDepartment_FullMethodName   = "/api.admin.team.Department/UpdateDepartment"
	Department_DeleteDepartment_FullMethodName   = "/api.admin.team.Department/DeleteDepartment"
	Department_GetDepartment_FullMethodName      = "/api.admin.team.Department/GetDepartment"
	Department_ListDepartment_FullMethodName     = "/api.admin.team.Department/ListDepartment"
	Department_SetDepartmentAdmin_FullMethodName = "/api.admin.team.Department/SetDepartmentAdmin"
)

// DepartmentClient is the client API for Department service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DepartmentClient interface {
	// 创建部门
	CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*CreateDepartmentReply, error)
	// 更新部门
	UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*UpdateDepartmentReply, error)
	// 删除部门
	DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*DeleteDepartmentReply, error)
	// 获取部门详情
	GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*GetDepartmentReply, error)
	// 获取部门列表（树）
	ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentReply, error)
	// 设置部门管理员
	SetDepartmentAdmin(ctx context.Context, in *SetDepartmentAdminRequest, opts ...grpc.CallOption) (*SetDepartmentAdminReply, error)
}

type departmentClient struct {
	cc grpc.ClientConnInterface
}

func NewDepartmentClient(cc grpc.ClientConnInterface) DepartmentClient {
	return &departmentClient{cc}
}

func (c *departmentClient) CreateDepartment(ctx context.Context, in *CreateDepartmentRequest, opts ...grpc.CallOption) (*CreateDepartmentReply, error) {
	out := new(CreateDepartmentReply)
	err := c.cc.Invoke(ctx, Department_CreateDepartment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentClient) UpdateDepartment(ctx context.Context, in *UpdateDepartmentRequest, opts ...grpc.CallOption) (*UpdateDepartmentReply, error) {
	out := new(UpdateDepartmentReply)
	err := c.cc.Invoke(ctx, Department_UpdateDepartment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentClient) DeleteDepartment(ctx context.Context, in *DeleteDepartmentRequest, opts ...grpc.CallOption) (*DeleteDepartmentReply, error) {
	out := new(DeleteDepartmentReply)
	err := c.cc.Invoke(ctx, Department_DeleteDepartment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentClient) GetDepartment(ctx context.Context, in *GetDepartmentRequest, opts ...grpc.CallOption) (*GetDepartmentReply, error) {
	out := new(GetDepartmentReply)
	err := c.cc.Invoke(ctx, Department_GetDepartment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentClient) ListDepartment(ctx context.Context, in *ListDepartmentRequest, opts ...grpc.CallOption) (*ListDepartmentReply, error) {
	out := new(ListDepartmentReply)
	err := c.cc.Invoke(ctx, Department_ListDepartment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *departmentClient) SetDepartmentAdmin(ctx context.Context, in *SetDepartmentAdminRequest, opts ...grpc.CallOption) (*SetDepartmentAdminReply, error) {
	out := new(SetDepartmentAdminReply)
	err := c.cc.Invoke(ctx, Department_SetDepartmentAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DepartmentServer is the server API for Department service.
// All implementations must embed UnimplementedDepartmentServer
// for forward compatibility
type DepartmentServer interface {
	// 创建部门
	CreateDepartment(context.Context, *CreateDepartmentRequest) (*CreateDepartmentReply, error)
	// 更新部门
	UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*UpdateDepartmentReply, error)
	// 删除部门
	DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*DeleteDepartmentReply, error)
	// 获取部门详情
	GetDepartment(context.Context, *GetDepartmentRequest) (*GetDepartmentReply, error)
	// 获取部门列表（树）
	ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentReply, error)
	// 设置部门管理员
	SetDepartmentAdmin(context.Context, *SetDepartmentAdminRequest) (*SetDepartmentAdminReply, error)
	mustEmbedUnimplementedDepartmentServer()
}

// UnimplementedDepartmentServer must be embedded to have forward compatible implementations.
type UnimplementedDepartmentServer struct {
}

func (UnimplementedDepartmentServer) CreateDepartment(context.Context, *CreateDepartmentRequest) (*CreateDepartmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDepartment not implemented")
}
func (UnimplementedDepartmentServer) UpdateDepartment(context.Context, *UpdateDepartmentRequest) (*UpdateDepartmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDepartment not implemented")
}
func (UnimplementedDepartmentServer) DeleteDepartment(context.Context, *DeleteDepartmentRequest) (*DeleteDepartmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepartment not implemented")
}
func (UnimplementedDepartmentServer) GetDepartment(context.Context, *GetDepartmentRequest) (*GetDepartmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDepartment not implemented")
}
func (UnimplementedDepartmentServer) ListDepartment(context.Context, *ListDepartmentRequest) (*ListDepartmentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDepartment not implemented")
}
func (UnimplementedDepartmentServer) SetDepartmentAdmin(context.Context, *SetDepartmentAdminRequest) (*SetDepartmentAdminReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDepartmentAdmin not implemented")
}
func (UnimplementedDepartmentServer) mustEmbedUnimplementedDepartmentServer() {}

// UnsafeDepartmentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DepartmentServer will
// result in compilation errors.
type UnsafeDepartmentServer interface {
	mustEmbedUnimplementedDepartmentServer()
}

func RegisterDepartmentServer(s grpc.ServiceRegistrar, srv DepartmentServer) {
	s.RegisterService(&Department_ServiceDesc, srv)
}

func _Department_CreateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).CreateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Department_CreateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).CreateDepartment(ctx, req.(*CreateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Department_UpdateDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).UpdateDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Department_UpdateDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).UpdateDepartment(ctx, req.(*UpdateDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Department_DeleteDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).DeleteDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Department_DeleteDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).DeleteDepartment(ctx, req.(*DeleteDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Department_GetDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).GetDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Department_GetDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).GetDepartment(ctx, req.(*GetDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Department_ListDepartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDepartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).ListDepartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Department_ListDepartment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).ListDepartment(ctx, req.(*ListDepartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Department_SetDepartmentAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDepartmentAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DepartmentServer).SetDepartmentAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Department_SetDepartmentAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DepartmentServer).SetDepartmentAdmin(ctx, req.(*SetDepartmentAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Department_ServiceDesc is the grpc.ServiceDesc for Department service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Department_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.team.Department",
	HandlerType: (*DepartmentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDepartment",
			Handler:    _Department_CreateDepartment_Handler,
		},
		{
			MethodName: "UpdateDepartment",
			Handler:    _Department_UpdateDepartment_Handler,
		},
		{
			MethodName: "DeleteDepartment",
			Handler:    _Department_DeleteDepartment_Handler,
		},
		{
			MethodName: "GetDepartment",
			Handler:    _Department_GetDepartment_Handler,
		},
		{
			MethodName: "ListDepartment",
			Handler:    _Department_ListDepartment_Handler,
		},
		{
			MethodName: "SetDepartmentAdmin",
			Handler:    _Department_SetDepartmentAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/team/department.proto",
}
