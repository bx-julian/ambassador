// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/route/v3/srds.proto

package envoy_service_route_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/service/discovery/v3"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type SrdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrdsDummy) Reset()         { *m = SrdsDummy{} }
func (m *SrdsDummy) String() string { return proto.CompactTextString(m) }
func (*SrdsDummy) ProtoMessage()    {}
func (*SrdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e461818506941dcb, []int{0}
}
func (m *SrdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SrdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SrdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SrdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrdsDummy.Merge(m, src)
}
func (m *SrdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *SrdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SrdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SrdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SrdsDummy)(nil), "envoy.service.route.v3.SrdsDummy")
}

func init() { proto.RegisterFile("envoy/service/route/v3/srds.proto", fileDescriptor_e461818506941dcb) }

var fileDescriptor_e461818506941dcb = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x1d, 0x0b, 0x8a, 0xb3, 0x6b, 0x16, 0x22, 0xb1, 0x0d, 0xb6, 0xb8, 0x90, 0xd2, 0x4e,
	0x6a, 0x03, 0x2e, 0xba, 0xac, 0xc5, 0x75, 0x69, 0x9e, 0x20, 0x26, 0xc7, 0x38, 0xd0, 0xce, 0x89,
	0x33, 0x93, 0xd1, 0xec, 0xc4, 0x55, 0x71, 0x2b, 0xb8, 0x70, 0xe1, 0xca, 0x57, 0x11, 0x5c, 0x0a,
	0xbe, 0x80, 0x14, 0xdf, 0xa0, 0x2f, 0x20, 0x99, 0xe4, 0xe6, 0x26, 0xdc, 0xcb, 0xe5, 0xde, 0xcd,
	0x5d, 0x26, 0xe7, 0x3b, 0xe7, 0xff, 0xff, 0xe1, 0xa7, 0x23, 0x10, 0x06, 0x0b, 0x5f, 0x81, 0x34,
	0x3c, 0x06, 0x5f, 0x62, 0xae, 0xc1, 0x37, 0x81, 0xaf, 0x64, 0xa2, 0x58, 0x26, 0x51, 0xa3, 0xf3,
	0xd0, 0x22, 0xac, 0x46, 0x98, 0x45, 0x98, 0x09, 0xdc, 0x49, 0x77, 0x35, 0xe1, 0x2a, 0x46, 0x03,
	0xb2, 0x28, 0xd7, 0x9b, 0x8f, 0xea, 0x86, 0x3b, 0x48, 0x11, 0xd3, 0x1d, 0xf8, 0x51, 0xc6, 0xfd,
	0x48, 0x08, 0xd4, 0x91, 0xe6, 0x28, 0x6a, 0x05, 0x77, 0x94, 0x27, 0x59, 0xd4, 0xfe, 0xef, 0x1b,
	0x90, 0x8a, 0xa3, 0xe0, 0x22, 0xad, 0x91, 0x27, 0x95, 0x58, 0x9b, 0x91, 0xa0, 0x30, 0x97, 0x31,
	0x54, 0xc4, 0x78, 0x42, 0x1f, 0x84, 0x32, 0x51, 0xeb, 0x7c, 0xbf, 0x2f, 0x96, 0xc3, 0x6f, 0x3f,
	0x0f, 0xde, 0x23, 0x5a, 0x5b, 0x8f, 0x32, 0xce, 0xcc, 0x82, 0x35, 0xe3, 0xc5, 0xa9, 0x47, 0x07,
	0x61, 0x8c, 0x19, 0x24, 0xdb, 0x32, 0x8d, 0x5a, 0x9f, 0xd9, 0x0d, 0xab, 0x30, 0xce, 0x7b, 0xea,
	0x84, 0x5a, 0x42, 0xb4, 0x6f, 0x53, 0xce, 0x94, 0x75, 0x9f, 0xe2, 0x3c, 0xa5, 0x09, 0x58, 0x73,
	0x63, 0x0b, 0xef, 0x72, 0x50, 0xda, 0x9d, 0x5d, 0x93, 0x56, 0x19, 0x0a, 0x05, 0xe3, 0x3b, 0xcf,
	0xc8, 0x9c, 0x38, 0x1f, 0x09, 0xed, 0xaf, 0x61, 0xa7, 0xa3, 0x8e, 0xf0, 0xf3, 0x2b, 0x4f, 0x95,
	0xf8, 0x05, 0xf5, 0xc5, 0x4d, 0x56, 0x3a, 0x16, 0xbe, 0x13, 0xda, 0x7f, 0x05, 0x3a, 0x7e, 0x7b,
	0x7b, 0xd9, 0xa7, 0x9f, 0xfe, 0xfc, 0xfb, 0x72, 0x77, 0x38, 0x7e, 0xdc, 0xe9, 0xce, 0x52, 0x59,
	0x03, 0x33, 0xdb, 0x38, 0x65, 0x91, 0xde, 0x92, 0x4c, 0xdc, 0x17, 0x9f, 0x7f, 0x7c, 0x3d, 0xdd,
	0x9f, 0xd3, 0x5a, 0x23, 0x46, 0xf1, 0x86, 0xa7, 0x4d, 0x2f, 0x59, 0xcb, 0xf6, 0x4b, 0x3b, 0xca,
	0xa5, 0x6d, 0xcc, 0x6a, 0xf5, 0xeb, 0xe8, 0x91, 0xdf, 0x47, 0x8f, 0xfc, 0x3d, 0x7a, 0x84, 0x3e,
	0xe5, 0x58, 0x1d, 0xc8, 0x24, 0x7e, 0x28, 0xd8, 0xe5, 0x25, 0x5f, 0xd9, 0x4e, 0x6d, 0xca, 0x82,
	0x6d, 0xc8, 0x81, 0x90, 0xd7, 0xf7, 0x6c, 0xd9, 0x82, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x85,
	0x27, 0xde, 0xd6, 0x38, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScopedRoutesDiscoveryServiceClient is the client API for ScopedRoutesDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScopedRoutesDiscoveryServiceClient interface {
	StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error)
	DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error)
	FetchScopedRoutes(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type scopedRoutesDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewScopedRoutesDiscoveryServiceClient(cc *grpc.ClientConn) ScopedRoutesDiscoveryServiceClient {
	return &scopedRoutesDiscoveryServiceClient{cc}
}

func (c *scopedRoutesDiscoveryServiceClient) StreamScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_StreamScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[0], "/envoy.service.route.v3.ScopedRoutesDiscoveryService/StreamScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceStreamScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) DeltaScopedRoutes(ctx context.Context, opts ...grpc.CallOption) (ScopedRoutesDiscoveryService_DeltaScopedRoutesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ScopedRoutesDiscoveryService_serviceDesc.Streams[1], "/envoy.service.route.v3.ScopedRoutesDiscoveryService/DeltaScopedRoutes", opts...)
	if err != nil {
		return nil, err
	}
	x := &scopedRoutesDiscoveryServiceDeltaScopedRoutesClient{stream}
	return x, nil
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesClient struct {
	grpc.ClientStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *scopedRoutesDiscoveryServiceClient) FetchScopedRoutes(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.route.v3.ScopedRoutesDiscoveryService/FetchScopedRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScopedRoutesDiscoveryServiceServer is the server API for ScopedRoutesDiscoveryService service.
type ScopedRoutesDiscoveryServiceServer interface {
	StreamScopedRoutes(ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error
	DeltaScopedRoutes(ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error
	FetchScopedRoutes(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedScopedRoutesDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedScopedRoutesDiscoveryServiceServer struct {
}

func (*UnimplementedScopedRoutesDiscoveryServiceServer) StreamScopedRoutes(srv ScopedRoutesDiscoveryService_StreamScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) DeltaScopedRoutes(srv ScopedRoutesDiscoveryService_DeltaScopedRoutesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaScopedRoutes not implemented")
}
func (*UnimplementedScopedRoutesDiscoveryServiceServer) FetchScopedRoutes(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchScopedRoutes not implemented")
}

func RegisterScopedRoutesDiscoveryServiceServer(s *grpc.Server, srv ScopedRoutesDiscoveryServiceServer) {
	s.RegisterService(&_ScopedRoutesDiscoveryService_serviceDesc, srv)
}

func _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).StreamScopedRoutes(&scopedRoutesDiscoveryServiceStreamScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_StreamScopedRoutesServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceStreamScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceStreamScopedRoutesServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ScopedRoutesDiscoveryServiceServer).DeltaScopedRoutes(&scopedRoutesDiscoveryServiceDeltaScopedRoutesServer{stream})
}

type ScopedRoutesDiscoveryService_DeltaScopedRoutesServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type scopedRoutesDiscoveryServiceDeltaScopedRoutesServer struct {
	grpc.ServerStream
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *scopedRoutesDiscoveryServiceDeltaScopedRoutesServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.route.v3.ScopedRoutesDiscoveryService/FetchScopedRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScopedRoutesDiscoveryServiceServer).FetchScopedRoutes(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ScopedRoutesDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.route.v3.ScopedRoutesDiscoveryService",
	HandlerType: (*ScopedRoutesDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchScopedRoutes",
			Handler:    _ScopedRoutesDiscoveryService_FetchScopedRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_StreamScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaScopedRoutes",
			Handler:       _ScopedRoutesDiscoveryService_DeltaScopedRoutes_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/route/v3/srds.proto",
}

func (m *SrdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SrdsDummy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SrdsDummy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintSrds(dAtA []byte, offset int, v uint64) int {
	offset -= sovSrds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SrdsDummy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSrds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSrds(x uint64) (n int) {
	return sovSrds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SrdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSrds
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SrdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SrdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSrds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSrds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSrds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSrds
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSrds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSrds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSrds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSrds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSrds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSrds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSrds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSrds = fmt.Errorf("proto: unexpected end of group")
)
