// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party/tink/proto/hmac_prf.proto

package hmac_prf_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common_go_proto "github.com/prizraksarvar/tink/proto/common_go_proto"
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

type HmacPrfParams struct {
	Hash                 common_go_proto.HashType `protobuf:"varint,1,opt,name=hash,proto3,enum=google.crypto.tink.HashType" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *HmacPrfParams) Reset()         { *m = HmacPrfParams{} }
func (m *HmacPrfParams) String() string { return proto.CompactTextString(m) }
func (*HmacPrfParams) ProtoMessage()    {}
func (*HmacPrfParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_547bbc464bb1c2f6, []int{0}
}

func (m *HmacPrfParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HmacPrfParams.Unmarshal(m, b)
}
func (m *HmacPrfParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HmacPrfParams.Marshal(b, m, deterministic)
}
func (m *HmacPrfParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HmacPrfParams.Merge(m, src)
}
func (m *HmacPrfParams) XXX_Size() int {
	return xxx_messageInfo_HmacPrfParams.Size(m)
}
func (m *HmacPrfParams) XXX_DiscardUnknown() {
	xxx_messageInfo_HmacPrfParams.DiscardUnknown(m)
}

var xxx_messageInfo_HmacPrfParams proto.InternalMessageInfo

func (m *HmacPrfParams) GetHash() common_go_proto.HashType {
	if m != nil {
		return m.Hash
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

// key_type: type.googleapis.com/google.crypto.tink.HmacPrfKey
type HmacPrfKey struct {
	Version              uint32         `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params               *HmacPrfParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KeyValue             []byte         `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HmacPrfKey) Reset()         { *m = HmacPrfKey{} }
func (m *HmacPrfKey) String() string { return proto.CompactTextString(m) }
func (*HmacPrfKey) ProtoMessage()    {}
func (*HmacPrfKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_547bbc464bb1c2f6, []int{1}
}

func (m *HmacPrfKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HmacPrfKey.Unmarshal(m, b)
}
func (m *HmacPrfKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HmacPrfKey.Marshal(b, m, deterministic)
}
func (m *HmacPrfKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HmacPrfKey.Merge(m, src)
}
func (m *HmacPrfKey) XXX_Size() int {
	return xxx_messageInfo_HmacPrfKey.Size(m)
}
func (m *HmacPrfKey) XXX_DiscardUnknown() {
	xxx_messageInfo_HmacPrfKey.DiscardUnknown(m)
}

var xxx_messageInfo_HmacPrfKey proto.InternalMessageInfo

func (m *HmacPrfKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HmacPrfKey) GetParams() *HmacPrfParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *HmacPrfKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

type HmacPrfKeyFormat struct {
	Params               *HmacPrfParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	KeySize              uint32         `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	Version              uint32         `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HmacPrfKeyFormat) Reset()         { *m = HmacPrfKeyFormat{} }
func (m *HmacPrfKeyFormat) String() string { return proto.CompactTextString(m) }
func (*HmacPrfKeyFormat) ProtoMessage()    {}
func (*HmacPrfKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_547bbc464bb1c2f6, []int{2}
}

func (m *HmacPrfKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HmacPrfKeyFormat.Unmarshal(m, b)
}
func (m *HmacPrfKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HmacPrfKeyFormat.Marshal(b, m, deterministic)
}
func (m *HmacPrfKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HmacPrfKeyFormat.Merge(m, src)
}
func (m *HmacPrfKeyFormat) XXX_Size() int {
	return xxx_messageInfo_HmacPrfKeyFormat.Size(m)
}
func (m *HmacPrfKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_HmacPrfKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_HmacPrfKeyFormat proto.InternalMessageInfo

func (m *HmacPrfKeyFormat) GetParams() *HmacPrfParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *HmacPrfKeyFormat) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

func (m *HmacPrfKeyFormat) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*HmacPrfParams)(nil), "google.crypto.tink.HmacPrfParams")
	proto.RegisterType((*HmacPrfKey)(nil), "google.crypto.tink.HmacPrfKey")
	proto.RegisterType((*HmacPrfKeyFormat)(nil), "google.crypto.tink.HmacPrfKeyFormat")
}

func init() {
	proto.RegisterFile("proto/hmac_prf.proto", fileDescriptor_547bbc464bb1c2f6)
}

var fileDescriptor_547bbc464bb1c2f6 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0x33, 0x31,
	0x18, 0x84, 0x49, 0xfb, 0xd1, 0x7e, 0x46, 0x2b, 0xb2, 0xa7, 0x55, 0x7b, 0xa8, 0x15, 0xa1, 0xa7,
	0xac, 0xd4, 0x93, 0x47, 0x7b, 0x90, 0x4a, 0x41, 0x96, 0xb5, 0x08, 0x7a, 0x09, 0x69, 0x4c, 0x37,
	0x61, 0x9b, 0xbe, 0x21, 0x9b, 0x16, 0xd2, 0x83, 0x07, 0x7f, 0x8a, 0xbf, 0x54, 0x9a, 0x56, 0xac,
	0xb8, 0x1e, 0x3c, 0xce, 0x32, 0xf3, 0xcc, 0xbc, 0x1b, 0x7c, 0xe1, 0xa4, 0xb2, 0x2f, 0xd4, 0x30,
	0xeb, 0x7c, 0xe2, 0xd4, 0xbc, 0x48, 0x8c, 0x05, 0x07, 0x89, 0xd4, 0x8c, 0x53, 0x63, 0xa7, 0x24,
	0xc8, 0x28, 0xca, 0x01, 0xf2, 0x99, 0x20, 0xdc, 0x7a, 0xe3, 0x80, 0xac, 0x8d, 0x27, 0xe7, 0xbf,
	0x44, 0x39, 0x68, 0x0d, 0xf3, 0x4d, 0xb0, 0x7b, 0x83, 0x5b, 0x43, 0xcd, 0x78, 0x6a, 0xa7, 0x29,
	0xb3, 0x4c, 0x97, 0xd1, 0x25, 0xfe, 0x27, 0x59, 0x29, 0x63, 0xd4, 0x41, 0xbd, 0xc3, 0x7e, 0x9b,
	0xfc, 0x04, 0x93, 0x21, 0x2b, 0xe5, 0xd8, 0x1b, 0x91, 0x05, 0x67, 0xf7, 0x15, 0xe3, 0x2d, 0x62,
	0x24, 0x7c, 0x14, 0xe3, 0xe6, 0x52, 0xd8, 0x52, 0xc1, 0x3c, 0x20, 0x5a, 0xd9, 0xa7, 0x8c, 0xae,
	0x71, 0xc3, 0x84, 0x8e, 0xb8, 0xd6, 0x41, 0xbd, 0xfd, 0xfe, 0x59, 0x25, 0x7b, 0x77, 0x4c, 0xb6,
	0x0d, 0x44, 0xa7, 0x78, 0xaf, 0x10, 0x9e, 0x2e, 0xd9, 0x6c, 0x21, 0xe2, 0x7a, 0x07, 0xf5, 0x0e,
	0xb2, 0xff, 0x85, 0xf0, 0x8f, 0x6b, 0xdd, 0x7d, 0x43, 0xf8, 0xe8, 0x6b, 0xc0, 0x2d, 0x58, 0xcd,
	0xdc, 0x4e, 0x19, 0xfa, 0x6b, 0xd9, 0x31, 0x5e, 0xb3, 0x69, 0xa9, 0x56, 0x22, 0x2c, 0x6d, 0x65,
	0xcd, 0x42, 0xf8, 0x07, 0xb5, 0x12, 0xbb, 0xc7, 0xd5, 0xbf, 0x1d, 0x37, 0x78, 0xc2, 0x6d, 0x0e,
	0xba, 0xaa, 0x24, 0xfc, 0xe7, 0x14, 0x3d, 0x93, 0x5c, 0x39, 0xb9, 0x98, 0x10, 0x0e, 0x3a, 0xd9,
	0xd8, 0xaa, 0xde, 0x93, 0xe6, 0x40, 0xc3, 0x97, 0xf7, 0x5a, 0x63, 0x7c, 0x77, 0x3f, 0x4a, 0x07,
	0x93, 0x46, 0xd0, 0x57, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xe3, 0x29, 0x03, 0x0b, 0x02,
	0x00, 0x00,
}
