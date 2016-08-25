// Code generated by protoc-gen-go.
// source: ab.proto
// DO NOT EDIT!

/*
Package atomicbroadcast is a generated protocol buffer package.

It is generated from these files:
	ab.proto

It has these top-level messages:
	BroadcastResponse
	BroadcastMessage
	SeekInfo
	Acknowledgement
	DeliverUpdate
	Block
	DeliverResponse
*/
package atomicbroadcast

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// These status codes are intended to resemble selected HTTP status codes
type Status int32

const (
	Status_SUCCESS             Status = 0
	Status_BAD_REQUEST         Status = 400
	Status_FORBIDDEN           Status = 403
	Status_NOT_FOUND           Status = 404
	Status_SERVICE_UNAVAILABLE Status = 503
)

var Status_name = map[int32]string{
	0:   "SUCCESS",
	400: "BAD_REQUEST",
	403: "FORBIDDEN",
	404: "NOT_FOUND",
	503: "SERVICE_UNAVAILABLE",
}
var Status_value = map[string]int32{
	"SUCCESS":             0,
	"BAD_REQUEST":         400,
	"FORBIDDEN":           403,
	"NOT_FOUND":           404,
	"SERVICE_UNAVAILABLE": 503,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

// Start may be specified to a specific block number, or may be request from the newest or oldest available
// The start location is always inclusive, so the first reply from NEWEST will contain the newest block at the time
// of reception, it will must not wait until a new block is created.  Similarly, when SPECIFIED, and SpecifiedNumber = 10
// The first block received must be block 10, not block 11
type SeekInfo_StartType int32

const (
	SeekInfo_NEWEST    SeekInfo_StartType = 0
	SeekInfo_OLDEST    SeekInfo_StartType = 1
	SeekInfo_SPECIFIED SeekInfo_StartType = 2
)

var SeekInfo_StartType_name = map[int32]string{
	0: "NEWEST",
	1: "OLDEST",
	2: "SPECIFIED",
}
var SeekInfo_StartType_value = map[string]int32{
	"NEWEST":    0,
	"OLDEST":    1,
	"SPECIFIED": 2,
}

func (x SeekInfo_StartType) String() string {
	return proto.EnumName(SeekInfo_StartType_name, int32(x))
}

type BroadcastResponse struct {
	Status Status `protobuf:"varint,1,opt,name=Status,enum=atomicbroadcast.Status" json:"Status,omitempty"`
}

func (m *BroadcastResponse) Reset()         { *m = BroadcastResponse{} }
func (m *BroadcastResponse) String() string { return proto.CompactTextString(m) }
func (*BroadcastResponse) ProtoMessage()    {}

type BroadcastMessage struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *BroadcastMessage) Reset()         { *m = BroadcastMessage{} }
func (m *BroadcastMessage) String() string { return proto.CompactTextString(m) }
func (*BroadcastMessage) ProtoMessage()    {}

type SeekInfo struct {
	Start           SeekInfo_StartType `protobuf:"varint,1,opt,name=Start,enum=atomicbroadcast.SeekInfo_StartType" json:"Start,omitempty"`
	SpecifiedNumber uint64             `protobuf:"varint,2,opt,name=SpecifiedNumber" json:"SpecifiedNumber,omitempty"`
	WindowSize      uint64             `protobuf:"varint,3,opt,name=WindowSize" json:"WindowSize,omitempty"`
}

func (m *SeekInfo) Reset()         { *m = SeekInfo{} }
func (m *SeekInfo) String() string { return proto.CompactTextString(m) }
func (*SeekInfo) ProtoMessage()    {}

type Acknowledgement struct {
	Number uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
}

func (m *Acknowledgement) Reset()         { *m = Acknowledgement{} }
func (m *Acknowledgement) String() string { return proto.CompactTextString(m) }
func (*Acknowledgement) ProtoMessage()    {}

// The update message either causes a seek to a new stream start with a new window, or acknowledges a received block and advances the base of the window
type DeliverUpdate struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverUpdate_Acknowledgement
	//	*DeliverUpdate_Seek
	Type isDeliverUpdate_Type `protobuf_oneof:"Type"`
}

func (m *DeliverUpdate) Reset()         { *m = DeliverUpdate{} }
func (m *DeliverUpdate) String() string { return proto.CompactTextString(m) }
func (*DeliverUpdate) ProtoMessage()    {}

type isDeliverUpdate_Type interface {
	isDeliverUpdate_Type()
}

type DeliverUpdate_Acknowledgement struct {
	Acknowledgement *Acknowledgement `protobuf:"bytes,1,opt,name=Acknowledgement,oneof"`
}
type DeliverUpdate_Seek struct {
	Seek *SeekInfo `protobuf:"bytes,2,opt,name=Seek,oneof"`
}

func (*DeliverUpdate_Acknowledgement) isDeliverUpdate_Type() {}
func (*DeliverUpdate_Seek) isDeliverUpdate_Type()            {}

func (m *DeliverUpdate) GetType() isDeliverUpdate_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverUpdate) GetAcknowledgement() *Acknowledgement {
	if x, ok := m.GetType().(*DeliverUpdate_Acknowledgement); ok {
		return x.Acknowledgement
	}
	return nil
}

func (m *DeliverUpdate) GetSeek() *SeekInfo {
	if x, ok := m.GetType().(*DeliverUpdate_Seek); ok {
		return x.Seek
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeliverUpdate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _DeliverUpdate_OneofMarshaler, _DeliverUpdate_OneofUnmarshaler, []interface{}{
		(*DeliverUpdate_Acknowledgement)(nil),
		(*DeliverUpdate_Seek)(nil),
	}
}

func _DeliverUpdate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeliverUpdate)
	// Type
	switch x := m.Type.(type) {
	case *DeliverUpdate_Acknowledgement:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Acknowledgement); err != nil {
			return err
		}
	case *DeliverUpdate_Seek:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Seek); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeliverUpdate.Type has unexpected type %T", x)
	}
	return nil
}

func _DeliverUpdate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeliverUpdate)
	switch tag {
	case 1: // Type.Acknowledgement
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Acknowledgement)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverUpdate_Acknowledgement{msg}
		return true, err
	case 2: // Type.Seek
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SeekInfo)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverUpdate_Seek{msg}
		return true, err
	default:
		return false, nil
	}
}

// This is a temporary data structure, meant to hold the place of the finalized block structure
// This must be a 'block' structure and not a 'batch' structure, although the terminology is slightly confusing
// The requirement is to allow for a consumer of the orderer to declare the unvalidated blockchain as the definitive
// blockchain, without breaking the hash chain or existing proof
type Block struct {
	Number   uint64              `protobuf:"varint,2,opt,name=Number" json:"Number,omitempty"`
	PrevHash []byte              `protobuf:"bytes,3,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	Proof    []byte              `protobuf:"bytes,4,opt,name=Proof,proto3" json:"Proof,omitempty"`
	Messages []*BroadcastMessage `protobuf:"bytes,5,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}

func (m *Block) GetMessages() []*BroadcastMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

type DeliverResponse struct {
	// Types that are valid to be assigned to Type:
	//	*DeliverResponse_Error
	//	*DeliverResponse_Block
	Type isDeliverResponse_Type `protobuf_oneof:"Type"`
}

func (m *DeliverResponse) Reset()         { *m = DeliverResponse{} }
func (m *DeliverResponse) String() string { return proto.CompactTextString(m) }
func (*DeliverResponse) ProtoMessage()    {}

type isDeliverResponse_Type interface {
	isDeliverResponse_Type()
}

type DeliverResponse_Error struct {
	Error Status `protobuf:"varint,1,opt,name=Error,enum=atomicbroadcast.Status,oneof"`
}
type DeliverResponse_Block struct {
	Block *Block `protobuf:"bytes,2,opt,name=Block,oneof"`
}

func (*DeliverResponse_Error) isDeliverResponse_Type() {}
func (*DeliverResponse_Block) isDeliverResponse_Type() {}

func (m *DeliverResponse) GetType() isDeliverResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DeliverResponse) GetError() Status {
	if x, ok := m.GetType().(*DeliverResponse_Error); ok {
		return x.Error
	}
	return Status_SUCCESS
}

func (m *DeliverResponse) GetBlock() *Block {
	if x, ok := m.GetType().(*DeliverResponse_Block); ok {
		return x.Block
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DeliverResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _DeliverResponse_OneofMarshaler, _DeliverResponse_OneofUnmarshaler, []interface{}{
		(*DeliverResponse_Error)(nil),
		(*DeliverResponse_Block)(nil),
	}
}

func _DeliverResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DeliverResponse)
	// Type
	switch x := m.Type.(type) {
	case *DeliverResponse_Error:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Error))
	case *DeliverResponse_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DeliverResponse.Type has unexpected type %T", x)
	}
	return nil
}

func _DeliverResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DeliverResponse)
	switch tag {
	case 1: // Type.Error
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &DeliverResponse_Error{Status(x)}
		return true, err
	case 2: // Type.Block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Block)
		err := b.DecodeMessage(msg)
		m.Type = &DeliverResponse_Block{msg}
		return true, err
	default:
		return false, nil
	}
}

func init() {
	proto.RegisterEnum("atomicbroadcast.Status", Status_name, Status_value)
	proto.RegisterEnum("atomicbroadcast.SeekInfo_StartType", SeekInfo_StartType_name, SeekInfo_StartType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for AtomicBroadcast service

type AtomicBroadcastClient interface {
	// broadcast receives a reply of Acknowledgement for each BroadcastMessage in order, indicating success or type of failure
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error)
	// deliver first requires an update containing a seek message, then a stream of block replies is received.
	// The receiver may choose to send an Acknowledgement for any block number it receives, however Acknowledgements must never be more than WindowSize apart
	// To avoid latency, clients will likely acknowledge before the WindowSize has been exhausted, preventing the server from stopping and waiting for an Acknowledgement
	Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error)
}

type atomicBroadcastClient struct {
	cc *grpc.ClientConn
}

func NewAtomicBroadcastClient(cc *grpc.ClientConn) AtomicBroadcastClient {
	return &atomicBroadcastClient{cc}
}

func (c *atomicBroadcastClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_BroadcastClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[0], c.cc, "/atomicbroadcast.AtomicBroadcast/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastBroadcastClient{stream}
	return x, nil
}

type AtomicBroadcast_BroadcastClient interface {
	Send(*BroadcastMessage) error
	Recv() (*BroadcastResponse, error)
	grpc.ClientStream
}

type atomicBroadcastBroadcastClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastBroadcastClient) Send(m *BroadcastMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastClient) Recv() (*BroadcastResponse, error) {
	m := new(BroadcastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *atomicBroadcastClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (AtomicBroadcast_DeliverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AtomicBroadcast_serviceDesc.Streams[1], c.cc, "/atomicbroadcast.AtomicBroadcast/Deliver", opts...)
	if err != nil {
		return nil, err
	}
	x := &atomicBroadcastDeliverClient{stream}
	return x, nil
}

type AtomicBroadcast_DeliverClient interface {
	Send(*DeliverUpdate) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type atomicBroadcastDeliverClient struct {
	grpc.ClientStream
}

func (x *atomicBroadcastDeliverClient) Send(m *DeliverUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AtomicBroadcast service

type AtomicBroadcastServer interface {
	// broadcast receives a reply of Acknowledgement for each BroadcastMessage in order, indicating success or type of failure
	Broadcast(AtomicBroadcast_BroadcastServer) error
	// deliver first requires an update containing a seek message, then a stream of block replies is received.
	// The receiver may choose to send an Acknowledgement for any block number it receives, however Acknowledgements must never be more than WindowSize apart
	// To avoid latency, clients will likely acknowledge before the WindowSize has been exhausted, preventing the server from stopping and waiting for an Acknowledgement
	Deliver(AtomicBroadcast_DeliverServer) error
}

func RegisterAtomicBroadcastServer(s *grpc.Server, srv AtomicBroadcastServer) {
	s.RegisterService(&_AtomicBroadcast_serviceDesc, srv)
}

func _AtomicBroadcast_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Broadcast(&atomicBroadcastBroadcastServer{stream})
}

type AtomicBroadcast_BroadcastServer interface {
	Send(*BroadcastResponse) error
	Recv() (*BroadcastMessage, error)
	grpc.ServerStream
}

type atomicBroadcastBroadcastServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastBroadcastServer) Send(m *BroadcastResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastBroadcastServer) Recv() (*BroadcastMessage, error) {
	m := new(BroadcastMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AtomicBroadcast_Deliver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AtomicBroadcastServer).Deliver(&atomicBroadcastDeliverServer{stream})
}

type AtomicBroadcast_DeliverServer interface {
	Send(*DeliverResponse) error
	Recv() (*DeliverUpdate, error)
	grpc.ServerStream
}

type atomicBroadcastDeliverServer struct {
	grpc.ServerStream
}

func (x *atomicBroadcastDeliverServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *atomicBroadcastDeliverServer) Recv() (*DeliverUpdate, error) {
	m := new(DeliverUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AtomicBroadcast_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomicbroadcast.AtomicBroadcast",
	HandlerType: (*AtomicBroadcastServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _AtomicBroadcast_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Deliver",
			Handler:       _AtomicBroadcast_Deliver_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
}
