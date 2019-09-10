// Code generated by protoc-gen-go.
// source: signer.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	signer.proto

It has these top-level messages:
	CreateKeyRequest
	KeyInfo
	KeyID
	Algorithm
	GetKeyInfoResponse
	PublicKey
	Signature
	SignatureRequest
	Void
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateKeyRequest struct {
	Algorithm string `protobuf:"bytes,1,opt,name=algorithm" json:"algorithm,omitempty"`
	Gun       string `protobuf:"bytes,2,opt,name=gun" json:"gun,omitempty"`
	Role      string `protobuf:"bytes,3,opt,name=role" json:"role,omitempty"`
}

func (m *CreateKeyRequest) Reset()                    { *m = CreateKeyRequest{} }
func (m *CreateKeyRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateKeyRequest) ProtoMessage()               {}
func (*CreateKeyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateKeyRequest) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

func (m *CreateKeyRequest) GetGun() string {
	if m != nil {
		return m.Gun
	}
	return ""
}

func (m *CreateKeyRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

// KeyInfo holds a KeyID that is used to reference the key and it's algorithm
type KeyInfo struct {
	KeyID     *KeyID     `protobuf:"bytes,1,opt,name=keyID" json:"keyID,omitempty"`
	Algorithm *Algorithm `protobuf:"bytes,2,opt,name=algorithm" json:"algorithm,omitempty"`
}

func (m *KeyInfo) Reset()                    { *m = KeyInfo{} }
func (m *KeyInfo) String() string            { return proto1.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()               {}
func (*KeyInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeyInfo) GetKeyID() *KeyID {
	if m != nil {
		return m.KeyID
	}
	return nil
}

func (m *KeyInfo) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

// KeyID holds an ID that is used to reference the key
type KeyID struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *KeyID) Reset()                    { *m = KeyID{} }
func (m *KeyID) String() string            { return proto1.CompactTextString(m) }
func (*KeyID) ProtoMessage()               {}
func (*KeyID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

// Type holds the type of crypto algorithm used
type Algorithm struct {
	Algorithm string `protobuf:"bytes,1,opt,name=algorithm" json:"algorithm,omitempty"`
}

func (m *Algorithm) Reset()                    { *m = Algorithm{} }
func (m *Algorithm) String() string            { return proto1.CompactTextString(m) }
func (*Algorithm) ProtoMessage()               {}
func (*Algorithm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Algorithm) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

// GetKeyInfoResponse returns the public key, the role, and the algorithm and key ID.
// For backwards compatibility, it doesn't embed a PublicKey object
type GetKeyInfoResponse struct {
	KeyInfo   *KeyInfo `protobuf:"bytes,1,opt,name=keyInfo" json:"keyInfo,omitempty"`
	PublicKey []byte   `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Role      string   `protobuf:"bytes,3,opt,name=role" json:"role,omitempty"`
}

func (m *GetKeyInfoResponse) Reset()                    { *m = GetKeyInfoResponse{} }
func (m *GetKeyInfoResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetKeyInfoResponse) ProtoMessage()               {}
func (*GetKeyInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetKeyInfoResponse) GetKeyInfo() *KeyInfo {
	if m != nil {
		return m.KeyInfo
	}
	return nil
}

func (m *GetKeyInfoResponse) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *GetKeyInfoResponse) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

// PublicKey has a KeyInfo that is used to reference the key, and opaque bytes of a publicKey
type PublicKey struct {
	KeyInfo   *KeyInfo `protobuf:"bytes,1,opt,name=keyInfo" json:"keyInfo,omitempty"`
	PublicKey []byte   `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto1.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PublicKey) GetKeyInfo() *KeyInfo {
	if m != nil {
		return m.KeyInfo
	}
	return nil
}

func (m *PublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

// Signature specifies a KeyInfo that was used for signing and signed content
type Signature struct {
	KeyInfo   *KeyInfo   `protobuf:"bytes,1,opt,name=keyInfo" json:"keyInfo,omitempty"`
	Algorithm *Algorithm `protobuf:"bytes,2,opt,name=algorithm" json:"algorithm,omitempty"`
	Content   []byte     `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto1.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Signature) GetKeyInfo() *KeyInfo {
	if m != nil {
		return m.KeyInfo
	}
	return nil
}

func (m *Signature) GetAlgorithm() *Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return nil
}

func (m *Signature) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// SignatureRequests specifies a KeyInfo, and content to be signed
type SignatureRequest struct {
	KeyID   *KeyID `protobuf:"bytes,1,opt,name=keyID" json:"keyID,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *SignatureRequest) Reset()                    { *m = SignatureRequest{} }
func (m *SignatureRequest) String() string            { return proto1.CompactTextString(m) }
func (*SignatureRequest) ProtoMessage()               {}
func (*SignatureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SignatureRequest) GetKeyID() *KeyID {
	if m != nil {
		return m.KeyID
	}
	return nil
}

func (m *SignatureRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// Void represents an empty message type
type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto1.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto1.RegisterType((*CreateKeyRequest)(nil), "proto.CreateKeyRequest")
	proto1.RegisterType((*KeyInfo)(nil), "proto.KeyInfo")
	proto1.RegisterType((*KeyID)(nil), "proto.KeyID")
	proto1.RegisterType((*Algorithm)(nil), "proto.Algorithm")
	proto1.RegisterType((*GetKeyInfoResponse)(nil), "proto.GetKeyInfoResponse")
	proto1.RegisterType((*PublicKey)(nil), "proto.PublicKey")
	proto1.RegisterType((*Signature)(nil), "proto.Signature")
	proto1.RegisterType((*SignatureRequest)(nil), "proto.SignatureRequest")
	proto1.RegisterType((*Void)(nil), "proto.Void")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KeyManagement service

type KeyManagementClient interface {
	// CreateKey creates as asymmetric key pair and returns the PublicKey
	CreateKey(ctx context.Context, in *CreateKeyRequest, opts ...grpc.CallOption) (*PublicKey, error)
	// DeleteKey deletes the key associated with a KeyID
	DeleteKey(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*Void, error)
	// GetKeyInfo returns the PublicKey associated with a KeyID
	GetKeyInfo(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*GetKeyInfoResponse, error)
}

type keyManagementClient struct {
	cc *grpc.ClientConn
}

func NewKeyManagementClient(cc *grpc.ClientConn) KeyManagementClient {
	return &keyManagementClient{cc}
}

func (c *keyManagementClient) CreateKey(ctx context.Context, in *CreateKeyRequest, opts ...grpc.CallOption) (*PublicKey, error) {
	out := new(PublicKey)
	err := grpc.Invoke(ctx, "/proto.KeyManagement/CreateKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementClient) DeleteKey(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/proto.KeyManagement/DeleteKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementClient) GetKeyInfo(ctx context.Context, in *KeyID, opts ...grpc.CallOption) (*GetKeyInfoResponse, error) {
	out := new(GetKeyInfoResponse)
	err := grpc.Invoke(ctx, "/proto.KeyManagement/GetKeyInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyManagement service

type KeyManagementServer interface {
	// CreateKey creates as asymmetric key pair and returns the PublicKey
	CreateKey(context.Context, *CreateKeyRequest) (*PublicKey, error)
	// DeleteKey deletes the key associated with a KeyID
	DeleteKey(context.Context, *KeyID) (*Void, error)
	// GetKeyInfo returns the PublicKey associated with a KeyID
	GetKeyInfo(context.Context, *KeyID) (*GetKeyInfoResponse, error)
}

func RegisterKeyManagementServer(s *grpc.Server, srv KeyManagementServer) {
	s.RegisterService(&_KeyManagement_serviceDesc, srv)
}

func _KeyManagement_CreateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServer).CreateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyManagement/CreateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServer).CreateKey(ctx, req.(*CreateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagement_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyManagement/DeleteKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServer).DeleteKey(ctx, req.(*KeyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagement_GetKeyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServer).GetKeyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KeyManagement/GetKeyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServer).GetKeyInfo(ctx, req.(*KeyID))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KeyManagement",
	HandlerType: (*KeyManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKey",
			Handler:    _KeyManagement_CreateKey_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _KeyManagement_DeleteKey_Handler,
		},
		{
			MethodName: "GetKeyInfo",
			Handler:    _KeyManagement_GetKeyInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signer.proto",
}

// Client API for Signer service

type SignerClient interface {
	// Sign calculates a cryptographic signature using the Key associated with a KeyID and returns the signature
	Sign(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*Signature, error)
}

type signerClient struct {
	cc *grpc.ClientConn
}

func NewSignerClient(cc *grpc.ClientConn) SignerClient {
	return &signerClient{cc}
}

func (c *signerClient) Sign(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*Signature, error) {
	out := new(Signature)
	err := grpc.Invoke(ctx, "/proto.Signer/Sign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Signer service

type SignerServer interface {
	// Sign calculates a cryptographic signature using the Key associated with a KeyID and returns the signature
	Sign(context.Context, *SignatureRequest) (*Signature, error)
}

func RegisterSignerServer(s *grpc.Server, srv SignerServer) {
	s.RegisterService(&_Signer_serviceDesc, srv)
}

func _Signer_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignerServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Signer/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignerServer).Sign(ctx, req.(*SignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Signer",
	HandlerType: (*SignerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _Signer_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "signer.proto",
}

func init() { proto1.RegisterFile("signer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0x3f, 0xef, 0xd3, 0x40,
	0x0c, 0x4d, 0xf3, 0xeb, 0x1f, 0x9d, 0x1b, 0xaa, 0xc8, 0x4b, 0x43, 0xc5, 0x80, 0x6e, 0x2a, 0x4b,
	0x87, 0x76, 0x80, 0x85, 0x01, 0x11, 0x09, 0x55, 0x11, 0x52, 0x95, 0x4a, 0xdd, 0x18, 0xd2, 0x62,
	0x42, 0xd4, 0xf4, 0x2e, 0x24, 0x97, 0x21, 0x13, 0x5f, 0x88, 0x0f, 0x89, 0x72, 0x49, 0xae, 0x69,
	0x41, 0x88, 0x4a, 0xbf, 0x29, 0xf6, 0xb3, 0xfd, 0xfc, 0xec, 0xf8, 0xc0, 0x29, 0x92, 0x58, 0x50,
	0xbe, 0xca, 0x72, 0xa9, 0x24, 0x8e, 0xf4, 0x87, 0x1f, 0xc0, 0xfd, 0x98, 0x53, 0xa4, 0x28, 0xa0,
	0x2a, 0xa4, 0x1f, 0x25, 0x15, 0x0a, 0x5f, 0x01, 0x8b, 0xd2, 0x58, 0xe6, 0x89, 0xfa, 0x7e, 0xf1,
	0x06, 0xaf, 0x07, 0x4b, 0x16, 0x5e, 0x01, 0x74, 0xe1, 0x29, 0x2e, 0x85, 0x67, 0x6b, 0xbc, 0x36,
	0x11, 0x61, 0x98, 0xcb, 0x94, 0xbc, 0x27, 0x0d, 0x69, 0x9b, 0x7f, 0x81, 0x49, 0x40, 0xd5, 0x56,
	0x7c, 0x93, 0xc8, 0x61, 0x74, 0xa6, 0x6a, 0xeb, 0x6b, 0xaa, 0xe9, 0xda, 0x69, 0x04, 0xac, 0xea,
	0xb0, 0x1f, 0x36, 0x21, 0x5c, 0xf5, 0x5b, 0xda, 0x3a, 0xcf, 0x6d, 0xf3, 0x3e, 0x74, 0x78, 0x4f,
	0x04, 0x9f, 0xc3, 0x48, 0xd7, 0xe3, 0x0c, 0xec, 0x96, 0x99, 0x85, 0xf6, 0xd6, 0xe7, 0x6f, 0x80,
	0x99, 0x82, 0x7f, 0x0f, 0xc2, 0x33, 0xc0, 0x4f, 0xa4, 0x5a, 0x95, 0x21, 0x15, 0x99, 0x14, 0x05,
	0xe1, 0x12, 0x26, 0xe7, 0x06, 0x6a, 0xf5, 0xce, 0x7a, 0x7a, 0xeb, 0xc4, 0x2e, 0x5c, 0xb3, 0x67,
	0xe5, 0x31, 0x4d, 0x4e, 0x01, 0x55, 0x5a, 0xb3, 0x13, 0x5e, 0x81, 0xbf, 0x2e, 0x65, 0x0f, 0x6c,
	0x67, 0x12, 0x9e, 0xa9, 0x11, 0xff, 0x09, 0x6c, 0x9f, 0xc4, 0x22, 0x52, 0x65, 0xfe, 0x88, 0xfa,
	0x07, 0x37, 0x8e, 0x1e, 0x4c, 0x4e, 0x52, 0x28, 0x12, 0x4a, 0x8f, 0xe4, 0x84, 0x9d, 0xcb, 0x77,
	0xe0, 0x1a, 0x01, 0xdd, 0x09, 0xfd, 0xcf, 0x3f, 0xef, 0x31, 0xda, 0xb7, 0x8c, 0x63, 0x18, 0x1e,
	0x64, 0xf2, 0x75, 0xfd, 0x6b, 0x00, 0x2f, 0x02, 0xaa, 0x3e, 0x47, 0x22, 0x8a, 0xe9, 0x42, 0x42,
	0xe1, 0x3b, 0x60, 0xe6, 0x5c, 0x71, 0xde, 0xb2, 0xde, 0x1f, 0xf0, 0xa2, 0x1b, 0xc4, 0x2c, 0x9b,
	0x5b, 0xb8, 0x04, 0xe6, 0x53, 0x4a, 0x4d, 0xe5, 0x8d, 0x9e, 0xc5, 0xb4, 0xf5, 0xea, 0x9e, 0xdc,
	0xc2, 0xb7, 0x00, 0xd7, 0xbb, 0xb8, 0x4b, 0x7d, 0xd9, 0x7a, 0x7f, 0x1e, 0x0e, 0xb7, 0xd6, 0xef,
	0x61, 0xbc, 0xd7, 0x4f, 0x0c, 0x37, 0x30, 0xac, 0x2d, 0xa3, 0xf0, 0x7e, 0x3f, 0x46, 0xa1, 0x09,
	0x70, 0xeb, 0x38, 0xd6, 0xd0, 0xe6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x92, 0x2c, 0xe0,
	0xa8, 0x03, 0x00, 0x00,
}