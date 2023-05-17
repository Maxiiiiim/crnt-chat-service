// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/chat/chat.proto

package chat

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
	ChatService_CreateUser_FullMethodName          = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/CreateUser"
	ChatService_GetUserById_FullMethodName         = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/GetUserById"
	ChatService_UpdateUserMeta_FullMethodName      = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/UpdateUserMeta"
	ChatService_SetLastActive_FullMethodName       = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/SetLastActive"
	ChatService_GetUsersLastActive_FullMethodName  = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/GetUsersLastActive"
	ChatService_CreateDialog_FullMethodName        = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/CreateDialog"
	ChatService_GetDialogById_FullMethodName       = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/GetDialogById"
	ChatService_UpdateDialogMeta_FullMethodName    = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/UpdateDialogMeta"
	ChatService_GetUserDialogs_FullMethodName      = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/GetUserDialogs"
	ChatService_JoinDialog_FullMethodName          = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/JoinDialog"
	ChatService_LeaveDialog_FullMethodName         = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/LeaveDialog"
	ChatService_CountUnreadMessages_FullMethodName = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/CountUnreadMessages"
	ChatService_SendMessage_FullMethodName         = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/SendMessage"
	ChatService_GetDialogMessages_FullMethodName   = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/GetDialogMessages"
	ChatService_SearchMessages_FullMethodName      = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/SearchMessages"
	ChatService_DeleteMessage_FullMethodName       = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/DeleteMessage"
	ChatService_UpdateMessage_FullMethodName       = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/UpdateMessage"
	ChatService_SendReply_FullMethodName           = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/SendReply"
	ChatService_GetReplies_FullMethodName          = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/GetReplies"
	ChatService_Ping_FullMethodName                = "/github.maxiiiiim.crnt_chat_service.api.chat.ChatService/Ping"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	// users
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
	UpdateUserMeta(ctx context.Context, in *UpdateUserMetaRequest, opts ...grpc.CallOption) (*UpdateUserMetaResponse, error)
	SetLastActive(ctx context.Context, in *SetLastActiveRequest, opts ...grpc.CallOption) (*SetLastActiveResponse, error)
	GetUsersLastActive(ctx context.Context, in *GetUsersLastActiveRequest, opts ...grpc.CallOption) (*GetUsersLastActiveResponse, error)
	// dialogs
	CreateDialog(ctx context.Context, in *CreateDialogRequest, opts ...grpc.CallOption) (*CreateDialogResponse, error)
	GetDialogById(ctx context.Context, in *GetDialogByIdRequest, opts ...grpc.CallOption) (*GetDialogByIdResponse, error)
	UpdateDialogMeta(ctx context.Context, in *UpdateDialogMetaRequest, opts ...grpc.CallOption) (*UpdateDialogMetaResponse, error)
	GetUserDialogs(ctx context.Context, in *GetUserDialogsRequest, opts ...grpc.CallOption) (*GetUserDialogsResponse, error)
	JoinDialog(ctx context.Context, in *JoinDialogRequest, opts ...grpc.CallOption) (*JoinDialogResponse, error)
	LeaveDialog(ctx context.Context, in *LeaveDialogRequest, opts ...grpc.CallOption) (*LeaveDialogResponse, error)
	CountUnreadMessages(ctx context.Context, in *CountUnreadMessagesRequest, opts ...grpc.CallOption) (*CountUnreadMessagesResponse, error)
	// messages
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	GetDialogMessages(ctx context.Context, in *GetDialogMessagesRequest, opts ...grpc.CallOption) (*GetDialogMessagesResponse, error)
	SearchMessages(ctx context.Context, in *SearchMessagesRequest, opts ...grpc.CallOption) (*SearchMessagesResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResponse, error)
	SendReply(ctx context.Context, in *SendReplyRequest, opts ...grpc.CallOption) (*SendReplyResponse, error)
	GetReplies(ctx context.Context, in *GetRepliesRequest, opts ...grpc.CallOption) (*GetRepliesResponse, error)
	// ping
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, ChatService_GetUserById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateUserMeta(ctx context.Context, in *UpdateUserMetaRequest, opts ...grpc.CallOption) (*UpdateUserMetaResponse, error) {
	out := new(UpdateUserMetaResponse)
	err := c.cc.Invoke(ctx, ChatService_UpdateUserMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SetLastActive(ctx context.Context, in *SetLastActiveRequest, opts ...grpc.CallOption) (*SetLastActiveResponse, error) {
	out := new(SetLastActiveResponse)
	err := c.cc.Invoke(ctx, ChatService_SetLastActive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetUsersLastActive(ctx context.Context, in *GetUsersLastActiveRequest, opts ...grpc.CallOption) (*GetUsersLastActiveResponse, error) {
	out := new(GetUsersLastActiveResponse)
	err := c.cc.Invoke(ctx, ChatService_GetUsersLastActive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateDialog(ctx context.Context, in *CreateDialogRequest, opts ...grpc.CallOption) (*CreateDialogResponse, error) {
	out := new(CreateDialogResponse)
	err := c.cc.Invoke(ctx, ChatService_CreateDialog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetDialogById(ctx context.Context, in *GetDialogByIdRequest, opts ...grpc.CallOption) (*GetDialogByIdResponse, error) {
	out := new(GetDialogByIdResponse)
	err := c.cc.Invoke(ctx, ChatService_GetDialogById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateDialogMeta(ctx context.Context, in *UpdateDialogMetaRequest, opts ...grpc.CallOption) (*UpdateDialogMetaResponse, error) {
	out := new(UpdateDialogMetaResponse)
	err := c.cc.Invoke(ctx, ChatService_UpdateDialogMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetUserDialogs(ctx context.Context, in *GetUserDialogsRequest, opts ...grpc.CallOption) (*GetUserDialogsResponse, error) {
	out := new(GetUserDialogsResponse)
	err := c.cc.Invoke(ctx, ChatService_GetUserDialogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinDialog(ctx context.Context, in *JoinDialogRequest, opts ...grpc.CallOption) (*JoinDialogResponse, error) {
	out := new(JoinDialogResponse)
	err := c.cc.Invoke(ctx, ChatService_JoinDialog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) LeaveDialog(ctx context.Context, in *LeaveDialogRequest, opts ...grpc.CallOption) (*LeaveDialogResponse, error) {
	out := new(LeaveDialogResponse)
	err := c.cc.Invoke(ctx, ChatService_LeaveDialog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CountUnreadMessages(ctx context.Context, in *CountUnreadMessagesRequest, opts ...grpc.CallOption) (*CountUnreadMessagesResponse, error) {
	out := new(CountUnreadMessagesResponse)
	err := c.cc.Invoke(ctx, ChatService_CountUnreadMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_SendMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetDialogMessages(ctx context.Context, in *GetDialogMessagesRequest, opts ...grpc.CallOption) (*GetDialogMessagesResponse, error) {
	out := new(GetDialogMessagesResponse)
	err := c.cc.Invoke(ctx, ChatService_GetDialogMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SearchMessages(ctx context.Context, in *SearchMessagesRequest, opts ...grpc.CallOption) (*SearchMessagesResponse, error) {
	out := new(SearchMessagesResponse)
	err := c.cc.Invoke(ctx, ChatService_SearchMessages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...grpc.CallOption) (*DeleteMessageResponse, error) {
	out := new(DeleteMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_DeleteMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*UpdateMessageResponse, error) {
	out := new(UpdateMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_UpdateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendReply(ctx context.Context, in *SendReplyRequest, opts ...grpc.CallOption) (*SendReplyResponse, error) {
	out := new(SendReplyResponse)
	err := c.cc.Invoke(ctx, ChatService_SendReply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetReplies(ctx context.Context, in *GetRepliesRequest, opts ...grpc.CallOption) (*GetRepliesResponse, error) {
	out := new(GetRepliesResponse)
	err := c.cc.Invoke(ctx, ChatService_GetReplies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, ChatService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations should embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	// users
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	UpdateUserMeta(context.Context, *UpdateUserMetaRequest) (*UpdateUserMetaResponse, error)
	SetLastActive(context.Context, *SetLastActiveRequest) (*SetLastActiveResponse, error)
	GetUsersLastActive(context.Context, *GetUsersLastActiveRequest) (*GetUsersLastActiveResponse, error)
	// dialogs
	CreateDialog(context.Context, *CreateDialogRequest) (*CreateDialogResponse, error)
	GetDialogById(context.Context, *GetDialogByIdRequest) (*GetDialogByIdResponse, error)
	UpdateDialogMeta(context.Context, *UpdateDialogMetaRequest) (*UpdateDialogMetaResponse, error)
	GetUserDialogs(context.Context, *GetUserDialogsRequest) (*GetUserDialogsResponse, error)
	JoinDialog(context.Context, *JoinDialogRequest) (*JoinDialogResponse, error)
	LeaveDialog(context.Context, *LeaveDialogRequest) (*LeaveDialogResponse, error)
	CountUnreadMessages(context.Context, *CountUnreadMessagesRequest) (*CountUnreadMessagesResponse, error)
	// messages
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	GetDialogMessages(context.Context, *GetDialogMessagesRequest) (*GetDialogMessagesResponse, error)
	SearchMessages(context.Context, *SearchMessagesRequest) (*SearchMessagesResponse, error)
	DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResponse, error)
	SendReply(context.Context, *SendReplyRequest) (*SendReplyResponse, error)
	GetReplies(context.Context, *GetRepliesRequest) (*GetRepliesResponse, error)
	// ping
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedChatServiceServer should be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedChatServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedChatServiceServer) UpdateUserMeta(context.Context, *UpdateUserMetaRequest) (*UpdateUserMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserMeta not implemented")
}
func (UnimplementedChatServiceServer) SetLastActive(context.Context, *SetLastActiveRequest) (*SetLastActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLastActive not implemented")
}
func (UnimplementedChatServiceServer) GetUsersLastActive(context.Context, *GetUsersLastActiveRequest) (*GetUsersLastActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersLastActive not implemented")
}
func (UnimplementedChatServiceServer) CreateDialog(context.Context, *CreateDialogRequest) (*CreateDialogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDialog not implemented")
}
func (UnimplementedChatServiceServer) GetDialogById(context.Context, *GetDialogByIdRequest) (*GetDialogByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDialogById not implemented")
}
func (UnimplementedChatServiceServer) UpdateDialogMeta(context.Context, *UpdateDialogMetaRequest) (*UpdateDialogMetaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDialogMeta not implemented")
}
func (UnimplementedChatServiceServer) GetUserDialogs(context.Context, *GetUserDialogsRequest) (*GetUserDialogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDialogs not implemented")
}
func (UnimplementedChatServiceServer) JoinDialog(context.Context, *JoinDialogRequest) (*JoinDialogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinDialog not implemented")
}
func (UnimplementedChatServiceServer) LeaveDialog(context.Context, *LeaveDialogRequest) (*LeaveDialogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveDialog not implemented")
}
func (UnimplementedChatServiceServer) CountUnreadMessages(context.Context, *CountUnreadMessagesRequest) (*CountUnreadMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUnreadMessages not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) GetDialogMessages(context.Context, *GetDialogMessagesRequest) (*GetDialogMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDialogMessages not implemented")
}
func (UnimplementedChatServiceServer) SearchMessages(context.Context, *SearchMessagesRequest) (*SearchMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMessages not implemented")
}
func (UnimplementedChatServiceServer) DeleteMessage(context.Context, *DeleteMessageRequest) (*DeleteMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedChatServiceServer) UpdateMessage(context.Context, *UpdateMessageRequest) (*UpdateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedChatServiceServer) SendReply(context.Context, *SendReplyRequest) (*SendReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReply not implemented")
}
func (UnimplementedChatServiceServer) GetReplies(context.Context, *GetRepliesRequest) (*GetRepliesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplies not implemented")
}
func (UnimplementedChatServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateUserMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateUserMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateUserMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateUserMeta(ctx, req.(*UpdateUserMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SetLastActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLastActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SetLastActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SetLastActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SetLastActive(ctx, req.(*SetLastActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetUsersLastActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersLastActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetUsersLastActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetUsersLastActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetUsersLastActive(ctx, req.(*GetUsersLastActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateDialog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDialogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateDialog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CreateDialog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateDialog(ctx, req.(*CreateDialogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetDialogById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDialogByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetDialogById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetDialogById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetDialogById(ctx, req.(*GetDialogByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateDialogMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDialogMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateDialogMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateDialogMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateDialogMeta(ctx, req.(*UpdateDialogMetaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetUserDialogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDialogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetUserDialogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetUserDialogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetUserDialogs(ctx, req.(*GetUserDialogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_JoinDialog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinDialogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).JoinDialog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_JoinDialog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).JoinDialog(ctx, req.(*JoinDialogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_LeaveDialog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveDialogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LeaveDialog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_LeaveDialog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LeaveDialog(ctx, req.(*LeaveDialogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CountUnreadMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountUnreadMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CountUnreadMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_CountUnreadMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CountUnreadMessages(ctx, req.(*CountUnreadMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetDialogMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDialogMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetDialogMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetDialogMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetDialogMessages(ctx, req.(*GetDialogMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SearchMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SearchMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SearchMessages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SearchMessages(ctx, req.(*SearchMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_DeleteMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeleteMessage(ctx, req.(*DeleteMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_UpdateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_SendReply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendReply(ctx, req.(*SendReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetReplies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRepliesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetReplies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetReplies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetReplies(ctx, req.(*GetRepliesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.maxiiiiim.crnt_chat_service.api.chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _ChatService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _ChatService_GetUserById_Handler,
		},
		{
			MethodName: "UpdateUserMeta",
			Handler:    _ChatService_UpdateUserMeta_Handler,
		},
		{
			MethodName: "SetLastActive",
			Handler:    _ChatService_SetLastActive_Handler,
		},
		{
			MethodName: "GetUsersLastActive",
			Handler:    _ChatService_GetUsersLastActive_Handler,
		},
		{
			MethodName: "CreateDialog",
			Handler:    _ChatService_CreateDialog_Handler,
		},
		{
			MethodName: "GetDialogById",
			Handler:    _ChatService_GetDialogById_Handler,
		},
		{
			MethodName: "UpdateDialogMeta",
			Handler:    _ChatService_UpdateDialogMeta_Handler,
		},
		{
			MethodName: "GetUserDialogs",
			Handler:    _ChatService_GetUserDialogs_Handler,
		},
		{
			MethodName: "JoinDialog",
			Handler:    _ChatService_JoinDialog_Handler,
		},
		{
			MethodName: "LeaveDialog",
			Handler:    _ChatService_LeaveDialog_Handler,
		},
		{
			MethodName: "CountUnreadMessages",
			Handler:    _ChatService_CountUnreadMessages_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
		{
			MethodName: "GetDialogMessages",
			Handler:    _ChatService_GetDialogMessages_Handler,
		},
		{
			MethodName: "SearchMessages",
			Handler:    _ChatService_SearchMessages_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _ChatService_DeleteMessage_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _ChatService_UpdateMessage_Handler,
		},
		{
			MethodName: "SendReply",
			Handler:    _ChatService_SendReply_Handler,
		},
		{
			MethodName: "GetReplies",
			Handler:    _ChatService_GetReplies_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _ChatService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/chat/chat.proto",
}