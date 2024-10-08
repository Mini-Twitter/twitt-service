// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserService_Create_FullMethodName                    = "/user.UserService/Create"
	UserService_GetProfile_FullMethodName                = "/user.UserService/GetProfile"
	UserService_UpdateProfile_FullMethodName             = "/user.UserService/UpdateProfile"
	UserService_ChangePassword_FullMethodName            = "/user.UserService/ChangePassword"
	UserService_ChangeProfileImage_FullMethodName        = "/user.UserService/ChangeProfileImage"
	UserService_FetchUsers_FullMethodName                = "/user.UserService/FetchUsers"
	UserService_ListOfFollowing_FullMethodName           = "/user.UserService/ListOfFollowing"
	UserService_ListOfFollowers_FullMethodName           = "/user.UserService/ListOfFollowers"
	UserService_ListOfFollowingByUsername_FullMethodName = "/user.UserService/ListOfFollowingByUsername"
	UserService_ListOfFollowersByUsername_FullMethodName = "/user.UserService/ListOfFollowersByUsername"
	UserService_DeleteUser_FullMethodName                = "/user.UserService/DeleteUser"
	UserService_Follow_FullMethodName                    = "/user.UserService/Follow"
	UserService_Unfollow_FullMethodName                  = "/user.UserService/Unfollow"
	UserService_GetUserFollowers_FullMethodName          = "/user.UserService/GetUserFollowers"
	UserService_GetUserFollows_FullMethodName            = "/user.UserService/GetUserFollows"
	UserService_MostPopularUser_FullMethodName           = "/user.UserService/MostPopularUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AuthService for user authentication and profile management
type UserServiceClient interface {
	// users
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*GetProfileResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	ChangeProfileImage(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Void, error)
	FetchUsers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*UserResponses, error)
	ListOfFollowing(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error)
	ListOfFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error)
	ListOfFollowingByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error)
	ListOfFollowersByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error)
	DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error)
	// subscribe
	Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowRes, error)
	Unfollow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*DFollowRes, error)
	GetUserFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error)
	GetUserFollows(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error)
	MostPopularUser(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *Id, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, UserService_GetProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, UserService_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeProfileImage(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_ChangeProfileImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FetchUsers(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*UserResponses, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponses)
	err := c.cc.Invoke(ctx, UserService_FetchUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowing(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Followings)
	err := c.cc.Invoke(ctx, UserService_ListOfFollowing_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Followers)
	err := c.cc.Invoke(ctx, UserService_ListOfFollowers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowingByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followings, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Followings)
	err := c.cc.Invoke(ctx, UserService_ListOfFollowingByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListOfFollowersByUsername(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Followers, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Followers)
	err := c.cc.Invoke(ctx, UserService_ListOfFollowersByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FollowRes)
	err := c.cc.Invoke(ctx, UserService_Follow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Unfollow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*DFollowRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DFollowRes)
	err := c.cc.Invoke(ctx, UserService_Unfollow_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFollowers(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Count)
	err := c.cc.Invoke(ctx, UserService_GetUserFollowers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFollows(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Count, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Count)
	err := c.cc.Invoke(ctx, UserService_GetUserFollows_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) MostPopularUser(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_MostPopularUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
//
// AuthService for user authentication and profile management
type UserServiceServer interface {
	// users
	Create(context.Context, *CreateRequest) (*UserResponse, error)
	GetProfile(context.Context, *Id) (*GetProfileResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UserResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	ChangeProfileImage(context.Context, *URL) (*Void, error)
	FetchUsers(context.Context, *Filter) (*UserResponses, error)
	ListOfFollowing(context.Context, *Id) (*Followings, error)
	ListOfFollowers(context.Context, *Id) (*Followers, error)
	ListOfFollowingByUsername(context.Context, *Id) (*Followings, error)
	ListOfFollowersByUsername(context.Context, *Id) (*Followers, error)
	DeleteUser(context.Context, *Id) (*Void, error)
	// subscribe
	Follow(context.Context, *FollowReq) (*FollowRes, error)
	Unfollow(context.Context, *FollowReq) (*DFollowRes, error)
	GetUserFollowers(context.Context, *Id) (*Count, error)
	GetUserFollows(context.Context, *Id) (*Count, error)
	MostPopularUser(context.Context, *Void) (*UserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Create(context.Context, *CreateRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) GetProfile(context.Context, *Id) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) ChangeProfileImage(context.Context, *URL) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeProfileImage not implemented")
}
func (UnimplementedUserServiceServer) FetchUsers(context.Context, *Filter) (*UserResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUsers not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowing(context.Context, *Id) (*Followings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowing not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowers(context.Context, *Id) (*Followers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowers not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowingByUsername(context.Context, *Id) (*Followings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowingByUsername not implemented")
}
func (UnimplementedUserServiceServer) ListOfFollowersByUsername(context.Context, *Id) (*Followers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOfFollowersByUsername not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *Id) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) Follow(context.Context, *FollowReq) (*FollowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedUserServiceServer) Unfollow(context.Context, *FollowReq) (*DFollowRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedUserServiceServer) GetUserFollowers(context.Context, *Id) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowers not implemented")
}
func (UnimplementedUserServiceServer) GetUserFollows(context.Context, *Id) (*Count, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollows not implemented")
}
func (UnimplementedUserServiceServer) MostPopularUser(context.Context, *Void) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostPopularUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeProfileImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeProfileImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangeProfileImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeProfileImage(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FetchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FetchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FetchUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FetchUsers(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListOfFollowing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowing(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListOfFollowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowers(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowingByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowingByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListOfFollowingByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowingByUsername(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListOfFollowersByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListOfFollowersByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListOfFollowersByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListOfFollowersByUsername(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Follow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Follow(ctx, req.(*FollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Unfollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Unfollow(ctx, req.(*FollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserFollowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFollowers(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFollows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFollows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserFollows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFollows(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_MostPopularUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).MostPopularUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_MostPopularUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).MostPopularUser(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserService_UpdateProfile_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "ChangeProfileImage",
			Handler:    _UserService_ChangeProfileImage_Handler,
		},
		{
			MethodName: "FetchUsers",
			Handler:    _UserService_FetchUsers_Handler,
		},
		{
			MethodName: "ListOfFollowing",
			Handler:    _UserService_ListOfFollowing_Handler,
		},
		{
			MethodName: "ListOfFollowers",
			Handler:    _UserService_ListOfFollowers_Handler,
		},
		{
			MethodName: "ListOfFollowingByUsername",
			Handler:    _UserService_ListOfFollowingByUsername_Handler,
		},
		{
			MethodName: "ListOfFollowersByUsername",
			Handler:    _UserService_ListOfFollowersByUsername_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _UserService_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _UserService_Unfollow_Handler,
		},
		{
			MethodName: "GetUserFollowers",
			Handler:    _UserService_GetUserFollowers_Handler,
		},
		{
			MethodName: "GetUserFollows",
			Handler:    _UserService_GetUserFollows_Handler,
		},
		{
			MethodName: "MostPopularUser",
			Handler:    _UserService_MostPopularUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
