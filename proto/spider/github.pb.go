// Code generated by protoc-gen-go.
// source: github.com/daviddengcn/gcse/proto/spider/github.proto
// DO NOT EDIT!

/*
Package gcse_spider is a generated protocol buffer package.

It is generated from these files:
	github.com/daviddengcn/gcse/proto/spider/github.proto

It has these top-level messages:
	FolderInfo
*/
package gcse_spider

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type FolderInfo struct {
	// E.g. "spider"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// E.g. "spider/github"
	Path    string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Sha     string `protobuf:"bytes,3,opt,name=sha" json:"sha,omitempty"`
	HtmlUrl string `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
}

func (m *FolderInfo) Reset()                    { *m = FolderInfo{} }
func (m *FolderInfo) String() string            { return proto.CompactTextString(m) }
func (*FolderInfo) ProtoMessage()               {}
func (*FolderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*FolderInfo)(nil), "gcse.spider.FolderInfo")
}

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x49, 0x2c, 0xcb, 0x4c, 0x49, 0x49, 0xcd, 0x4b,
	0x4f, 0xce, 0xd3, 0x4f, 0x4f, 0x2e, 0x4e, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e,
	0xc8, 0x4c, 0x49, 0x2d, 0xd2, 0x87, 0x2a, 0x04, 0x8b, 0x09, 0x71, 0x83, 0xe4, 0xf5, 0x20, 0x32,
	0x4a, 0x89, 0x5c, 0x5c, 0x6e, 0xf9, 0x39, 0x40, 0x96, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x10, 0x17,
	0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x0d, 0x12, 0x2b,
	0x48, 0x2c, 0xc9, 0x90, 0x60, 0x82, 0x88, 0x81, 0xd8, 0x42, 0x02, 0x5c, 0xcc, 0xc5, 0x19, 0x89,
	0x12, 0xcc, 0x60, 0x21, 0x10, 0x53, 0x48, 0x92, 0x8b, 0x23, 0xa3, 0x24, 0x37, 0x27, 0xbe, 0xb4,
	0x28, 0x47, 0x82, 0x05, 0x2c, 0xcc, 0x0e, 0xe2, 0x87, 0x16, 0xe5, 0x24, 0xb1, 0x81, 0xad, 0x35,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xd1, 0x93, 0xed, 0xaf, 0x00, 0x00, 0x00,
}
