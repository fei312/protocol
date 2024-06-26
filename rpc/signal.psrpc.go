// Code generated by protoc-gen-psrpc v0.5.1, DO NOT EDIT.
// source: rpc/signal.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)

var _ = version.PsrpcVersion_0_5

// =======================
// Signal Client Interface
// =======================

type SignalClient[NodeIdTopicType ~string] interface {
	RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error)
}

// ===========================
// Signal ServerImpl Interface
// ===========================

type SignalServerImpl interface {
	RelaySignal(psrpc.ServerStream[*RelaySignalResponse, *RelaySignalRequest]) error
}

// =======================
// Signal Server Interface
// =======================

type SignalServer[NodeIdTopicType ~string] interface {
	RegisterRelaySignalTopic(nodeId NodeIdTopicType) error
	DeregisterRelaySignalTopic(nodeId NodeIdTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// Signal Client
// =============

type signalClient[NodeIdTopicType ~string] struct {
	client *client.RPCClient
}

// NewSignalClient creates a psrpc client that implements the SignalClient interface.
func NewSignalClient[NodeIdTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (SignalClient[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("RelaySignal", false, false, false, true)

	rpcClient, err := client.NewRPCClientWithStreams(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &signalClient[NodeIdTopicType]{
		client: rpcClient,
	}, nil
}

func (c *signalClient[NodeIdTopicType]) RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error) {
	return client.OpenStream[*RelaySignalRequest, *RelaySignalResponse](ctx, c.client, "RelaySignal", []string{string(nodeId)}, opts...)
}

// =============
// Signal Server
// =============

type signalServer[NodeIdTopicType ~string] struct {
	svc SignalServerImpl
	rpc *server.RPCServer
}

// NewSignalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewSignalServer[NodeIdTopicType ~string](svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (SignalServer[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("RelaySignal", false, false, false, true)
	return &signalServer[NodeIdTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *signalServer[NodeIdTopicType]) RegisterRelaySignalTopic(nodeId NodeIdTopicType) error {
	return server.RegisterStreamHandler(s.rpc, "RelaySignal", []string{string(nodeId)}, s.svc.RelaySignal, nil)
}

func (s *signalServer[NodeIdTopicType]) DeregisterRelaySignalTopic(nodeId NodeIdTopicType) {
	s.rpc.DeregisterHandler("RelaySignal", []string{string(nodeId)})
}

func (s *signalServer[NodeIdTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *signalServer[NodeIdTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor7 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x6f, 0xe2, 0x30,
	0x10, 0xc5, 0x84, 0x4f, 0xb3, 0x48, 0x59, 0xb3, 0x0b, 0x51, 0x2e, 0x9b, 0xe5, 0x94, 0x43, 0x95,
	0x48, 0xa9, 0x7a, 0xeb, 0x89, 0x43, 0xc5, 0xd9, 0x51, 0x0f, 0xed, 0x05, 0x05, 0x63, 0x51, 0xab,
	0xc1, 0x36, 0x1e, 0x53, 0x89, 0x9f, 0xd0, 0xbf, 0xd3, 0x3f, 0xd5, 0xbf, 0x51, 0x11, 0xf3, 0x91,
	0x8a, 0xde, 0x66, 0xde, 0x7b, 0x33, 0xef, 0xd9, 0x83, 0x7d, 0xa3, 0x59, 0x0a, 0x62, 0x2d, 0x8b,
	0x32, 0xd1, 0x46, 0x59, 0x45, 0x3c, 0xa3, 0x59, 0x38, 0x54, 0xda, 0x0a, 0x25, 0xc1, 0x61, 0xe1,
	0xb8, 0x14, 0x6f, 0xfc, 0x55, 0xd8, 0x85, 0x90, 0x96, 0x9b, 0xb3, 0x36, 0xfc, 0x7d, 0xc2, 0x8d,
	0x65, 0x0e, 0x9a, 0x7e, 0x22, 0x4c, 0x28, 0x2f, 0x8b, 0x7d, 0x5e, 0x2d, 0xa5, 0x7c, 0xbb, 0xe3,
	0x60, 0xc9, 0x3d, 0x1e, 0x82, 0x2d, 0x8c, 0x5d, 0x00, 0x07, 0x10, 0x4a, 0x06, 0x28, 0x42, 0xf1,
	0x20, 0xfb, 0x9b, 0x1c, 0x37, 0x24, 0xf9, 0x81, 0xcd, 0x1d, 0x39, 0x6f, 0xd0, 0x5f, 0x50, 0xeb,
	0xc9, 0x0d, 0xee, 0xe6, 0x0f, 0x8f, 0xf9, 0x5e, 0xb2, 0xa0, 0x59, 0xcd, 0xf9, 0x97, 0x39, 0x87,
	0xcf, 0x1b, 0xf4, 0x24, 0x21, 0x19, 0xee, 0x19, 0x67, 0x0b, 0x81, 0x17, 0x79, 0xf1, 0x20, 0x1b,
	0x5f, 0xe4, 0xf5, 0x54, 0xf4, 0xac, 0x23, 0x3e, 0xf6, 0x80, 0x6f, 0x83, 0x56, 0x84, 0xe2, 0x16,
	0x3d, 0x94, 0xe4, 0x0f, 0x6e, 0xb3, 0x52, 0x01, 0x0f, 0xda, 0x11, 0x8a, 0x7b, 0xd4, 0x35, 0xb3,
	0x3e, 0xee, 0x6e, 0x38, 0x40, 0xb1, 0xe6, 0x53, 0x8b, 0x47, 0xdf, 0x1e, 0x0a, 0x5a, 0x49, 0xe0,
	0xe4, 0x0e, 0xf7, 0xcd, 0xb1, 0x86, 0xa0, 0x59, 0xd9, 0x4f, 0xae, 0xec, 0x1d, 0x4f, 0x2f, 0xca,
	0x53, 0x00, 0xef, 0x87, 0x00, 0xad, 0x5a, 0x80, 0x8c, 0xe1, 0x8e, 0x5b, 0x42, 0x9e, 0xf0, 0xa0,
	0xe6, 0x4f, 0x26, 0x89, 0xd1, 0x2c, 0xb9, 0xfe, 0xfa, 0x30, 0xb8, 0x26, 0x9c, 0xe9, 0x74, 0xf2,
	0xf1, 0x8e, 0x46, 0x3e, 0x0a, 0x87, 0xa4, 0x2b, 0xd5, 0x8a, 0x2f, 0xc4, 0xea, 0x70, 0x96, 0x08,
	0xcd, 0xfe, 0x3f, 0xff, 0x5b, 0x0b, 0xfb, 0xb2, 0x5b, 0x26, 0x4c, 0x6d, 0xd2, 0x63, 0xf8, 0xb4,
	0x3a, 0x30, 0x53, 0x65, 0x6a, 0x34, 0x5b, 0x76, 0xaa, 0xee, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x61, 0x6a, 0xda, 0x00, 0x41, 0x02, 0x00, 0x00,
}
