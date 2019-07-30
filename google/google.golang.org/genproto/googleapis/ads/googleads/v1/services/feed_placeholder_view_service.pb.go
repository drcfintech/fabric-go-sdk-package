// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/feed_placeholder_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for
// [FeedPlaceholderViewService.GetFeedPlaceholderView][google.ads.googleads.v1.services.FeedPlaceholderViewService.GetFeedPlaceholderView].
type GetFeedPlaceholderViewRequest struct {
	// The resource name of the feed placeholder view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedPlaceholderViewRequest) Reset()         { *m = GetFeedPlaceholderViewRequest{} }
func (m *GetFeedPlaceholderViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedPlaceholderViewRequest) ProtoMessage()    {}
func (*GetFeedPlaceholderViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c806439cff2804f, []int{0}
}

func (m *GetFeedPlaceholderViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Unmarshal(m, b)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.Merge(m, src)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Size(m)
}
func (m *GetFeedPlaceholderViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedPlaceholderViewRequest proto.InternalMessageInfo

func (m *GetFeedPlaceholderViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedPlaceholderViewRequest)(nil), "google.ads.googleads.v1.services.GetFeedPlaceholderViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/feed_placeholder_view_service.proto", fileDescriptor_8c806439cff2804f)
}

var fileDescriptor_8c806439cff2804f = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0x26, 0xb9, 0x70, 0xe1, 0x86, 0x7b, 0x37, 0x59, 0x5c, 0x24, 0x28, 0x96, 0xda, 0x85, 0x74,
	0x31, 0x43, 0x14, 0x8a, 0x8c, 0x54, 0x49, 0x29, 0xd6, 0x95, 0x94, 0x0a, 0x59, 0x48, 0x20, 0x8c,
	0x99, 0x63, 0x0c, 0x24, 0x99, 0x98, 0x49, 0xd3, 0x85, 0xb8, 0x71, 0xe1, 0x0b, 0xf8, 0x06, 0x2e,
	0x7d, 0x14, 0x77, 0xe2, 0x2b, 0xb8, 0x72, 0xe5, 0x23, 0x48, 0x3a, 0x9d, 0x88, 0xd2, 0xd8, 0xdd,
	0xc7, 0x9c, 0xef, 0xe7, 0xfc, 0x8c, 0x31, 0x0c, 0x39, 0x0f, 0x63, 0xc0, 0x94, 0x09, 0x2c, 0x61,
	0x85, 0x4a, 0x1b, 0x0b, 0xc8, 0xcb, 0x28, 0x00, 0x81, 0x2f, 0x00, 0x98, 0x9f, 0xc5, 0x34, 0x80,
	0x4b, 0x1e, 0x33, 0xc8, 0xfd, 0x32, 0x82, 0x99, 0xbf, 0x28, 0xa3, 0x2c, 0xe7, 0x05, 0x37, 0x5b,
	0x52, 0x8a, 0x28, 0x13, 0xa8, 0x76, 0x41, 0xa5, 0x8d, 0x94, 0x8b, 0xd5, 0x6f, 0xca, 0xc9, 0x41,
	0xf0, 0x69, 0xde, 0x18, 0x24, 0x03, 0xac, 0x75, 0x25, 0xcf, 0x22, 0x4c, 0xd3, 0x94, 0x17, 0xb4,
	0x88, 0x78, 0x2a, 0x64, 0xb5, 0x3d, 0x34, 0x36, 0x46, 0x50, 0x1c, 0x01, 0xb0, 0xf1, 0xa7, 0xdc,
	0x8d, 0x60, 0x36, 0x81, 0xab, 0x29, 0x88, 0xc2, 0xdc, 0x32, 0xfe, 0xa9, 0x1c, 0x3f, 0xa5, 0x09,
	0xac, 0x69, 0x2d, 0x6d, 0xfb, 0xcf, 0xe4, 0xaf, 0x7a, 0x3c, 0xa1, 0x09, 0xec, 0xbc, 0x6b, 0x86,
	0xb5, 0xc4, 0xe3, 0x54, 0x8e, 0x60, 0x3e, 0x6b, 0xc6, 0xff, 0xe5, 0x29, 0xe6, 0x21, 0x5a, 0x35,
	0x3f, 0xfa, 0xb1, 0x3f, 0xab, 0xd7, 0x68, 0x50, 0xaf, 0x07, 0x2d, 0x91, 0xb7, 0x0f, 0x6e, 0x5f,
	0x5e, 0xef, 0xf5, 0x3d, 0xb3, 0x57, 0x6d, 0xf2, 0xfa, 0xcb, 0x88, 0xfd, 0x60, 0x2a, 0x0a, 0x9e,
	0x40, 0x2e, 0x70, 0x77, 0xbe, 0xda, 0x6f, 0x5a, 0x81, 0xbb, 0x37, 0x83, 0x3b, 0xdd, 0xe8, 0x04,
	0x3c, 0x59, 0xd9, 0xfe, 0x60, 0xb3, 0x79, 0x31, 0xe3, 0xea, 0x04, 0x63, 0xed, 0xec, 0x78, 0x61,
	0x12, 0xf2, 0x98, 0xa6, 0x21, 0xe2, 0x79, 0x88, 0x43, 0x48, 0xe7, 0x07, 0x52, 0x17, 0xcf, 0x22,
	0xd1, 0xfc, 0xd1, 0xf6, 0x15, 0x78, 0xd0, 0x7f, 0x8d, 0x1c, 0xe7, 0x51, 0x6f, 0x8d, 0xa4, 0xa1,
	0xc3, 0x04, 0x92, 0xb0, 0x42, 0xae, 0x8d, 0x16, 0xc1, 0xe2, 0x49, 0x51, 0x3c, 0x87, 0x09, 0xaf,
	0xa6, 0x78, 0xae, 0xed, 0x29, 0xca, 0x9b, 0xde, 0x91, 0xef, 0x84, 0x38, 0x4c, 0x10, 0x52, 0x93,
	0x08, 0x71, 0x6d, 0x42, 0x14, 0xed, 0xfc, 0xf7, 0xbc, 0xcf, 0xdd, 0x8f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xee, 0xb8, 0x0d, 0xf6, 0x0f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedPlaceholderViewServiceClient is the client API for FeedPlaceholderViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedPlaceholderViewServiceClient interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error)
}

type feedPlaceholderViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedPlaceholderViewServiceClient(cc *grpc.ClientConn) FeedPlaceholderViewServiceClient {
	return &feedPlaceholderViewServiceClient{cc}
}

func (c *feedPlaceholderViewServiceClient) GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error) {
	out := new(resources.FeedPlaceholderView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.FeedPlaceholderViewService/GetFeedPlaceholderView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedPlaceholderViewServiceServer is the server API for FeedPlaceholderViewService service.
type FeedPlaceholderViewServiceServer interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(context.Context, *GetFeedPlaceholderViewRequest) (*resources.FeedPlaceholderView, error)
}

func RegisterFeedPlaceholderViewServiceServer(s *grpc.Server, srv FeedPlaceholderViewServiceServer) {
	s.RegisterService(&_FeedPlaceholderViewService_serviceDesc, srv)
}

func _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedPlaceholderViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.FeedPlaceholderViewService/GetFeedPlaceholderView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, req.(*GetFeedPlaceholderViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedPlaceholderViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.FeedPlaceholderViewService",
	HandlerType: (*FeedPlaceholderViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedPlaceholderView",
			Handler:    _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/feed_placeholder_view_service.proto",
}
