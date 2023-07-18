// Code generated by goctl. DO NOT EDIT.
// Source: ws.proto

package wsclient

import (
	"context"

	"ccps.com/service/ws/ws"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddTagsReq     = ws.AddTagsReq
	AddTagsResp    = ws.AddTagsResp
	Client         = ws.Client
	ClientsReq     = ws.ClientsReq
	ClientsRes     = ws.ClientsRes
	CountReq       = ws.CountReq
	CountRes       = ws.CountRes
	JoinRoomReq    = ws.JoinRoomReq
	JoinRoomRes    = ws.JoinRoomRes
	LeaveRoomReq   = ws.LeaveRoomReq
	LeaveRoomRes   = ws.LeaveRoomRes
	PingReq        = ws.PingReq
	PingRes        = ws.PingRes
	PushReq        = ws.PushReq
	PushRes        = ws.PushRes
	RemoveTagsReq  = ws.RemoveTagsReq
	RemoveTagsResp = ws.RemoveTagsResp
	SelectOptions  = ws.SelectOptions

	Ws interface {
		Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
		// JoinRoom 加入房间
		JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomRes, error)
		// LeaveRoom 离开房间
		LeaveRoom(ctx context.Context, in *LeaveRoomReq, opts ...grpc.CallOption) (*LeaveRoomRes, error)
		// AddTag 添加标签
		AddTag(ctx context.Context, in *AddTagsReq, opts ...grpc.CallOption) (*AddTagsResp, error)
		// RemoveTag 移除标签
		RemoveTag(ctx context.Context, in *RemoveTagsReq, opts ...grpc.CallOption) (*RemoveTagsResp, error)
		// Count 获取客户端连接数量
		Count(ctx context.Context, in *CountReq, opts ...grpc.CallOption) (*CountRes, error)
		// Push 推送消息
		Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushRes, error)
		// Clients 查询客户端
		Clients(ctx context.Context, in *ClientsReq, opts ...grpc.CallOption) (*ClientsRes, error)
	}

	defaultWs struct {
		cli zrpc.Client
	}
)

func NewWs(cli zrpc.Client) Ws {
	return &defaultWs{
		cli: cli,
	}
}

func (m *defaultWs) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

// JoinRoom 加入房间
func (m *defaultWs) JoinRoom(ctx context.Context, in *JoinRoomReq, opts ...grpc.CallOption) (*JoinRoomRes, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.JoinRoom(ctx, in, opts...)
}

// LeaveRoom 离开房间
func (m *defaultWs) LeaveRoom(ctx context.Context, in *LeaveRoomReq, opts ...grpc.CallOption) (*LeaveRoomRes, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.LeaveRoom(ctx, in, opts...)
}

// AddTag 添加标签
func (m *defaultWs) AddTag(ctx context.Context, in *AddTagsReq, opts ...grpc.CallOption) (*AddTagsResp, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.AddTag(ctx, in, opts...)
}

// RemoveTag 移除标签
func (m *defaultWs) RemoveTag(ctx context.Context, in *RemoveTagsReq, opts ...grpc.CallOption) (*RemoveTagsResp, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.RemoveTag(ctx, in, opts...)
}

// Count 获取客户端连接数量
func (m *defaultWs) Count(ctx context.Context, in *CountReq, opts ...grpc.CallOption) (*CountRes, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.Count(ctx, in, opts...)
}

// Push 推送消息
func (m *defaultWs) Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushRes, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.Push(ctx, in, opts...)
}

// Clients 查询客户端
func (m *defaultWs) Clients(ctx context.Context, in *ClientsReq, opts ...grpc.CallOption) (*ClientsRes, error) {
	client := ws.NewWsClient(m.cli.Conn())
	return client.Clients(ctx, in, opts...)
}