// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: functional_base.proto

/*
Package translator is a generated protocol buffer package.

It is generated from these files:
	functional_base.proto

It has these top-level messages:
	FunctionalFilterRouteConfig
*/
package translator

import (
	fmt "fmt"

	proto "github.com/gogo/protobuf/proto"

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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FunctionalFilterRouteConfig struct {
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName,proto3" json:"function_name,omitempty"`
}

func (m *FunctionalFilterRouteConfig) Reset()         { *m = FunctionalFilterRouteConfig{} }
func (m *FunctionalFilterRouteConfig) String() string { return proto.CompactTextString(m) }
func (*FunctionalFilterRouteConfig) ProtoMessage()    {}
func (*FunctionalFilterRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorFunctionalBase, []int{0}
}

func (m *FunctionalFilterRouteConfig) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionalFilterRouteConfig)(nil), "envoy.api.v2.filter.http.FunctionalFilterRouteConfig")
}

func init() { proto.RegisterFile("functional_base.proto", fileDescriptorFunctionalBase) }

var fileDescriptorFunctionalBase = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2b, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0x4b, 0xcc, 0x89, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x48, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0xd2,
	0x4b, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xcb, 0x28, 0x29, 0x29, 0x50, 0x72, 0xe2, 0x92, 0x76,
	0x83, 0x6b, 0x71, 0x03, 0x4b, 0x04, 0xe5, 0x97, 0x96, 0xa4, 0x3a, 0xe7, 0xe7, 0xa5, 0x65, 0xa6,
	0x0b, 0x29, 0x73, 0xf1, 0xc2, 0x4c, 0x8c, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0xe2, 0x81, 0x09, 0xfa, 0x25, 0xe6, 0xa6, 0x3a, 0xf1, 0x44, 0x71, 0x95, 0x14, 0x25,
	0xe6, 0x15, 0xe7, 0x24, 0x96, 0xe4, 0x17, 0x25, 0xb1, 0x81, 0xad, 0x34, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x44, 0x28, 0x4d, 0x42, 0x8b, 0x00, 0x00, 0x00,
}
