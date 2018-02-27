// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: directory.proto

/*
	Package directoryapi is a generated protocol buffer package.

	Package directoryapi defines the directory api.

	It is generated from these files:
		directory.proto

	It has these top-level messages:
		RegisterInwayRequest
		RegisterInwayResponse
		ListServicesRequest
		ListServicesResponse
		Service
*/
package directoryapi

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RegisterInwayRequest struct {
	InwayAddress string   `protobuf:"bytes,1,opt,name=inway_address,json=inwayAddress,proto3" json:"inway_address,omitempty"`
	ServiceNames []string `protobuf:"bytes,2,rep,name=service_names,json=serviceNames" json:"service_names,omitempty"`
}

func (m *RegisterInwayRequest) Reset()                    { *m = RegisterInwayRequest{} }
func (m *RegisterInwayRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterInwayRequest) ProtoMessage()               {}
func (*RegisterInwayRequest) Descriptor() ([]byte, []int) { return fileDescriptorDirectory, []int{0} }

type RegisterInwayResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *RegisterInwayResponse) Reset()                    { *m = RegisterInwayResponse{} }
func (m *RegisterInwayResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterInwayResponse) ProtoMessage()               {}
func (*RegisterInwayResponse) Descriptor() ([]byte, []int) { return fileDescriptorDirectory, []int{1} }

type ListServicesRequest struct {
}

func (m *ListServicesRequest) Reset()                    { *m = ListServicesRequest{} }
func (m *ListServicesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListServicesRequest) ProtoMessage()               {}
func (*ListServicesRequest) Descriptor() ([]byte, []int) { return fileDescriptorDirectory, []int{2} }

type ListServicesResponse struct {
	Error    string     `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Services []*Service `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
}

func (m *ListServicesResponse) Reset()                    { *m = ListServicesResponse{} }
func (m *ListServicesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListServicesResponse) ProtoMessage()               {}
func (*ListServicesResponse) Descriptor() ([]byte, []int) { return fileDescriptorDirectory, []int{3} }

type Service struct {
	Name             string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OrganizationName string   `protobuf:"bytes,2,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	InwayAddresses   []string `protobuf:"bytes,3,rep,name=inway_addresses,json=inwayAddresses" json:"inway_addresses,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptorDirectory, []int{4} }

func init() {
	proto.RegisterType((*RegisterInwayRequest)(nil), "directoryapi.RegisterInwayRequest")
	proto.RegisterType((*RegisterInwayResponse)(nil), "directoryapi.RegisterInwayResponse")
	proto.RegisterType((*ListServicesRequest)(nil), "directoryapi.ListServicesRequest")
	proto.RegisterType((*ListServicesResponse)(nil), "directoryapi.ListServicesResponse")
	proto.RegisterType((*Service)(nil), "directoryapi.Service")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Directory service

type DirectoryClient interface {
	// RegisterInway registers an inway for a given service
	RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error)
	// ListServices lists all services and their gateways.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
}

type directoryClient struct {
	cc *grpc.ClientConn
}

func NewDirectoryClient(cc *grpc.ClientConn) DirectoryClient {
	return &directoryClient{cc}
}

func (c *directoryClient) RegisterInway(ctx context.Context, in *RegisterInwayRequest, opts ...grpc.CallOption) (*RegisterInwayResponse, error) {
	out := new(RegisterInwayResponse)
	err := grpc.Invoke(ctx, "/directoryapi.Directory/RegisterInway", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := grpc.Invoke(ctx, "/directoryapi.Directory/ListServices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Directory service

type DirectoryServer interface {
	// RegisterInway registers an inway for a given service
	RegisterInway(context.Context, *RegisterInwayRequest) (*RegisterInwayResponse, error)
	// ListServices lists all services and their gateways.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
}

func RegisterDirectoryServer(s *grpc.Server, srv DirectoryServer) {
	s.RegisterService(&_Directory_serviceDesc, srv)
}

func _Directory_RegisterInway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterInwayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).RegisterInway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/RegisterInway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).RegisterInway(ctx, req.(*RegisterInwayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Directory_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directoryapi.Directory/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Directory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "directoryapi.Directory",
	HandlerType: (*DirectoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterInway",
			Handler:    _Directory_RegisterInway_Handler,
		},
		{
			MethodName: "ListServices",
			Handler:    _Directory_ListServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directory.proto",
}

func (m *RegisterInwayRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterInwayRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.InwayAddress) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDirectory(dAtA, i, uint64(len(m.InwayAddress)))
		i += copy(dAtA[i:], m.InwayAddress)
	}
	if len(m.ServiceNames) > 0 {
		for _, s := range m.ServiceNames {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *RegisterInwayResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterInwayResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDirectory(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	return i, nil
}

func (m *ListServicesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListServicesRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ListServicesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListServicesResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDirectory(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if len(m.Services) > 0 {
		for _, msg := range m.Services {
			dAtA[i] = 0x12
			i++
			i = encodeVarintDirectory(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Service) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Service) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDirectory(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.OrganizationName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDirectory(dAtA, i, uint64(len(m.OrganizationName)))
		i += copy(dAtA[i:], m.OrganizationName)
	}
	if len(m.InwayAddresses) > 0 {
		for _, s := range m.InwayAddresses {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Directory(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Directory(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintDirectory(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegisterInwayRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.InwayAddress)
	if l > 0 {
		n += 1 + l + sovDirectory(uint64(l))
	}
	if len(m.ServiceNames) > 0 {
		for _, s := range m.ServiceNames {
			l = len(s)
			n += 1 + l + sovDirectory(uint64(l))
		}
	}
	return n
}

func (m *RegisterInwayResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovDirectory(uint64(l))
	}
	return n
}

func (m *ListServicesRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ListServicesResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovDirectory(uint64(l))
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovDirectory(uint64(l))
		}
	}
	return n
}

func (m *Service) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDirectory(uint64(l))
	}
	l = len(m.OrganizationName)
	if l > 0 {
		n += 1 + l + sovDirectory(uint64(l))
	}
	if len(m.InwayAddresses) > 0 {
		for _, s := range m.InwayAddresses {
			l = len(s)
			n += 1 + l + sovDirectory(uint64(l))
		}
	}
	return n
}

func sovDirectory(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDirectory(x uint64) (n int) {
	return sovDirectory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterInwayRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDirectory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterInwayRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterInwayRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InwayAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InwayAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceNames = append(m.ServiceNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDirectory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDirectory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterInwayResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDirectory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterInwayResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterInwayResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDirectory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDirectory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListServicesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDirectory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListServicesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListServicesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipDirectory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDirectory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ListServicesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDirectory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ListServicesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListServicesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &Service{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDirectory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDirectory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Service) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDirectory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Service: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Service: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrganizationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrganizationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InwayAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDirectory
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InwayAddresses = append(m.InwayAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDirectory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDirectory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDirectory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDirectory
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
					return 0, ErrIntOverflowDirectory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDirectory
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthDirectory
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDirectory
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipDirectory(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthDirectory = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDirectory   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("directory.proto", fileDescriptorDirectory) }

var fileDescriptorDirectory = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4e, 0xfa, 0x40,
	0x14, 0xc5, 0x29, 0xfc, 0x3f, 0x64, 0x04, 0xc1, 0x11, 0x92, 0xa6, 0x31, 0x0d, 0x0e, 0x0b, 0x4d,
	0x0c, 0x6d, 0xc4, 0x27, 0xd0, 0xb8, 0x31, 0x31, 0x2e, 0xea, 0xce, 0x0d, 0x0e, 0x30, 0xd6, 0x49,
	0xa0, 0x53, 0xe7, 0x0e, 0x12, 0x5c, 0xfa, 0x0a, 0xbe, 0x14, 0x4b, 0x13, 0x5f, 0x00, 0x89, 0x0f,
	0x62, 0x3a, 0x1d, 0x48, 0x6b, 0x08, 0xbb, 0xb9, 0xa7, 0xa7, 0xa7, 0xe7, 0xfe, 0x3a, 0xa8, 0x36,
	0xe4, 0x92, 0x0d, 0x94, 0x90, 0x33, 0x2f, 0x96, 0x42, 0x09, 0x5c, 0x59, 0x0b, 0x34, 0xe6, 0xce,
	0x61, 0x28, 0x44, 0x38, 0x62, 0x3e, 0x8d, 0xb9, 0x4f, 0xa3, 0x48, 0x28, 0xaa, 0xb8, 0x88, 0x20,
	0xf5, 0x3a, 0x9d, 0x90, 0xab, 0xa7, 0x49, 0xdf, 0x1b, 0x88, 0xb1, 0x1f, 0x8a, 0x50, 0xf8, 0x5a,
	0xee, 0x4f, 0x1e, 0xf5, 0xa4, 0x07, 0x7d, 0x4a, 0xed, 0xe4, 0x01, 0x35, 0x02, 0x16, 0x72, 0x50,
	0x4c, 0x5e, 0x47, 0x53, 0x3a, 0x0b, 0xd8, 0xf3, 0x84, 0x81, 0xc2, 0x6d, 0x54, 0xe5, 0xc9, 0xdc,
	0xa3, 0xc3, 0xa1, 0x64, 0x00, 0xb6, 0xd5, 0xb2, 0x4e, 0xca, 0x41, 0x45, 0x8b, 0x17, 0xa9, 0x96,
	0x98, 0x80, 0xc9, 0x17, 0x3e, 0x60, 0xbd, 0x88, 0x8e, 0x19, 0xd8, 0xc5, 0x56, 0x29, 0x31, 0x19,
	0xf1, 0x36, 0xd1, 0x48, 0x07, 0x35, 0x7f, 0x7d, 0x01, 0x62, 0x11, 0x01, 0xc3, 0x0d, 0xf4, 0x97,
	0x49, 0x29, 0xa4, 0x89, 0x4e, 0x07, 0xd2, 0x44, 0x07, 0x37, 0x1c, 0xd4, 0x5d, 0x1a, 0x01, 0xa6,
	0x0f, 0xe9, 0xa1, 0x46, 0x5e, 0xde, 0x16, 0x82, 0xcf, 0xd0, 0x8e, 0xe9, 0x90, 0x76, 0xda, 0xed,
	0x36, 0xbd, 0x2c, 0x43, 0xcf, 0xe4, 0x04, 0x6b, 0x1b, 0x01, 0xf4, 0xdf, 0x88, 0x18, 0xa3, 0x3f,
	0xc9, 0x3a, 0x26, 0x52, 0x9f, 0xf1, 0x29, 0xda, 0x17, 0x32, 0xa4, 0x11, 0x7f, 0xd5, 0xb4, 0xf5,
	0xbe, 0x76, 0x51, 0x1b, 0xea, 0xd9, 0x07, 0xc9, 0xce, 0xf8, 0x18, 0xd5, 0x72, 0xf0, 0x18, 0xd8,
	0x25, 0x4d, 0x66, 0x2f, 0x8b, 0x8f, 0x41, 0x77, 0x61, 0xa1, 0xf2, 0xd5, 0xaa, 0x17, 0xbe, 0x47,
	0xd5, 0x1c, 0x29, 0x4c, 0xf2, 0xa5, 0x37, 0xfd, 0x28, 0xa7, 0xbd, 0xd5, 0x93, 0x52, 0x22, 0x05,
	0x3c, 0x45, 0x95, 0x2c, 0x3f, 0x7c, 0x94, 0x7f, 0x6d, 0x03, 0x72, 0x87, 0x6c, 0xb3, 0x98, 0xe0,
	0xd6, 0xdb, 0xe7, 0xf7, 0x7b, 0xd1, 0xc1, 0xb6, 0xbf, 0xf6, 0xfa, 0x23, 0x0e, 0xaa, 0xb3, 0xe2,
	0x7a, 0x59, 0x9f, 0x7f, 0xb9, 0x85, 0xf9, 0xd2, 0xb5, 0x3e, 0x96, 0xae, 0xb5, 0x58, 0xba, 0x56,
	0xff, 0x9f, 0xbe, 0x79, 0xe7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x65, 0x8b, 0x38, 0xe7,
	0x02, 0x00, 0x00,
}
