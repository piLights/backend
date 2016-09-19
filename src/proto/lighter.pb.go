// Code generated by protoc-gen-go.
// source: src/proto/lighter.proto
// DO NOT EDIT!

/*
Package LighterGRPC is a generated protocol buffer package.

It is generated from these files:
	src/proto/lighter.proto

It has these top-level messages:
	ColorMessage
	StateMessage
	Confirmation
	InitMessage
	LoadConfigRequest
	ServerConfig
	ChangeParameterMessage
	FadeTime
	IPVersion
	Pins
	ScheduledSwitch
*/
package LighterGRPC

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IPVersion_Version int32

const (
	IPVersion_DUAL     IPVersion_Version = 0
	IPVersion_IPV4ONLY IPVersion_Version = 1
	IPVersion_IPV6ONLY IPVersion_Version = 2
)

var IPVersion_Version_name = map[int32]string{
	0: "DUAL",
	1: "IPV4ONLY",
	2: "IPV6ONLY",
}
var IPVersion_Version_value = map[string]int32{
	"DUAL":     0,
	"IPV4ONLY": 1,
	"IPV6ONLY": 2,
}

func (x IPVersion_Version) String() string {
	return proto.EnumName(IPVersion_Version_name, int32(x))
}
func (IPVersion_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type ColorMessage struct {
	Onstate  bool   `protobuf:"varint,1,opt,name=onstate" json:"onstate,omitempty"`
	R        int32  `protobuf:"varint,2,opt,name=r" json:"r,omitempty"`
	G        int32  `protobuf:"varint,3,opt,name=g" json:"g,omitempty"`
	B        int32  `protobuf:"varint,4,opt,name=b" json:"b,omitempty"`
	Opacity  int32  `protobuf:"varint,5,opt,name=opacity" json:"opacity,omitempty"`
	DeviceID string `protobuf:"bytes,6,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,7,opt,name=password" json:"password,omitempty"`
}

func (m *ColorMessage) Reset()                    { *m = ColorMessage{} }
func (m *ColorMessage) String() string            { return proto.CompactTextString(m) }
func (*ColorMessage) ProtoMessage()               {}
func (*ColorMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StateMessage struct {
	Onstate  bool   `protobuf:"varint,1,opt,name=onstate" json:"onstate,omitempty"`
	DeviceID string `protobuf:"bytes,2,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *StateMessage) Reset()                    { *m = StateMessage{} }
func (m *StateMessage) String() string            { return proto.CompactTextString(m) }
func (*StateMessage) ProtoMessage()               {}
func (*StateMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Confirmation struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *Confirmation) Reset()                    { *m = Confirmation{} }
func (m *Confirmation) String() string            { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()               {}
func (*Confirmation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type InitMessage struct {
	DeviceID string `protobuf:"bytes,1,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *InitMessage) Reset()                    { *m = InitMessage{} }
func (m *InitMessage) String() string            { return proto.CompactTextString(m) }
func (*InitMessage) ProtoMessage()               {}
func (*InitMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type LoadConfigRequest struct {
	DeviceID string `protobuf:"bytes,1,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoadConfigRequest) Reset()                    { *m = LoadConfigRequest{} }
func (m *LoadConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadConfigRequest) ProtoMessage()               {}
func (*LoadConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ServerConfig struct {
	ServerName string     `protobuf:"bytes,1,opt,name=serverName" json:"serverName,omitempty"`
	FadeTime   *FadeTime  `protobuf:"bytes,2,opt,name=fadeTime" json:"fadeTime,omitempty"`
	IpVersion  *IPVersion `protobuf:"bytes,3,opt,name=ipVersion" json:"ipVersion,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServerConfig) GetFadeTime() *FadeTime {
	if m != nil {
		return m.FadeTime
	}
	return nil
}

func (m *ServerConfig) GetIpVersion() *IPVersion {
	if m != nil {
		return m.IpVersion
	}
	return nil
}

type ChangeParameterMessage struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	// Types that are valid to be assigned to Parameter:
	//	*ChangeParameterMessage_ServerName
	//	*ChangeParameterMessage_FadeTime
	//	*ChangeParameterMessage_Ipversion
	//	*ChangeParameterMessage_Pins
	Parameter isChangeParameterMessage_Parameter `protobuf_oneof:"parameter"`
}

func (m *ChangeParameterMessage) Reset()                    { *m = ChangeParameterMessage{} }
func (m *ChangeParameterMessage) String() string            { return proto.CompactTextString(m) }
func (*ChangeParameterMessage) ProtoMessage()               {}
func (*ChangeParameterMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isChangeParameterMessage_Parameter interface {
	isChangeParameterMessage_Parameter()
}

type ChangeParameterMessage_ServerName struct {
	ServerName string `protobuf:"bytes,2,opt,name=serverName,oneof"`
}
type ChangeParameterMessage_FadeTime struct {
	FadeTime *FadeTime `protobuf:"bytes,3,opt,name=fadeTime,oneof"`
}
type ChangeParameterMessage_Ipversion struct {
	Ipversion *IPVersion `protobuf:"bytes,4,opt,name=ipversion,oneof"`
}
type ChangeParameterMessage_Pins struct {
	Pins *Pins `protobuf:"bytes,5,opt,name=pins,oneof"`
}

func (*ChangeParameterMessage_ServerName) isChangeParameterMessage_Parameter() {}
func (*ChangeParameterMessage_FadeTime) isChangeParameterMessage_Parameter()   {}
func (*ChangeParameterMessage_Ipversion) isChangeParameterMessage_Parameter()  {}
func (*ChangeParameterMessage_Pins) isChangeParameterMessage_Parameter()       {}

func (m *ChangeParameterMessage) GetParameter() isChangeParameterMessage_Parameter {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func (m *ChangeParameterMessage) GetServerName() string {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_ServerName); ok {
		return x.ServerName
	}
	return ""
}

func (m *ChangeParameterMessage) GetFadeTime() *FadeTime {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_FadeTime); ok {
		return x.FadeTime
	}
	return nil
}

func (m *ChangeParameterMessage) GetIpversion() *IPVersion {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_Ipversion); ok {
		return x.Ipversion
	}
	return nil
}

func (m *ChangeParameterMessage) GetPins() *Pins {
	if x, ok := m.GetParameter().(*ChangeParameterMessage_Pins); ok {
		return x.Pins
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ChangeParameterMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ChangeParameterMessage_OneofMarshaler, _ChangeParameterMessage_OneofUnmarshaler, _ChangeParameterMessage_OneofSizer, []interface{}{
		(*ChangeParameterMessage_ServerName)(nil),
		(*ChangeParameterMessage_FadeTime)(nil),
		(*ChangeParameterMessage_Ipversion)(nil),
		(*ChangeParameterMessage_Pins)(nil),
	}
}

func _ChangeParameterMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ChangeParameterMessage)
	// parameter
	switch x := m.Parameter.(type) {
	case *ChangeParameterMessage_ServerName:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ServerName)
	case *ChangeParameterMessage_FadeTime:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FadeTime); err != nil {
			return err
		}
	case *ChangeParameterMessage_Ipversion:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ipversion); err != nil {
			return err
		}
	case *ChangeParameterMessage_Pins:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pins); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ChangeParameterMessage.Parameter has unexpected type %T", x)
	}
	return nil
}

func _ChangeParameterMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ChangeParameterMessage)
	switch tag {
	case 2: // parameter.serverName
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Parameter = &ChangeParameterMessage_ServerName{x}
		return true, err
	case 3: // parameter.fadeTime
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FadeTime)
		err := b.DecodeMessage(msg)
		m.Parameter = &ChangeParameterMessage_FadeTime{msg}
		return true, err
	case 4: // parameter.ipversion
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IPVersion)
		err := b.DecodeMessage(msg)
		m.Parameter = &ChangeParameterMessage_Ipversion{msg}
		return true, err
	case 5: // parameter.pins
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pins)
		err := b.DecodeMessage(msg)
		m.Parameter = &ChangeParameterMessage_Pins{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ChangeParameterMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ChangeParameterMessage)
	// parameter
	switch x := m.Parameter.(type) {
	case *ChangeParameterMessage_ServerName:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.ServerName)))
		n += len(x.ServerName)
	case *ChangeParameterMessage_FadeTime:
		s := proto.Size(x.FadeTime)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ChangeParameterMessage_Ipversion:
		s := proto.Size(x.Ipversion)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ChangeParameterMessage_Pins:
		s := proto.Size(x.Pins)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type FadeTime struct {
	Millis int32 `protobuf:"varint,1,opt,name=millis" json:"millis,omitempty"`
}

func (m *FadeTime) Reset()                    { *m = FadeTime{} }
func (m *FadeTime) String() string            { return proto.CompactTextString(m) }
func (*FadeTime) ProtoMessage()               {}
func (*FadeTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type IPVersion struct {
	Version IPVersion_Version `protobuf:"varint,1,opt,name=version,enum=LighterGRPC.IPVersion_Version" json:"version,omitempty"`
}

func (m *IPVersion) Reset()                    { *m = IPVersion{} }
func (m *IPVersion) String() string            { return proto.CompactTextString(m) }
func (*IPVersion) ProtoMessage()               {}
func (*IPVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Pins struct {
	RedPin   int32 `protobuf:"varint,1,opt,name=redPin" json:"redPin,omitempty"`
	GreenPin int32 `protobuf:"varint,2,opt,name=greenPin" json:"greenPin,omitempty"`
	BluePin  int32 `protobuf:"varint,3,opt,name=bluePin" json:"bluePin,omitempty"`
	WhitePin int32 `protobuf:"varint,4,opt,name=whitePin" json:"whitePin,omitempty"`
}

func (m *Pins) Reset()                    { *m = Pins{} }
func (m *Pins) String() string            { return proto.CompactTextString(m) }
func (*Pins) ProtoMessage()               {}
func (*Pins) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ScheduledSwitch struct {
	Time     int64  `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Onstate  bool   `protobuf:"varint,2,opt,name=onstate" json:"onstate,omitempty"`
	DeviceID string `protobuf:"bytes,3,opt,name=deviceID" json:"deviceID,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
}

func (m *ScheduledSwitch) Reset()                    { *m = ScheduledSwitch{} }
func (m *ScheduledSwitch) String() string            { return proto.CompactTextString(m) }
func (*ScheduledSwitch) ProtoMessage()               {}
func (*ScheduledSwitch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*ColorMessage)(nil), "LighterGRPC.ColorMessage")
	proto.RegisterType((*StateMessage)(nil), "LighterGRPC.StateMessage")
	proto.RegisterType((*Confirmation)(nil), "LighterGRPC.Confirmation")
	proto.RegisterType((*InitMessage)(nil), "LighterGRPC.InitMessage")
	proto.RegisterType((*LoadConfigRequest)(nil), "LighterGRPC.LoadConfigRequest")
	proto.RegisterType((*ServerConfig)(nil), "LighterGRPC.ServerConfig")
	proto.RegisterType((*ChangeParameterMessage)(nil), "LighterGRPC.ChangeParameterMessage")
	proto.RegisterType((*FadeTime)(nil), "LighterGRPC.FadeTime")
	proto.RegisterType((*IPVersion)(nil), "LighterGRPC.IPVersion")
	proto.RegisterType((*Pins)(nil), "LighterGRPC.Pins")
	proto.RegisterType((*ScheduledSwitch)(nil), "LighterGRPC.ScheduledSwitch")
	proto.RegisterEnum("LighterGRPC.IPVersion_Version", IPVersion_Version_name, IPVersion_Version_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Lighter service

type LighterClient interface {
	SetColor(ctx context.Context, in *ColorMessage, opts ...grpc.CallOption) (*Confirmation, error)
	CheckConnection(ctx context.Context, in *InitMessage, opts ...grpc.CallOption) (Lighter_CheckConnectionClient, error)
	SwitchState(ctx context.Context, in *StateMessage, opts ...grpc.CallOption) (*Confirmation, error)
	LoadServerConfig(ctx context.Context, in *LoadConfigRequest, opts ...grpc.CallOption) (*ServerConfig, error)
	ChangeServerParameter(ctx context.Context, in *ChangeParameterMessage, opts ...grpc.CallOption) (*Confirmation, error)
	ScheduleSwitchState(ctx context.Context, in *ScheduledSwitch, opts ...grpc.CallOption) (*Confirmation, error)
}

type lighterClient struct {
	cc *grpc.ClientConn
}

func NewLighterClient(cc *grpc.ClientConn) LighterClient {
	return &lighterClient{cc}
}

func (c *lighterClient) SetColor(ctx context.Context, in *ColorMessage, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/setColor", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) CheckConnection(ctx context.Context, in *InitMessage, opts ...grpc.CallOption) (Lighter_CheckConnectionClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Lighter_serviceDesc.Streams[0], c.cc, "/LighterGRPC.Lighter/checkConnection", opts...)
	if err != nil {
		return nil, err
	}
	x := &lighterCheckConnectionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Lighter_CheckConnectionClient interface {
	Recv() (*ColorMessage, error)
	grpc.ClientStream
}

type lighterCheckConnectionClient struct {
	grpc.ClientStream
}

func (x *lighterCheckConnectionClient) Recv() (*ColorMessage, error) {
	m := new(ColorMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lighterClient) SwitchState(ctx context.Context, in *StateMessage, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/switchState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) LoadServerConfig(ctx context.Context, in *LoadConfigRequest, opts ...grpc.CallOption) (*ServerConfig, error) {
	out := new(ServerConfig)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/loadServerConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) ChangeServerParameter(ctx context.Context, in *ChangeParameterMessage, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/changeServerParameter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lighterClient) ScheduleSwitchState(ctx context.Context, in *ScheduledSwitch, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/LighterGRPC.Lighter/scheduleSwitchState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lighter service

type LighterServer interface {
	SetColor(context.Context, *ColorMessage) (*Confirmation, error)
	CheckConnection(*InitMessage, Lighter_CheckConnectionServer) error
	SwitchState(context.Context, *StateMessage) (*Confirmation, error)
	LoadServerConfig(context.Context, *LoadConfigRequest) (*ServerConfig, error)
	ChangeServerParameter(context.Context, *ChangeParameterMessage) (*Confirmation, error)
	ScheduleSwitchState(context.Context, *ScheduledSwitch) (*Confirmation, error)
}

func RegisterLighterServer(s *grpc.Server, srv LighterServer) {
	s.RegisterService(&_Lighter_serviceDesc, srv)
}

func _Lighter_SetColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColorMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).SetColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/SetColor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).SetColor(ctx, req.(*ColorMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_CheckConnection_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InitMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LighterServer).CheckConnection(m, &lighterCheckConnectionServer{stream})
}

type Lighter_CheckConnectionServer interface {
	Send(*ColorMessage) error
	grpc.ServerStream
}

type lighterCheckConnectionServer struct {
	grpc.ServerStream
}

func (x *lighterCheckConnectionServer) Send(m *ColorMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Lighter_SwitchState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).SwitchState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/SwitchState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).SwitchState(ctx, req.(*StateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_LoadServerConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).LoadServerConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/LoadServerConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).LoadServerConfig(ctx, req.(*LoadConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_ChangeServerParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeParameterMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).ChangeServerParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/ChangeServerParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).ChangeServerParameter(ctx, req.(*ChangeParameterMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lighter_ScheduleSwitchState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduledSwitch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LighterServer).ScheduleSwitchState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LighterGRPC.Lighter/ScheduleSwitchState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LighterServer).ScheduleSwitchState(ctx, req.(*ScheduledSwitch))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lighter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LighterGRPC.Lighter",
	HandlerType: (*LighterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "setColor",
			Handler:    _Lighter_SetColor_Handler,
		},
		{
			MethodName: "switchState",
			Handler:    _Lighter_SwitchState_Handler,
		},
		{
			MethodName: "loadServerConfig",
			Handler:    _Lighter_LoadServerConfig_Handler,
		},
		{
			MethodName: "changeServerParameter",
			Handler:    _Lighter_ChangeServerParameter_Handler,
		},
		{
			MethodName: "scheduleSwitchState",
			Handler:    _Lighter_ScheduleSwitchState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "checkConnection",
			Handler:       _Lighter_CheckConnection_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("src/proto/lighter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x8e, 0x13, 0x43, 0x9c, 0x13, 0x74, 0x81, 0x41, 0x80, 0x2f, 0xba, 0x42, 0xc8, 0x77, 0x51,
	0x56, 0xa1, 0x0d, 0x08, 0x75, 0x5b, 0xa0, 0x3f, 0xb4, 0x29, 0x44, 0x4e, 0x5b, 0x89, 0x5d, 0x1d,
	0x7b, 0x70, 0x46, 0x18, 0xdb, 0x9d, 0x71, 0x12, 0xf5, 0x49, 0xfa, 0x02, 0x7d, 0xae, 0xbe, 0x47,
	0x77, 0x9d, 0x19, 0x7b, 0x1c, 0x4f, 0x0a, 0xae, 0xd4, 0x5d, 0xbe, 0x39, 0x7f, 0xdf, 0x39, 0x3e,
	0xe7, 0x0b, 0xec, 0x32, 0xea, 0x1f, 0xa5, 0x34, 0xc9, 0x92, 0xa3, 0x88, 0x84, 0x93, 0x0c, 0xd3,
	0x9e, 0x44, 0xa8, 0x3b, 0xc8, 0xe1, 0x6b, 0x77, 0x78, 0xee, 0x7c, 0x37, 0x60, 0xed, 0x3c, 0x89,
	0x12, 0xfa, 0x1e, 0x33, 0xe6, 0x85, 0x18, 0xd9, 0xd0, 0x4e, 0x62, 0x96, 0x79, 0x19, 0xb6, 0x8d,
	0x03, 0xe3, 0xd0, 0x72, 0x15, 0x44, 0x6b, 0x60, 0x50, 0xbb, 0xc9, 0xdf, 0x56, 0x5c, 0x83, 0x0a,
	0x14, 0xda, 0xad, 0x1c, 0x85, 0x02, 0x8d, 0x6d, 0x33, 0x47, 0x63, 0x99, 0x23, 0xf5, 0x7c, 0x92,
	0x7d, 0xb5, 0x57, 0xe4, 0x9b, 0x82, 0x68, 0x0f, 0xac, 0x00, 0xcf, 0x88, 0x8f, 0x2f, 0x2f, 0xec,
	0x55, 0x6e, 0xea, 0xb8, 0x25, 0x16, 0xb6, 0xd4, 0x63, 0x6c, 0x9e, 0xd0, 0xc0, 0x6e, 0xe7, 0x36,
	0x85, 0x9d, 0xcf, 0xb0, 0x36, 0x12, 0x24, 0xfe, 0xcc, 0xb2, 0x5a, 0xa1, 0x59, 0x53, 0xa1, 0xb5,
	0x54, 0xe1, 0x50, 0xcc, 0x21, 0xbe, 0x25, 0xf4, 0xde, 0xcb, 0x48, 0x12, 0x8b, 0x0a, 0x6c, 0xea,
	0xfb, 0xbc, 0x9e, 0xaa, 0x50, 0x40, 0xe7, 0x25, 0x74, 0x2f, 0x63, 0x92, 0x29, 0x2a, 0xd5, 0x82,
	0x46, 0x4d, 0xc1, 0xe6, 0x52, 0xc1, 0x77, 0xb0, 0x39, 0x48, 0xbc, 0x40, 0x16, 0x0d, 0x5d, 0xfc,
	0x65, 0x8a, 0x59, 0xf6, 0xd7, 0xc9, 0xbe, 0xf1, 0xcf, 0x38, 0xc2, 0x74, 0x86, 0x69, 0x9e, 0x0f,
	0xed, 0x03, 0x30, 0x89, 0xaf, 0xbc, 0x7b, 0x5c, 0xa4, 0xaa, 0xbc, 0xa0, 0x67, 0x60, 0xdd, 0x7a,
	0x01, 0xfe, 0x40, 0xb8, 0x55, 0x24, 0xeb, 0xf6, 0xb7, 0x7b, 0x95, 0xbd, 0xe8, 0xbd, 0x2a, 0x8c,
	0x6e, 0xe9, 0x86, 0x4e, 0xa0, 0x43, 0xd2, 0x4f, 0x98, 0x32, 0x3e, 0x1e, 0x39, 0xbe, 0x6e, 0x7f,
	0x47, 0x8b, 0xb9, 0x1c, 0x16, 0x56, 0x77, 0xe1, 0xe8, 0xfc, 0x34, 0x60, 0xe7, 0x7c, 0xe2, 0xc5,
	0x21, 0x1e, 0x7a, 0x94, 0x57, 0xe6, 0xce, 0x95, 0xc9, 0x95, 0x0d, 0x19, 0x7a, 0x43, 0xe8, 0x40,
	0xe3, 0x2f, 0xdb, 0x7d, 0xd3, 0xd0, 0x3a, 0x38, 0xae, 0x74, 0xd0, 0xaa, 0xe9, 0x80, 0x87, 0x2d,
	0x7a, 0x38, 0x15, 0x3d, 0xcc, 0x8a, 0x1e, 0xcc, 0xba, 0x1e, 0x78, 0xd8, 0xc2, 0x15, 0x3d, 0x01,
	0x33, 0x25, 0x31, 0x93, 0xeb, 0xdc, 0xed, 0x6f, 0x6a, 0x21, 0x43, 0x6e, 0xe0, 0xde, 0xd2, 0xe1,
	0xac, 0x0b, 0x9d, 0x54, 0xf5, 0xe9, 0x38, 0x60, 0x29, 0x16, 0x68, 0x07, 0x56, 0xef, 0x49, 0x14,
	0x91, 0x7c, 0x9d, 0x56, 0xdc, 0x02, 0x39, 0x33, 0xe8, 0x94, 0x35, 0xd1, 0x73, 0x68, 0x2b, 0x72,
	0xc2, 0xeb, 0x9f, 0xfe, 0xfe, 0xc3, 0xe4, 0x7a, 0x6a, 0xd0, 0xca, 0xdd, 0x39, 0x82, 0xb6, 0x4a,
	0x62, 0x81, 0x79, 0xf1, 0xf1, 0xc5, 0x60, 0xa3, 0xc1, 0xaf, 0xd2, 0xe2, 0x21, 0x27, 0xd7, 0x57,
	0x83, 0x9b, 0x0d, 0xa3, 0x40, 0xa7, 0x12, 0x35, 0x9d, 0x14, 0x4c, 0x41, 0x5c, 0xf0, 0xa2, 0x38,
	0xe0, 0x3f, 0x15, 0xaf, 0x1c, 0x89, 0x8f, 0x13, 0x52, 0x8c, 0x63, 0x61, 0xc9, 0x8f, 0xbe, 0xc4,
	0xe2, 0x36, 0xc6, 0xd1, 0x14, 0x0b, 0x53, 0xae, 0x00, 0x0a, 0x8a, 0xa8, 0xf9, 0x84, 0x64, 0xd2,
	0x94, 0xcb, 0x41, 0x89, 0x9d, 0x39, 0xac, 0x8f, 0xfc, 0x09, 0x0e, 0xa6, 0x11, 0x0e, 0x46, 0x73,
	0x92, 0xf9, 0x13, 0x84, 0xc0, 0xcc, 0x48, 0xb1, 0x9f, 0x2d, 0x57, 0xfe, 0xae, 0x9e, 0x76, 0xf3,
	0xf1, 0xd3, 0x6e, 0xd5, 0x1c, 0x87, 0xa9, 0xef, 0x52, 0xff, 0x47, 0x0b, 0xda, 0xc5, 0x18, 0xd1,
	0x19, 0x58, 0x0c, 0x67, 0x52, 0xf1, 0xd0, 0xbf, 0xda, 0x70, 0xab, 0x2a, 0xb8, 0xb7, 0x6c, 0x5a,
	0x08, 0x83, 0xd3, 0x40, 0x6f, 0x61, 0x9d, 0xf7, 0xe1, 0xdf, 0xf1, 0xe7, 0x18, 0xfb, 0xb9, 0x5a,
	0xe8, 0xdf, 0x69, 0x21, 0x0f, 0x7b, 0x8f, 0x17, 0x71, 0x1a, 0x4f, 0x0d, 0xc4, 0xc5, 0x84, 0xc9,
	0x59, 0x48, 0x79, 0x5b, 0xa2, 0x54, 0x95, 0xbc, 0x7a, 0x4a, 0xd7, 0xb0, 0x11, 0x71, 0x31, 0xd1,
	0x25, 0x40, 0x0b, 0xf8, 0x4d, 0x6b, 0x96, 0x12, 0x56, 0x43, 0x79, 0xc2, 0x1b, 0xd8, 0xf6, 0xe5,
	0xd5, 0xe6, 0xef, 0xe5, 0xed, 0xa2, 0xff, 0x75, 0x1a, 0x0f, 0x5e, 0x76, 0x3d, 0xd7, 0x21, 0x6c,
	0xb1, 0x62, 0x0f, 0x46, 0x95, 0xd6, 0xff, 0xd3, 0xe9, 0xe8, 0x9b, 0x52, 0x9b, 0xf1, 0x6c, 0x17,
	0xb6, 0x02, 0xdc, 0x0b, 0xe9, 0x34, 0x26, 0xfe, 0x1d, 0xee, 0x15, 0x7f, 0x77, 0x43, 0x63, 0xbc,
	0x2a, 0xff, 0xf1, 0x8e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x51, 0x57, 0xa7, 0x21, 0x0c, 0x07,
	0x00, 0x00,
}
