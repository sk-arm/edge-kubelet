// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_MessageType int32

const (
	Message_Open  Message_MessageType = 0
	Message_Close Message_MessageType = 1
	Message_Data  Message_MessageType = 2
)

var Message_MessageType_name = map[int32]string{
	0: "Open",
	1: "Close",
	2: "Data",
}
var Message_MessageType_value = map[string]int32{
	"Open":  0,
	"Close": 1,
	"Data":  2,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_protocol_8abd51438c247ef6, []int{0, 0}
}

type ClosePayload_CloseCode int32

const (
	ClosePayload_Normal            ClosePayload_CloseCode = 0
	ClosePayload_Abnormal          ClosePayload_CloseCode = 1
	ClosePayload_ProtocolError     ClosePayload_CloseCode = 2
	ClosePayload_NodeRouteNotFound ClosePayload_CloseCode = 3
)

var ClosePayload_CloseCode_name = map[int32]string{
	0: "Normal",
	1: "Abnormal",
	2: "ProtocolError",
	3: "NodeRouteNotFound",
}
var ClosePayload_CloseCode_value = map[string]int32{
	"Normal":            0,
	"Abnormal":          1,
	"ProtocolError":     2,
	"NodeRouteNotFound": 3,
}

func (x ClosePayload_CloseCode) String() string {
	return proto.EnumName(ClosePayload_CloseCode_name, int32(x))
}
func (ClosePayload_CloseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_protocol_8abd51438c247ef6, []int{2, 0}
}

type Message struct {
	Type Message_MessageType `protobuf:"varint,1,opt,name=type,enum=Message_MessageType" json:"type,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Message_OpenPayload
	//	*Message_ClosePayload
	//	*Message_DataPayload
	Payload              isMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_8abd51438c247ef6, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_OpenPayload struct {
	OpenPayload *OpenPayload `protobuf:"bytes,2,opt,name=open_payload,json=openPayload,oneof"`
}
type Message_ClosePayload struct {
	ClosePayload *ClosePayload `protobuf:"bytes,3,opt,name=close_payload,json=closePayload,oneof"`
}
type Message_DataPayload struct {
	DataPayload []byte `protobuf:"bytes,4,opt,name=data_payload,json=dataPayload,proto3,oneof"`
}

func (*Message_OpenPayload) isMessage_Payload()  {}
func (*Message_ClosePayload) isMessage_Payload() {}
func (*Message_DataPayload) isMessage_Payload()  {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetType() Message_MessageType {
	if m != nil {
		return m.Type
	}
	return Message_Open
}

func (m *Message) GetOpenPayload() *OpenPayload {
	if x, ok := m.GetPayload().(*Message_OpenPayload); ok {
		return x.OpenPayload
	}
	return nil
}

func (m *Message) GetClosePayload() *ClosePayload {
	if x, ok := m.GetPayload().(*Message_ClosePayload); ok {
		return x.ClosePayload
	}
	return nil
}

func (m *Message) GetDataPayload() []byte {
	if x, ok := m.GetPayload().(*Message_DataPayload); ok {
		return x.DataPayload
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_OpenPayload)(nil),
		(*Message_ClosePayload)(nil),
		(*Message_DataPayload)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_OpenPayload:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OpenPayload); err != nil {
			return err
		}
	case *Message_ClosePayload:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClosePayload); err != nil {
			return err
		}
	case *Message_DataPayload:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.DataPayload)
	case nil:
	default:
		return fmt.Errorf("Message.Payload has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 2: // payload.open_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OpenPayload)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_OpenPayload{msg}
		return true, err
	case 3: // payload.close_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClosePayload)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_ClosePayload{msg}
		return true, err
	case 4: // payload.data_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Payload = &Message_DataPayload{x}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_OpenPayload:
		s := proto.Size(x.OpenPayload)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ClosePayload:
		s := proto.Size(x.ClosePayload)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_DataPayload:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.DataPayload)))
		n += len(x.DataPayload)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type OpenPayload struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenPayload) Reset()         { *m = OpenPayload{} }
func (m *OpenPayload) String() string { return proto.CompactTextString(m) }
func (*OpenPayload) ProtoMessage()    {}
func (*OpenPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_8abd51438c247ef6, []int{1}
}
func (m *OpenPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenPayload.Unmarshal(m, b)
}
func (m *OpenPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenPayload.Marshal(b, m, deterministic)
}
func (dst *OpenPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenPayload.Merge(dst, src)
}
func (m *OpenPayload) XXX_Size() int {
	return xxx_messageInfo_OpenPayload.Size(m)
}
func (m *OpenPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenPayload.DiscardUnknown(m)
}

var xxx_messageInfo_OpenPayload proto.InternalMessageInfo

func (m *OpenPayload) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *OpenPayload) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ClosePayload struct {
	Code                 ClosePayload_CloseCode `protobuf:"varint,1,opt,name=code,enum=ClosePayload_CloseCode" json:"code,omitempty"`
	Reason               string                 `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ClosePayload) Reset()         { *m = ClosePayload{} }
func (m *ClosePayload) String() string { return proto.CompactTextString(m) }
func (*ClosePayload) ProtoMessage()    {}
func (*ClosePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_8abd51438c247ef6, []int{2}
}
func (m *ClosePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClosePayload.Unmarshal(m, b)
}
func (m *ClosePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClosePayload.Marshal(b, m, deterministic)
}
func (dst *ClosePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClosePayload.Merge(dst, src)
}
func (m *ClosePayload) XXX_Size() int {
	return xxx_messageInfo_ClosePayload.Size(m)
}
func (m *ClosePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClosePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ClosePayload proto.InternalMessageInfo

func (m *ClosePayload) GetCode() ClosePayload_CloseCode {
	if m != nil {
		return m.Code
	}
	return ClosePayload_Normal
}

func (m *ClosePayload) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*OpenPayload)(nil), "OpenPayload")
	proto.RegisterType((*ClosePayload)(nil), "ClosePayload")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("ClosePayload_CloseCode", ClosePayload_CloseCode_name, ClosePayload_CloseCode_value)
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_protocol_8abd51438c247ef6) }

var fileDescriptor_protocol_8abd51438c247ef6 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x14, 0x64, 0xa1, 0x1f, 0xd0, 0xd7, 0x85, 0x94, 0x97, 0x4f, 0xe5, 0x48, 0xea, 0xa5, 0x89, 0xa6,
	0x89, 0xe8, 0xcd, 0x93, 0xa2, 0xc6, 0x8b, 0x40, 0x1a, 0xef, 0x66, 0xe9, 0x6e, 0xbc, 0xd4, 0xbe,
	0x66, 0x5b, 0x0e, 0xfc, 0x1f, 0x7f, 0xa5, 0x27, 0xb3, 0xdb, 0xa5, 0xd4, 0xd3, 0xee, 0x4c, 0x66,
	0x26, 0x33, 0x79, 0x30, 0x2d, 0x35, 0xd5, 0x94, 0x51, 0x9e, 0xd8, 0x4f, 0xf4, 0xc3, 0x60, 0xf4,
	0xa6, 0xaa, 0x4a, 0x7c, 0x2a, 0x8c, 0xc1, 0xab, 0x0f, 0xa5, 0x9a, 0xb3, 0x05, 0x8b, 0xa7, 0xcb,
	0xff, 0x89, 0xe3, 0x8f, 0xef, 0xfb, 0xa1, 0x54, 0xa9, 0x55, 0xe0, 0x0d, 0x70, 0x2a, 0x55, 0xf1,
	0x51, 0x8a, 0x43, 0x4e, 0x42, 0xce, 0xfb, 0x0b, 0x16, 0x07, 0x4b, 0x9e, 0x6c, 0x4a, 0x55, 0x6c,
	0x1b, 0xee, 0xb5, 0x97, 0x06, 0x74, 0x82, 0x78, 0x07, 0x93, 0x2c, 0xa7, 0x4a, 0xb5, 0x9e, 0x81,
	0xf5, 0x4c, 0x92, 0x95, 0x61, 0x4f, 0x26, 0x9e, 0x75, 0x30, 0x5e, 0x02, 0x97, 0xa2, 0x16, 0xad,
	0xc9, 0x5b, 0xb0, 0x98, 0x9b, 0x68, 0xc3, 0x3a, 0x51, 0x74, 0x0d, 0x41, 0xa7, 0x22, 0x8e, 0xc1,
	0x33, 0x3d, 0xc2, 0x1e, 0xfa, 0xf0, 0xcf, 0xa6, 0x87, 0xcc, 0x90, 0x4f, 0xa2, 0x16, 0x61, 0xff,
	0xd1, 0x87, 0x91, 0x4b, 0x8b, 0xee, 0x21, 0xe8, 0x34, 0x46, 0x04, 0xaf, 0x20, 0xd9, 0xec, 0xf7,
	0x53, 0xfb, 0xc7, 0x39, 0x8c, 0x84, 0x94, 0x5a, 0x55, 0x95, 0x1d, 0xe9, 0xa7, 0x47, 0x18, 0x7d,
	0x33, 0xe0, 0xdd, 0xee, 0x78, 0x05, 0x5e, 0x76, 0xb4, 0x4f, 0x97, 0x17, 0x7f, 0x86, 0x35, 0x60,
	0x45, 0x52, 0xa5, 0x56, 0x84, 0xe7, 0x30, 0xd4, 0x4a, 0x54, 0x54, 0xb8, 0x58, 0x87, 0xa2, 0x0d,
	0xf8, 0xad, 0x14, 0x01, 0x86, 0x6b, 0xd2, 0x5f, 0x22, 0x0f, 0x7b, 0xc8, 0x61, 0xfc, 0xb0, 0x2b,
	0x1a, 0xc4, 0x70, 0x06, 0x93, 0xad, 0x3b, 0xe4, 0xb3, 0xd6, 0xa4, 0xc3, 0x3e, 0x9e, 0xc1, 0x6c,
	0x6d, 0xf2, 0x69, 0x5f, 0xab, 0x35, 0xd5, 0x2f, 0xb4, 0x2f, 0x64, 0x38, 0xd8, 0x0d, 0xed, 0x9d,
	0x6f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x8d, 0xca, 0x07, 0xf9, 0x01, 0x00, 0x00,
}
