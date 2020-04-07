// Code generated by protoc-gen-go. DO NOT EDIT.
// source: callback.proto

package callback

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A type of request message
type Mode int32

const (
	Mode_Invoke    Mode = 0
	Mode_NewStream Mode = 1
	Mode_Header    Mode = 2
	Mode_Trailer   Mode = 3
	Mode_Close     Mode = 4
	Mode_Msg       Mode = 5
	Mode_Recv      Mode = 6
	Mode_Error     Mode = 7
)

var Mode_name = map[int32]string{
	0: "Invoke",
	1: "NewStream",
	2: "Header",
	3: "Trailer",
	4: "Close",
	5: "Msg",
	6: "Recv",
	7: "Error",
}

var Mode_value = map[string]int32{
	"Invoke":    0,
	"NewStream": 1,
	"Header":    2,
	"Trailer":   3,
	"Close":     4,
	"Msg":       5,
	"Recv":      6,
	"Error":     7,
}

func (x Mode) String() string {
	return proto.EnumName(Mode_name, int32(x))
}

func (Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6cf7fe261a3a1c45, []int{0}
}

// Request passed
type Request struct {
	Mode                 Mode              `protobuf:"varint,1,opt,name=mode,proto3,enum=callback.Mode" json:"mode,omitempty"`
	Method               string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	OpId                 int64             `protobuf:"varint,3,opt,name=opId,proto3" json:"opId,omitempty"`
	Args                 *any.Any          `protobuf:"bytes,4,opt,name=args,proto3" json:"args,omitempty"`
	Reply                *any.Any          `protobuf:"bytes,5,opt,name=reply,proto3" json:"reply,omitempty"`
	Descr                *StreamDescriptor `protobuf:"bytes,6,opt,name=descr,proto3" json:"descr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cf7fe261a3a1c45, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_Invoke
}

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *Request) GetArgs() *any.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Request) GetReply() *any.Any {
	if m != nil {
		return m.Reply
	}
	return nil
}

func (m *Request) GetDescr() *StreamDescriptor {
	if m != nil {
		return m.Descr
	}
	return nil
}

type StreamDescriptor struct {
	StreamName string `protobuf:"bytes,1,opt,name=StreamName,proto3" json:"StreamName,omitempty"`
	// At least one of these is true.
	ServerStreams        bool     `protobuf:"varint,2,opt,name=ServerStreams,proto3" json:"ServerStreams,omitempty"`
	ClientStreams        bool     `protobuf:"varint,3,opt,name=ClientStreams,proto3" json:"ClientStreams,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamDescriptor) Reset()         { *m = StreamDescriptor{} }
func (m *StreamDescriptor) String() string { return proto.CompactTextString(m) }
func (*StreamDescriptor) ProtoMessage()    {}
func (*StreamDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cf7fe261a3a1c45, []int{1}
}

func (m *StreamDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamDescriptor.Unmarshal(m, b)
}
func (m *StreamDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamDescriptor.Marshal(b, m, deterministic)
}
func (m *StreamDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamDescriptor.Merge(m, src)
}
func (m *StreamDescriptor) XXX_Size() int {
	return xxx_messageInfo_StreamDescriptor.Size(m)
}
func (m *StreamDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_StreamDescriptor proto.InternalMessageInfo

func (m *StreamDescriptor) GetStreamName() string {
	if m != nil {
		return m.StreamName
	}
	return ""
}

func (m *StreamDescriptor) GetServerStreams() bool {
	if m != nil {
		return m.ServerStreams
	}
	return false
}

func (m *StreamDescriptor) GetClientStreams() bool {
	if m != nil {
		return m.ClientStreams
	}
	return false
}

type MDValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MDValue) Reset()         { *m = MDValue{} }
func (m *MDValue) String() string { return proto.CompactTextString(m) }
func (*MDValue) ProtoMessage()    {}
func (*MDValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cf7fe261a3a1c45, []int{2}
}

func (m *MDValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDValue.Unmarshal(m, b)
}
func (m *MDValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDValue.Marshal(b, m, deterministic)
}
func (m *MDValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDValue.Merge(m, src)
}
func (m *MDValue) XXX_Size() int {
	return xxx_messageInfo_MDValue.Size(m)
}
func (m *MDValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MDValue.DiscardUnknown(m)
}

var xxx_messageInfo_MDValue proto.InternalMessageInfo

func (m *MDValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MDValue) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type MD struct {
	Mds                  []*MDValue `protobuf:"bytes,1,rep,name=mds,proto3" json:"mds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MD) Reset()         { *m = MD{} }
func (m *MD) String() string { return proto.CompactTextString(m) }
func (*MD) ProtoMessage()    {}
func (*MD) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cf7fe261a3a1c45, []int{3}
}

func (m *MD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MD.Unmarshal(m, b)
}
func (m *MD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MD.Marshal(b, m, deterministic)
}
func (m *MD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MD.Merge(m, src)
}
func (m *MD) XXX_Size() int {
	return xxx_messageInfo_MD.Size(m)
}
func (m *MD) XXX_DiscardUnknown() {
	xxx_messageInfo_MD.DiscardUnknown(m)
}

var xxx_messageInfo_MD proto.InternalMessageInfo

func (m *MD) GetMds() []*MDValue {
	if m != nil {
		return m.Mds
	}
	return nil
}

// Response send
type Response struct {
	Mode                 Mode            `protobuf:"varint,1,opt,name=mode,proto3,enum=callback.Mode" json:"mode,omitempty"`
	Status               *ResponseStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Method               string          `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	OpId                 int64           `protobuf:"varint,4,opt,name=opId,proto3" json:"opId,omitempty"`
	Response             *any.Any        `protobuf:"bytes,5,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cf7fe261a3a1c45, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_Invoke
}

func (m *Response) GetStatus() *ResponseStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Response) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Response) GetOpId() int64 {
	if m != nil {
		return m.OpId
	}
	return 0
}

func (m *Response) GetResponse() *any.Any {
	if m != nil {
		return m.Response
	}
	return nil
}

type ResponseStatus struct {
	Code                 int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details              []*any.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ResponseStatus) Reset()         { *m = ResponseStatus{} }
func (m *ResponseStatus) String() string { return proto.CompactTextString(m) }
func (*ResponseStatus) ProtoMessage()    {}
func (*ResponseStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cf7fe261a3a1c45, []int{5}
}

func (m *ResponseStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseStatus.Unmarshal(m, b)
}
func (m *ResponseStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseStatus.Marshal(b, m, deterministic)
}
func (m *ResponseStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseStatus.Merge(m, src)
}
func (m *ResponseStatus) XXX_Size() int {
	return xxx_messageInfo_ResponseStatus.Size(m)
}
func (m *ResponseStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseStatus proto.InternalMessageInfo

func (m *ResponseStatus) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ResponseStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ResponseStatus) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterEnum("callback.Mode", Mode_name, Mode_value)
	proto.RegisterType((*Request)(nil), "callback.Request")
	proto.RegisterType((*StreamDescriptor)(nil), "callback.StreamDescriptor")
	proto.RegisterType((*MDValue)(nil), "callback.MDValue")
	proto.RegisterType((*MD)(nil), "callback.MD")
	proto.RegisterType((*Response)(nil), "callback.Response")
	proto.RegisterType((*ResponseStatus)(nil), "callback.ResponseStatus")
}

func init() { proto.RegisterFile("callback.proto", fileDescriptor_6cf7fe261a3a1c45) }

var fileDescriptor_6cf7fe261a3a1c45 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x71, 0xfc, 0x95, 0x4c, 0xd4, 0xd4, 0x1d, 0x55, 0xc8, 0xea, 0x01, 0x59, 0x86, 0x83,
	0xe9, 0x21, 0x84, 0x70, 0xe5, 0x82, 0x12, 0xa4, 0xf6, 0x90, 0x22, 0x6d, 0x10, 0xf7, 0x8d, 0x3d,
	0x84, 0x28, 0xb6, 0xd7, 0xec, 0x3a, 0x41, 0xb9, 0xf0, 0x68, 0xbc, 0x0b, 0x6f, 0x82, 0xd6, 0x6b,
	0xb7, 0x09, 0x1f, 0x15, 0x87, 0x48, 0x33, 0xf3, 0xff, 0x69, 0x26, 0x33, 0xff, 0x35, 0x8c, 0x52,
	0x9e, 0xe7, 0x2b, 0x9e, 0x6e, 0xc7, 0x95, 0x14, 0xb5, 0xc0, 0x7e, 0x97, 0x5f, 0x5d, 0x56, 0xf5,
	0xa1, 0x22, 0xf5, 0x8a, 0x97, 0x07, 0xfd, 0x33, 0x7a, 0xfc, 0xd3, 0x02, 0x9f, 0xd1, 0xd7, 0x1d,
	0xa9, 0x1a, 0x63, 0x70, 0x0a, 0x91, 0x51, 0x68, 0x45, 0x56, 0x32, 0x9a, 0x8e, 0xc6, 0xf7, 0xad,
	0x16, 0x22, 0x23, 0xd6, 0x68, 0xf8, 0x14, 0xbc, 0x82, 0xea, 0x2f, 0x22, 0x0b, 0x7b, 0x91, 0x95,
	0x0c, 0x58, 0x9b, 0x21, 0x82, 0x23, 0xaa, 0xdb, 0x2c, 0xb4, 0x23, 0x2b, 0xb1, 0x59, 0x13, 0x63,
	0x02, 0x0e, 0x97, 0x6b, 0x15, 0x3a, 0x91, 0x95, 0x0c, 0xa7, 0x97, 0xe3, 0xb5, 0x10, 0xeb, 0x9c,
	0xcc, 0xe0, 0xd5, 0xee, 0xf3, 0xf8, 0x5d, 0x79, 0x60, 0x0d, 0x81, 0xd7, 0xe0, 0x4a, 0xaa, 0xf2,
	0x43, 0xe8, 0x3e, 0x82, 0x1a, 0x04, 0x27, 0xe0, 0x66, 0xa4, 0x52, 0x19, 0x7a, 0x0d, 0x7b, 0xf5,
	0xf0, 0x37, 0x97, 0xb5, 0x24, 0x5e, 0xcc, 0xb5, 0xb8, 0xa9, 0x6a, 0x21, 0x99, 0x01, 0xe3, 0xef,
	0x10, 0xfc, 0x2e, 0xe1, 0x33, 0x00, 0x53, 0xbb, 0xe3, 0x85, 0xd9, 0x78, 0xc0, 0x8e, 0x2a, 0xf8,
	0x02, 0xce, 0x96, 0x24, 0xf7, 0x24, 0x4d, 0x4d, 0x35, 0xeb, 0xf6, 0xd9, 0x69, 0x51, 0x53, 0xb3,
	0x7c, 0x43, 0x65, 0xdd, 0x51, 0xb6, 0xa1, 0x4e, 0x8a, 0xf1, 0x6b, 0xf0, 0x17, 0xf3, 0x4f, 0x3c,
	0xdf, 0x11, 0x06, 0x60, 0x6f, 0xe9, 0xd0, 0xce, 0xd3, 0x21, 0x5e, 0x82, 0xbb, 0xd7, 0x52, 0xd8,
	0x8b, 0xec, 0x64, 0xc0, 0x4c, 0x12, 0xbf, 0x84, 0xde, 0x62, 0x8e, 0xcf, 0xc1, 0x2e, 0x32, 0x15,
	0x5a, 0x91, 0x9d, 0x0c, 0xa7, 0x17, 0x47, 0x7e, 0x98, 0x6e, 0x4c, 0xab, 0xf1, 0x0f, 0x0b, 0xfa,
	0x8c, 0x54, 0x25, 0x4a, 0x45, 0xff, 0x65, 0xe1, 0x04, 0x3c, 0x55, 0xf3, 0x7a, 0x67, 0x76, 0x1a,
	0x4e, 0xc3, 0x07, 0xaa, 0xeb, 0xb3, 0x6c, 0x74, 0xd6, 0x72, 0x47, 0xa6, 0xdb, 0x7f, 0x35, 0xdd,
	0x39, 0x32, 0x7d, 0x02, 0x7d, 0xd9, 0x76, 0x79, 0xd4, 0xcd, 0x7b, 0x2a, 0x2e, 0x61, 0x74, 0x3a,
	0x57, 0xf7, 0x4d, 0xbb, 0x2d, 0x5c, 0xd6, 0xc4, 0x18, 0x82, 0x5f, 0x90, 0x52, 0x7c, 0x4d, 0xed,
	0xcb, 0xeb, 0x52, 0x1c, 0x83, 0x9f, 0x51, 0xcd, 0x37, 0xb9, 0x3e, 0xbf, 0xfd, 0xcf, 0x81, 0x1d,
	0x74, 0x9d, 0x82, 0xa3, 0xaf, 0x81, 0x00, 0xde, 0x6d, 0xb9, 0x17, 0x5b, 0x0a, 0x9e, 0xe0, 0x19,
	0x0c, 0xee, 0xe8, 0x9b, 0x31, 0x2c, 0xb0, 0xb4, 0x74, 0x43, 0x3c, 0x23, 0x19, 0xf4, 0x70, 0x08,
	0xfe, 0x47, 0xc9, 0x37, 0x39, 0xc9, 0xc0, 0xc6, 0x01, 0xb8, 0xb3, 0x5c, 0x28, 0x0a, 0x1c, 0xf4,
	0xc1, 0x5e, 0xa8, 0x75, 0xe0, 0x62, 0x1f, 0x1c, 0x46, 0xe9, 0x3e, 0xf0, 0xb4, 0xfa, 0x5e, 0x4a,
	0x21, 0x03, 0x7f, 0xfa, 0x01, 0xce, 0x67, 0xed, 0x55, 0xf5, 0x93, 0xd9, 0xa4, 0x84, 0x6f, 0xe1,
	0xfc, 0x86, 0x97, 0x59, 0x4e, 0x9d, 0xa0, 0x10, 0xff, 0x3c, 0xfd, 0xd5, 0xc5, 0x71, 0xad, 0xf9,
	0x30, 0x13, 0x6b, 0x62, 0xad, 0xbc, 0x66, 0x99, 0x37, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfc,
	0x27, 0x94, 0x43, 0xe1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CallbackServiceClient is the client API for CallbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallbackServiceClient interface {
	HandleCallbacks(ctx context.Context, opts ...grpc.CallOption) (CallbackService_HandleCallbacksClient, error)
}

type callbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCallbackServiceClient(cc grpc.ClientConnInterface) CallbackServiceClient {
	return &callbackServiceClient{cc}
}

func (c *callbackServiceClient) HandleCallbacks(ctx context.Context, opts ...grpc.CallOption) (CallbackService_HandleCallbacksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CallbackService_serviceDesc.Streams[0], "/callback.CallbackService/HandleCallbacks", opts...)
	if err != nil {
		return nil, err
	}
	x := &callbackServiceHandleCallbacksClient{stream}
	return x, nil
}

type CallbackService_HandleCallbacksClient interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ClientStream
}

type callbackServiceHandleCallbacksClient struct {
	grpc.ClientStream
}

func (x *callbackServiceHandleCallbacksClient) Send(m *Response) error {
	return x.ClientStream.SendMsg(m)
}

func (x *callbackServiceHandleCallbacksClient) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CallbackServiceServer is the server API for CallbackService service.
type CallbackServiceServer interface {
	HandleCallbacks(CallbackService_HandleCallbacksServer) error
}

// UnimplementedCallbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCallbackServiceServer struct {
}

func (*UnimplementedCallbackServiceServer) HandleCallbacks(srv CallbackService_HandleCallbacksServer) error {
	return status.Errorf(codes.Unimplemented, "method HandleCallbacks not implemented")
}

func RegisterCallbackServiceServer(s *grpc.Server, srv CallbackServiceServer) {
	s.RegisterService(&_CallbackService_serviceDesc, srv)
}

func _CallbackService_HandleCallbacks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CallbackServiceServer).HandleCallbacks(&callbackServiceHandleCallbacksServer{stream})
}

type CallbackService_HandleCallbacksServer interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ServerStream
}

type callbackServiceHandleCallbacksServer struct {
	grpc.ServerStream
}

func (x *callbackServiceHandleCallbacksServer) Send(m *Request) error {
	return x.ServerStream.SendMsg(m)
}

func (x *callbackServiceHandleCallbacksServer) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CallbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "callback.CallbackService",
	HandlerType: (*CallbackServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HandleCallbacks",
			Handler:       _CallbackService_HandleCallbacks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "callback.proto",
}