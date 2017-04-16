// Code generated by protoc-gen-go.
// source: team.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	team.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	GetRequest
	GetResponse
	IsAvailableRequest
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type CreateRequest struct {
	Name               string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName        string            `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Email              string            `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	UserName           string            `protobuf:"bytes,4,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Password           string            `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	InviteEmails       []string          `protobuf:"bytes,6,rep,name=invite_emails,json=inviteEmails" json:"invite_emails,omitempty"`
	SubscriptionType   string            `protobuf:"bytes,7,opt,name=subscription_type,json=subscriptionType" json:"subscription_type,omitempty"`
	ClientIp           string            `protobuf:"bytes,8,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	PaymentMethodNonce string            `protobuf:"bytes,9,opt,name=payment_method_nonce,json=paymentMethodNonce" json:"payment_method_nonce,omitempty"`
	Options            map[string]string `protobuf:"bytes,10,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRequest) GetInviteEmails() []string {
	if m != nil {
		return m.InviteEmails
	}
	return nil
}

func (m *CreateRequest) GetSubscriptionType() string {
	if m != nil {
		return m.SubscriptionType
	}
	return ""
}

func (m *CreateRequest) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *CreateRequest) GetPaymentMethodNonce() string {
	if m != nil {
		return m.PaymentMethodNonce
	}
	return ""
}

func (m *CreateRequest) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

type CreateResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string                  `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetResponse) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

type IsAvailableRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *IsAvailableRequest) Reset()                    { *m = IsAvailableRequest{} }
func (m *IsAvailableRequest) String() string            { return proto.CompactTextString(m) }
func (*IsAvailableRequest) ProtoMessage()               {}
func (*IsAvailableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IsAvailableRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "appscode.namespace.v1beta1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "appscode.namespace.v1beta1.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "appscode.namespace.v1beta1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "appscode.namespace.v1beta1.GetResponse")
	proto.RegisterType((*IsAvailableRequest)(nil), "appscode.namespace.v1beta1.IsAvailableRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Teams service

type TeamsClient interface {
	// Creates a new namespace, if name is valid and no namespace with same name exists.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Gets a namespace, if exists. This can be used to check existence of a namespace.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Check if a namespace name is available meaning name is valid and no namespace with same name exists.
	IsAvailable(ctx context.Context, in *IsAvailableRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type teamsClient struct {
	cc *grpc.ClientConn
}

func NewTeamsClient(cc *grpc.ClientConn) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) IsAvailable(ctx context.Context, in *IsAvailableRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.namespace.v1beta1.Teams/IsAvailable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Teams service

type TeamsServer interface {
	// Creates a new namespace, if name is valid and no namespace with same name exists.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Gets a namespace, if exists. This can be used to check existence of a namespace.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Check if a namespace name is available meaning name is valid and no namespace with same name exists.
	IsAvailable(context.Context, *IsAvailableRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_IsAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).IsAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.namespace.v1beta1.Teams/IsAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).IsAvailable(ctx, req.(*IsAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.namespace.v1beta1.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Teams_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Teams_Get_Handler,
		},
		{
			MethodName: "IsAvailable",
			Handler:    _Teams_IsAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team.proto",
}

func init() { proto.RegisterFile("team.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x95, 0x66, 0xeb, 0xda, 0xd3, 0xed, 0xa7, 0xfd, 0xac, 0x49, 0x44, 0xe1, 0x5f, 0x09,
	0x12, 0x94, 0x4d, 0x24, 0x5b, 0x41, 0x08, 0x55, 0x48, 0xb0, 0xa1, 0x69, 0xda, 0xc5, 0xc6, 0x14,
	0x26, 0x2e, 0xb8, 0x89, 0xdc, 0xc4, 0xda, 0x0c, 0x49, 0x6c, 0x62, 0xb7, 0x28, 0x42, 0xdc, 0xec,
	0x15, 0xb8, 0xe1, 0x86, 0x27, 0xe0, 0x9a, 0x37, 0xe0, 0x0d, 0x78, 0x05, 0x1e, 0x80, 0x47, 0x40,
	0xb6, 0x93, 0xad, 0xd3, 0xb4, 0x52, 0xb8, 0x89, 0xec, 0xf3, 0x39, 0x7f, 0xbe, 0x3e, 0x3e, 0x31,
	0x80, 0x24, 0x38, 0xf3, 0x79, 0xc1, 0x24, 0x43, 0x2e, 0xe6, 0x5c, 0xc4, 0x2c, 0x21, 0x7e, 0x8e,
	0x33, 0x22, 0x38, 0x8e, 0x89, 0x3f, 0xde, 0x18, 0x12, 0x89, 0x37, 0xdc, 0x6b, 0x47, 0x8c, 0x1d,
	0xa5, 0x24, 0xc0, 0x9c, 0x06, 0x38, 0xcf, 0x99, 0xc4, 0x92, 0xb2, 0x5c, 0x98, 0x48, 0xf7, 0x46,
	0x1d, 0x79, 0x09, 0xbf, 0x79, 0x8e, 0x27, 0xb2, 0xe4, 0x44, 0x04, 0xfa, 0x6b, 0x1c, 0xbc, 0xef,
	0x36, 0x2c, 0x3d, 0x2f, 0x08, 0x96, 0x24, 0x24, 0xef, 0x46, 0x44, 0x48, 0x84, 0x60, 0x4e, 0xa9,
	0x70, 0xac, 0xae, 0xd5, 0x6b, 0x87, 0x7a, 0x8d, 0x6e, 0xc1, 0x62, 0x42, 0x05, 0x4f, 0x71, 0x19,
	0x69, 0xd6, 0xd0, 0xac, 0x53, 0xd9, 0xf6, 0x95, 0xcb, 0x0a, 0xcc, 0x93, 0x0c, 0xd3, 0xd4, 0xb1,
	0x35, 0x33, 0x1b, 0x74, 0x15, 0xda, 0x23, 0x41, 0x0a, 0x13, 0x35, 0xa7, 0x49, 0x4b, 0x19, 0x74,
	0x88, 0x0b, 0x2d, 0x8e, 0x85, 0x78, 0xcf, 0x8a, 0xc4, 0x99, 0x37, 0xac, 0xde, 0xa3, 0xdb, 0xb0,
	0x44, 0xf3, 0x31, 0x95, 0x24, 0xd2, 0x89, 0x84, 0xd3, 0xec, 0xda, 0xbd, 0x76, 0xb8, 0x68, 0x8c,
	0xdb, 0xda, 0x86, 0xd6, 0xe0, 0x7f, 0x31, 0x1a, 0x8a, 0xb8, 0xa0, 0x5c, 0x1d, 0x3a, 0x52, 0x07,
	0x73, 0x16, 0x74, 0xa6, 0xe5, 0x49, 0x70, 0x58, 0x72, 0xa2, 0xa4, 0xc4, 0x29, 0x25, 0xb9, 0x8c,
	0x28, 0x77, 0x5a, 0xa6, 0x9c, 0x31, 0xec, 0x72, 0xb4, 0x0e, 0x2b, 0x1c, 0x97, 0x99, 0xa2, 0x19,
	0x91, 0xc7, 0x2c, 0x89, 0x72, 0x96, 0xc7, 0xc4, 0x69, 0x6b, 0x3f, 0x54, 0xb1, 0x3d, 0x8d, 0xf6,
	0x15, 0x41, 0x07, 0xb0, 0xc0, 0x74, 0x72, 0xe1, 0x40, 0xd7, 0xee, 0x75, 0xfa, 0x8f, 0xfc, 0xcb,
	0x6f, 0xd1, 0x3f, 0xd7, 0x62, 0xff, 0x85, 0x09, 0xdc, 0xce, 0x65, 0x51, 0x86, 0x75, 0x1a, 0x77,
	0x00, 0x8b, 0x93, 0x00, 0x2d, 0x83, 0xfd, 0x96, 0x94, 0xd5, 0x3d, 0xa8, 0xa5, 0xea, 0xf1, 0x18,
	0xa7, 0xa3, 0xba, 0xff, 0x66, 0x33, 0x68, 0x3c, 0xb6, 0xbc, 0x4d, 0xf8, 0xaf, 0x2e, 0x21, 0x38,
	0xcb, 0x05, 0x41, 0x01, 0x34, 0x85, 0xc4, 0x72, 0x24, 0x74, 0x82, 0x4e, 0xff, 0xca, 0x99, 0x3c,
	0x33, 0x06, 0xfe, 0x4b, 0x8d, 0xc3, 0xca, 0xcd, 0xeb, 0x02, 0xec, 0x10, 0x39, 0x65, 0x0a, 0xbc,
	0x10, 0x3a, 0xda, 0xe3, 0x1f, 0x2b, 0xa8, 0x9c, 0xfc, 0x98, 0x26, 0x95, 0x7a, 0xbd, 0xf6, 0x7a,
	0x80, 0x76, 0xc5, 0xe6, 0x18, 0xd3, 0x14, 0x0f, 0xd3, 0x69, 0x33, 0xd8, 0xff, 0x65, 0xc3, 0xfc,
	0x21, 0xc1, 0x99, 0x40, 0x5f, 0x2c, 0x68, 0x9a, 0xd3, 0xa2, 0x7b, 0x33, 0x37, 0xdd, 0x5d, 0x9d,
	0xc5, 0xd5, 0x1c, 0xcd, 0x7b, 0x72, 0xf2, 0xcd, 0x69, 0xb4, 0xac, 0x93, 0x1f, 0x3f, 0x3f, 0x35,
	0xd6, 0xbd, 0xb5, 0x20, 0x3a, 0xf7, 0x1b, 0x9d, 0x86, 0x07, 0x55, 0x78, 0xa0, 0xfe, 0x65, 0x11,
	0xbc, 0x11, 0x2c, 0x1f, 0x58, 0xab, 0xe8, 0xb3, 0x05, 0xf6, 0x0e, 0x91, 0xe8, 0xce, 0xb4, 0x8a,
	0x67, 0xbd, 0x76, 0xef, 0xfe, 0xd1, 0xaf, 0x92, 0xf5, 0x6c, 0x42, 0xd6, 0x43, 0xd4, 0x9f, 0x51,
	0xd6, 0x07, 0x05, 0x3e, 0x6a, 0x75, 0xe8, 0xab, 0x05, 0x9d, 0x89, 0x7e, 0x23, 0x7f, 0x5a, 0xe9,
	0x8b, 0x17, 0xe3, 0x5e, 0xbf, 0x70, 0xc7, 0xaf, 0x18, 0x4d, 0x4e, 0x05, 0xee, 0x4d, 0x08, 0xdc,
	0x44, 0x4f, 0xff, 0x4e, 0x20, 0x15, 0xf7, 0x71, 0x5d, 0x4f, 0xab, 0xdd, 0x1a, 0x80, 0x17, 0xb3,
	0xec, 0xac, 0x24, 0xe6, 0xf4, 0xa2, 0xcc, 0xad, 0xb6, 0x9a, 0x8a, 0x03, 0xf5, 0x9a, 0x1d, 0x58,
	0xaf, 0x17, 0x2a, 0xeb, 0xb0, 0xa9, 0xdf, 0xb7, 0x07, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x50,
	0x27, 0x9a, 0x7f, 0x68, 0x05, 0x00, 0x00,
}
