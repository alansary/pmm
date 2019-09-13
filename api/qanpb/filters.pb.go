// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qanpb/filters.proto

package qanpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// FiltersRequest contains period for which we need filters.
type FiltersRequest struct {
	PeriodStartFrom      *timestamp.Timestamp `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	MainMetricName       string               `protobuf:"bytes,3,opt,name=main_metric_name,json=mainMetricName,proto3" json:"main_metric_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FiltersRequest) Reset()         { *m = FiltersRequest{} }
func (m *FiltersRequest) String() string { return proto.CompactTextString(m) }
func (*FiltersRequest) ProtoMessage()    {}
func (*FiltersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{0}
}

func (m *FiltersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiltersRequest.Unmarshal(m, b)
}
func (m *FiltersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiltersRequest.Marshal(b, m, deterministic)
}
func (m *FiltersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiltersRequest.Merge(m, src)
}
func (m *FiltersRequest) XXX_Size() int {
	return xxx_messageInfo_FiltersRequest.Size(m)
}
func (m *FiltersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FiltersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FiltersRequest proto.InternalMessageInfo

func (m *FiltersRequest) GetPeriodStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartFrom
	}
	return nil
}

func (m *FiltersRequest) GetPeriodStartTo() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartTo
	}
	return nil
}

func (m *FiltersRequest) GetMainMetricName() string {
	if m != nil {
		return m.MainMetricName
	}
	return ""
}

// FiltersReply is map of labels for given period by key.
// Key is label's name and value is label's value and how many times it occur.
type FiltersReply struct {
	Labels               map[string]*ListLabels `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FiltersReply) Reset()         { *m = FiltersReply{} }
func (m *FiltersReply) String() string { return proto.CompactTextString(m) }
func (*FiltersReply) ProtoMessage()    {}
func (*FiltersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{1}
}

func (m *FiltersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiltersReply.Unmarshal(m, b)
}
func (m *FiltersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiltersReply.Marshal(b, m, deterministic)
}
func (m *FiltersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiltersReply.Merge(m, src)
}
func (m *FiltersReply) XXX_Size() int {
	return xxx_messageInfo_FiltersReply.Size(m)
}
func (m *FiltersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FiltersReply.DiscardUnknown(m)
}

var xxx_messageInfo_FiltersReply proto.InternalMessageInfo

func (m *FiltersReply) GetLabels() map[string]*ListLabels {
	if m != nil {
		return m.Labels
	}
	return nil
}

// ListLabels is list of label's values: duplicates are impossible.
type ListLabels struct {
	Name                 []*Values `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListLabels) Reset()         { *m = ListLabels{} }
func (m *ListLabels) String() string { return proto.CompactTextString(m) }
func (*ListLabels) ProtoMessage()    {}
func (*ListLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{2}
}

func (m *ListLabels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLabels.Unmarshal(m, b)
}
func (m *ListLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLabels.Marshal(b, m, deterministic)
}
func (m *ListLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLabels.Merge(m, src)
}
func (m *ListLabels) XXX_Size() int {
	return xxx_messageInfo_ListLabels.Size(m)
}
func (m *ListLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLabels.DiscardUnknown(m)
}

var xxx_messageInfo_ListLabels proto.InternalMessageInfo

func (m *ListLabels) GetName() []*Values {
	if m != nil {
		return m.Name
	}
	return nil
}

// Values is label values and main metric percent and per second.
type Values struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	MainMetricPercent    float32  `protobuf:"fixed32,2,opt,name=main_metric_percent,json=mainMetricPercent,proto3" json:"main_metric_percent,omitempty"`
	MainMetricPerSec     float32  `protobuf:"fixed32,3,opt,name=main_metric_per_sec,json=mainMetricPerSec,proto3" json:"main_metric_per_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Values) Reset()         { *m = Values{} }
func (m *Values) String() string { return proto.CompactTextString(m) }
func (*Values) ProtoMessage()    {}
func (*Values) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{3}
}

func (m *Values) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Values.Unmarshal(m, b)
}
func (m *Values) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Values.Marshal(b, m, deterministic)
}
func (m *Values) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Values.Merge(m, src)
}
func (m *Values) XXX_Size() int {
	return xxx_messageInfo_Values.Size(m)
}
func (m *Values) XXX_DiscardUnknown() {
	xxx_messageInfo_Values.DiscardUnknown(m)
}

var xxx_messageInfo_Values proto.InternalMessageInfo

func (m *Values) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Values) GetMainMetricPercent() float32 {
	if m != nil {
		return m.MainMetricPercent
	}
	return 0
}

func (m *Values) GetMainMetricPerSec() float32 {
	if m != nil {
		return m.MainMetricPerSec
	}
	return 0
}

func init() {
	proto.RegisterType((*FiltersRequest)(nil), "qan.FiltersRequest")
	proto.RegisterType((*FiltersReply)(nil), "qan.FiltersReply")
	proto.RegisterMapType((map[string]*ListLabels)(nil), "qan.FiltersReply.LabelsEntry")
	proto.RegisterType((*ListLabels)(nil), "qan.ListLabels")
	proto.RegisterType((*Values)(nil), "qan.Values")
}

func init() { proto.RegisterFile("qanpb/filters.proto", fileDescriptor_2662fef59f58e02c) }

var fileDescriptor_2662fef59f58e02c = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5f, 0x6b, 0x13, 0x41,
	0x10, 0xe7, 0xee, 0x6c, 0x4a, 0x27, 0xda, 0x24, 0x1b, 0x1f, 0x8e, 0x43, 0x6d, 0x38, 0x10, 0x82,
	0xd0, 0x3b, 0x89, 0x08, 0xd2, 0xc7, 0x82, 0x29, 0x48, 0x2b, 0xb2, 0x2d, 0x22, 0xbe, 0x1c, 0x9b,
	0x38, 0x29, 0x87, 0xb7, 0xbb, 0x77, 0xbb, 0x93, 0x42, 0x1e, 0x7c, 0xf1, 0x1b, 0x88, 0xdf, 0xc9,
	0x2f, 0xe0, 0x57, 0xf0, 0x83, 0x48, 0x76, 0x2f, 0x36, 0x21, 0x0f, 0xbe, 0xdd, 0xcc, 0xef, 0xcf,
	0xce, 0xfc, 0x6e, 0x60, 0xd8, 0x08, 0x55, 0xcf, 0xf2, 0x45, 0x59, 0x11, 0x1a, 0x9b, 0xd5, 0x46,
	0x93, 0x66, 0x51, 0x23, 0x54, 0xf2, 0xe4, 0x56, 0xeb, 0xdb, 0x0a, 0x73, 0x51, 0x97, 0xb9, 0x50,
	0x4a, 0x93, 0xa0, 0x52, 0xab, 0x96, 0x92, 0x9c, 0xb4, 0xa8, 0xab, 0x66, 0xcb, 0x45, 0x4e, 0xa5,
	0x44, 0x4b, 0x42, 0xd6, 0x9e, 0x90, 0xfe, 0x0a, 0xe0, 0x78, 0xea, 0x5d, 0x39, 0x36, 0x4b, 0xb4,
	0xc4, 0xa6, 0x30, 0xa8, 0xd1, 0x94, 0xfa, 0x4b, 0x61, 0x49, 0x18, 0x2a, 0x16, 0x46, 0xcb, 0x38,
	0x18, 0x05, 0xe3, 0xee, 0x24, 0xc9, 0xbc, 0x5f, 0xb6, 0xf1, 0xcb, 0x6e, 0x36, 0x7e, 0xbc, 0xe7,
	0x45, 0xd7, 0x6b, 0xcd, 0xd4, 0x68, 0xc9, 0xce, 0xa1, 0xb7, 0xe3, 0x43, 0x3a, 0x0e, 0xff, 0xeb,
	0xf2, 0x68, 0xcb, 0xe5, 0x46, 0xb3, 0x31, 0xf4, 0xa5, 0x28, 0x55, 0x21, 0x91, 0x4c, 0x39, 0x2f,
	0x94, 0x90, 0x18, 0x47, 0xa3, 0x60, 0x7c, 0xc4, 0x8f, 0xd7, 0xfd, 0x2b, 0xd7, 0x7e, 0x2f, 0x24,
	0xa6, 0x3f, 0x02, 0x78, 0xf8, 0x6f, 0x91, 0xba, 0x5a, 0xb1, 0xd7, 0xd0, 0xa9, 0xc4, 0x0c, 0x2b,
	0x1b, 0x07, 0xa3, 0x68, 0xdc, 0x9d, 0x3c, 0xcd, 0x1a, 0xa1, 0xb2, 0x6d, 0x4a, 0x76, 0xe9, 0xf0,
	0xb7, 0x8a, 0xcc, 0x8a, 0xb7, 0xe4, 0xe4, 0x1d, 0x74, 0xb7, 0xda, 0xac, 0x0f, 0xd1, 0x57, 0x5c,
	0xb9, 0xf5, 0x8f, 0xf8, 0xfa, 0x93, 0x3d, 0x87, 0x83, 0x3b, 0x51, 0x2d, 0xb1, 0x5d, 0xa6, 0xe7,
	0x6c, 0x2f, 0x4b, 0x4b, 0x5e, 0xc6, 0x3d, 0x7a, 0x16, 0xbe, 0x09, 0xd2, 0x53, 0x80, 0x7b, 0x80,
	0x9d, 0xc0, 0x03, 0x37, 0xbf, 0x1f, 0xa7, 0xeb, 0x74, 0x1f, 0xd7, 0x5c, 0xcb, 0x1d, 0x90, 0x7e,
	0x83, 0x8e, 0xaf, 0xd9, 0xe3, 0xcd, 0x1b, 0xfe, 0x5d, 0x5f, 0xb0, 0x0c, 0x86, 0xdb, 0x61, 0xd4,
	0x68, 0xe6, 0xa8, 0xc8, 0xcd, 0x11, 0xf2, 0xc1, 0x7d, 0x1e, 0x1f, 0x3c, 0xc0, 0x4e, 0xf7, 0xf8,
	0x85, 0xc5, 0xb9, 0xcb, 0x2f, 0xe4, 0xfd, 0x1d, 0xfe, 0x35, 0xce, 0x27, 0x9f, 0xe0, 0xb0, 0x4d,
	0x87, 0x5d, 0x41, 0x74, 0x81, 0xc4, 0x86, 0xbb, 0x91, 0xb9, 0xf3, 0x48, 0x06, 0x7b, 0x39, 0xa6,
	0xcf, 0xbe, 0xff, 0xfe, 0xf3, 0x33, 0x8c, 0xd3, 0x61, 0x7e, 0xf7, 0x32, 0x6f, 0x84, 0xca, 0x5b,
	0x34, 0xbf, 0x40, 0x3a, 0x0b, 0x5e, 0x9c, 0x1f, 0x7e, 0x3e, 0x70, 0xf7, 0x3b, 0xeb, 0xb8, 0x3f,
	0xfe, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x4e, 0x55, 0xf3, 0xcf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FiltersClient is the client API for Filters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FiltersClient interface {
	// Get gets map of metrics names.
	Get(ctx context.Context, in *FiltersRequest, opts ...grpc.CallOption) (*FiltersReply, error)
}

type filtersClient struct {
	cc *grpc.ClientConn
}

func NewFiltersClient(cc *grpc.ClientConn) FiltersClient {
	return &filtersClient{cc}
}

func (c *filtersClient) Get(ctx context.Context, in *FiltersRequest, opts ...grpc.CallOption) (*FiltersReply, error) {
	out := new(FiltersReply)
	err := c.cc.Invoke(ctx, "/qan.Filters/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FiltersServer is the server API for Filters service.
type FiltersServer interface {
	// Get gets map of metrics names.
	Get(context.Context, *FiltersRequest) (*FiltersReply, error)
}

// UnimplementedFiltersServer can be embedded to have forward compatible implementations.
type UnimplementedFiltersServer struct {
}

func (*UnimplementedFiltersServer) Get(ctx context.Context, req *FiltersRequest) (*FiltersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterFiltersServer(s *grpc.Server, srv FiltersServer) {
	s.RegisterService(&_Filters_serviceDesc, srv)
}

func _Filters_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiltersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.Filters/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiltersServer).Get(ctx, req.(*FiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Filters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qan.Filters",
	HandlerType: (*FiltersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Filters_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/filters.proto",
}