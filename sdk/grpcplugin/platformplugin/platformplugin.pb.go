// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platformplugin.proto

/*
Package platformplugin is a generated protocol buffer package.

It is generated from these files:
	platformplugin.proto

It has these top-level messages:
	PlatformPluginManifest
	DeployQuery
	DeployResult
	DeployStatusQuery
*/
package platformplugin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

// Log represents an execution log
// Generate *.pb.go files with:
// 	protoc --go_out=plugins=grpc:. ./log.pb.go
// 	protoc-go-inject-tag -input=./log.pb.go
// 	=> github.com/favadi/protoc-go-inject-tag
type PlatformPluginManifest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version     string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Author      string `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
}

func (m *PlatformPluginManifest) Reset()                    { *m = PlatformPluginManifest{} }
func (m *PlatformPluginManifest) String() string            { return proto.CompactTextString(m) }
func (*PlatformPluginManifest) ProtoMessage()               {}
func (*PlatformPluginManifest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PlatformPluginManifest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlatformPluginManifest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PlatformPluginManifest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PlatformPluginManifest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

type DeployQuery struct {
	Options map[string]string `protobuf:"bytes,1,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DeployQuery) Reset()                    { *m = DeployQuery{} }
func (m *DeployQuery) String() string            { return proto.CompactTextString(m) }
func (*DeployQuery) ProtoMessage()               {}
func (*DeployQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeployQuery) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

type DeployResult struct {
	Status  string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
}

func (m *DeployResult) Reset()                    { *m = DeployResult{} }
func (m *DeployResult) String() string            { return proto.CompactTextString(m) }
func (*DeployResult) ProtoMessage()               {}
func (*DeployResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeployResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeployResult) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type DeployStatusQuery struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *DeployStatusQuery) Reset()                    { *m = DeployStatusQuery{} }
func (m *DeployStatusQuery) String() string            { return proto.CompactTextString(m) }
func (*DeployStatusQuery) ProtoMessage()               {}
func (*DeployStatusQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeployStatusQuery) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*PlatformPluginManifest)(nil), "platformplugin.PlatformPluginManifest")
	proto.RegisterType((*DeployQuery)(nil), "platformplugin.DeployQuery")
	proto.RegisterType((*DeployResult)(nil), "platformplugin.DeployResult")
	proto.RegisterType((*DeployStatusQuery)(nil), "platformplugin.DeployStatusQuery")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PlatformPlugin service

type PlatformPluginClient interface {
	Manifest(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PlatformPluginManifest, error)
	Deploy(ctx context.Context, in *DeployQuery, opts ...grpc.CallOption) (*DeployResult, error)
	DeployStatus(ctx context.Context, in *DeployStatusQuery, opts ...grpc.CallOption) (*DeployResult, error)
}

type platformPluginClient struct {
	cc *grpc.ClientConn
}

func NewPlatformPluginClient(cc *grpc.ClientConn) PlatformPluginClient {
	return &platformPluginClient{cc}
}

func (c *platformPluginClient) Manifest(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*PlatformPluginManifest, error) {
	out := new(PlatformPluginManifest)
	err := grpc.Invoke(ctx, "/platformplugin.PlatformPlugin/Manifest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformPluginClient) Deploy(ctx context.Context, in *DeployQuery, opts ...grpc.CallOption) (*DeployResult, error) {
	out := new(DeployResult)
	err := grpc.Invoke(ctx, "/platformplugin.PlatformPlugin/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *platformPluginClient) DeployStatus(ctx context.Context, in *DeployStatusQuery, opts ...grpc.CallOption) (*DeployResult, error) {
	out := new(DeployResult)
	err := grpc.Invoke(ctx, "/platformplugin.PlatformPlugin/DeployStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlatformPlugin service

type PlatformPluginServer interface {
	Manifest(context.Context, *google_protobuf.Empty) (*PlatformPluginManifest, error)
	Deploy(context.Context, *DeployQuery) (*DeployResult, error)
	DeployStatus(context.Context, *DeployStatusQuery) (*DeployResult, error)
}

func RegisterPlatformPluginServer(s *grpc.Server, srv PlatformPluginServer) {
	s.RegisterService(&_PlatformPlugin_serviceDesc, srv)
}

func _PlatformPlugin_Manifest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformPluginServer).Manifest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformplugin.PlatformPlugin/Manifest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformPluginServer).Manifest(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformPlugin_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformPluginServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformplugin.PlatformPlugin/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformPluginServer).Deploy(ctx, req.(*DeployQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlatformPlugin_DeployStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployStatusQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlatformPluginServer).DeployStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platformplugin.PlatformPlugin/DeployStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlatformPluginServer).DeployStatus(ctx, req.(*DeployStatusQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlatformPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "platformplugin.PlatformPlugin",
	HandlerType: (*PlatformPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Manifest",
			Handler:    _PlatformPlugin_Manifest_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _PlatformPlugin_Deploy_Handler,
		},
		{
			MethodName: "DeployStatus",
			Handler:    _PlatformPlugin_DeployStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "platformplugin.proto",
}

func init() { proto.RegisterFile("platformplugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0xef, 0xd2, 0x30,
	0x14, 0xfe, 0x75, 0x20, 0xe8, 0x83, 0x10, 0xad, 0x84, 0x2c, 0xc3, 0x03, 0xce, 0xc4, 0xe0, 0xa5,
	0x24, 0x78, 0x31, 0x9c, 0x0c, 0x81, 0x03, 0x89, 0xc6, 0x31, 0x6f, 0xde, 0xc6, 0x56, 0xc6, 0xc2,
	0xb6, 0x2e, 0x6d, 0x47, 0xb2, 0x9b, 0xff, 0x80, 0xff, 0xaf, 0x47, 0xb3, 0xb6, 0x33, 0x1b, 0x21,
	0x7a, 0xeb, 0xf7, 0xbe, 0xef, 0xbd, 0xf6, 0xfb, 0x5e, 0x61, 0x5a, 0xa4, 0x81, 0x3c, 0x33, 0x9e,
	0x15, 0x69, 0x19, 0x27, 0x39, 0x29, 0x38, 0x93, 0x0c, 0x4f, 0xba, 0x55, 0x67, 0x1e, 0x33, 0x16,
	0xa7, 0x74, 0xa5, 0xd8, 0x53, 0x79, 0x5e, 0xd1, 0xac, 0x90, 0x95, 0x16, 0xbb, 0x3f, 0x11, 0xcc,
	0x3c, 0xa3, 0xf7, 0x94, 0xfe, 0x6b, 0x90, 0x27, 0x67, 0x2a, 0x24, 0xc6, 0xd0, 0xcf, 0x83, 0x8c,
	0xda, 0x68, 0x81, 0x96, 0x2f, 0x7c, 0x75, 0xc6, 0x36, 0x0c, 0x6f, 0x94, 0x8b, 0x84, 0xe5, 0xb6,
	0xa5, 0xca, 0x0d, 0xc4, 0x0b, 0x18, 0x45, 0x54, 0x84, 0x3c, 0x29, 0x64, 0xcd, 0xf6, 0x14, 0xdb,
	0x2e, 0xe1, 0x19, 0x0c, 0x82, 0x52, 0x5e, 0x18, 0xb7, 0xfb, 0x8a, 0x34, 0xc8, 0xfd, 0x85, 0x60,
	0xb4, 0xa3, 0x45, 0xca, 0xaa, 0x63, 0x49, 0x79, 0x85, 0xb7, 0x30, 0x64, 0xaa, 0x43, 0xd8, 0x68,
	0xd1, 0x5b, 0x8e, 0xd6, 0x4b, 0x72, 0xe7, 0xb3, 0xa5, 0x26, 0xdf, 0xb4, 0x74, 0x9f, 0x4b, 0x5e,
	0xf9, 0x4d, 0xa3, 0xb3, 0x81, 0x71, 0x9b, 0xc0, 0x2f, 0xa1, 0x77, 0xa5, 0x95, 0xb1, 0x52, 0x1f,
	0xf1, 0x14, 0x9e, 0xdd, 0x82, 0xb4, 0xa4, 0xc6, 0x87, 0x06, 0x1b, 0xeb, 0x13, 0x72, 0x3f, 0xc3,
	0x58, 0x5f, 0xe0, 0x53, 0x51, 0xa6, 0xb2, 0x7e, 0xb7, 0x90, 0x81, 0x2c, 0x85, 0x69, 0x37, 0xa8,
	0xce, 0x22, 0xa2, 0x32, 0x48, 0x52, 0xd1, 0x64, 0x61, 0xa0, 0xfb, 0x0e, 0x5e, 0xe9, 0x09, 0xdf,
	0x95, 0x52, 0xdb, 0x9a, 0x80, 0x75, 0xd8, 0x99, 0x11, 0xd6, 0x61, 0xb7, 0xfe, 0x8d, 0x60, 0xd2,
	0x4d, 0x1e, 0x7f, 0x81, 0xe7, 0x7f, 0xd3, 0x9f, 0x11, 0xbd, 0x36, 0xd2, 0xac, 0x8d, 0xec, 0xeb,
	0xb5, 0x39, 0xef, 0xef, 0xc3, 0x78, 0xbc, 0x3d, 0xf7, 0x09, 0xef, 0x61, 0xa0, 0x5f, 0x81, 0xe7,
	0xff, 0x08, 0xd0, 0x79, 0xf3, 0x98, 0xd4, 0xe6, 0xdd, 0x27, 0x7c, 0x6c, 0xe2, 0xd0, 0x66, 0xf0,
	0xdb, 0xc7, 0xfa, 0x96, 0xd5, 0xff, 0x8d, 0xdc, 0xfa, 0xf0, 0x21, 0x64, 0x19, 0x61, 0xb7, 0x0b,
	0x09, 0x23, 0x41, 0x44, 0x74, 0x25, 0x31, 0x2f, 0xc2, 0xe6, 0x17, 0x77, 0x7a, 0xb7, 0xaf, 0xbb,
	0x06, 0xbd, 0x3a, 0x0e, 0x0f, 0xfd, 0xb8, 0xfb, 0xe5, 0xa7, 0x81, 0xca, 0xe9, 0xe3, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x46, 0x66, 0x07, 0xe7, 0x14, 0x03, 0x00, 0x00,
}
