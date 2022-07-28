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
// source: third_party/tink/proto/aes_gcm_hkdf_streaming.proto

package aes_gcm_hkdf_streaming_go_proto

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

type AesGcmHkdfStreamingParams struct {
	CiphertextSegmentSize uint32                   `protobuf:"varint,1,opt,name=ciphertext_segment_size,json=ciphertextSegmentSize,proto3" json:"ciphertext_segment_size,omitempty"`
	DerivedKeySize        uint32                   `protobuf:"varint,2,opt,name=derived_key_size,json=derivedKeySize,proto3" json:"derived_key_size,omitempty"`
	HkdfHashType          common_go_proto.HashType `protobuf:"varint,3,opt,name=hkdf_hash_type,json=hkdfHashType,proto3,enum=google.crypto.tink.HashType" json:"hkdf_hash_type,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *AesGcmHkdfStreamingParams) Reset()         { *m = AesGcmHkdfStreamingParams{} }
func (m *AesGcmHkdfStreamingParams) String() string { return proto.CompactTextString(m) }
func (*AesGcmHkdfStreamingParams) ProtoMessage()    {}
func (*AesGcmHkdfStreamingParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dba2d882aaf5933, []int{0}
}

func (m *AesGcmHkdfStreamingParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AesGcmHkdfStreamingParams.Unmarshal(m, b)
}
func (m *AesGcmHkdfStreamingParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AesGcmHkdfStreamingParams.Marshal(b, m, deterministic)
}
func (m *AesGcmHkdfStreamingParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AesGcmHkdfStreamingParams.Merge(m, src)
}
func (m *AesGcmHkdfStreamingParams) XXX_Size() int {
	return xxx_messageInfo_AesGcmHkdfStreamingParams.Size(m)
}
func (m *AesGcmHkdfStreamingParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AesGcmHkdfStreamingParams.DiscardUnknown(m)
}

var xxx_messageInfo_AesGcmHkdfStreamingParams proto.InternalMessageInfo

func (m *AesGcmHkdfStreamingParams) GetCiphertextSegmentSize() uint32 {
	if m != nil {
		return m.CiphertextSegmentSize
	}
	return 0
}

func (m *AesGcmHkdfStreamingParams) GetDerivedKeySize() uint32 {
	if m != nil {
		return m.DerivedKeySize
	}
	return 0
}

func (m *AesGcmHkdfStreamingParams) GetHkdfHashType() common_go_proto.HashType {
	if m != nil {
		return m.HkdfHashType
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

type AesGcmHkdfStreamingKeyFormat struct {
	Version              uint32                     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Params               *AesGcmHkdfStreamingParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	KeySize              uint32                     `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AesGcmHkdfStreamingKeyFormat) Reset()         { *m = AesGcmHkdfStreamingKeyFormat{} }
func (m *AesGcmHkdfStreamingKeyFormat) String() string { return proto.CompactTextString(m) }
func (*AesGcmHkdfStreamingKeyFormat) ProtoMessage()    {}
func (*AesGcmHkdfStreamingKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dba2d882aaf5933, []int{1}
}

func (m *AesGcmHkdfStreamingKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AesGcmHkdfStreamingKeyFormat.Unmarshal(m, b)
}
func (m *AesGcmHkdfStreamingKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AesGcmHkdfStreamingKeyFormat.Marshal(b, m, deterministic)
}
func (m *AesGcmHkdfStreamingKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AesGcmHkdfStreamingKeyFormat.Merge(m, src)
}
func (m *AesGcmHkdfStreamingKeyFormat) XXX_Size() int {
	return xxx_messageInfo_AesGcmHkdfStreamingKeyFormat.Size(m)
}
func (m *AesGcmHkdfStreamingKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_AesGcmHkdfStreamingKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_AesGcmHkdfStreamingKeyFormat proto.InternalMessageInfo

func (m *AesGcmHkdfStreamingKeyFormat) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *AesGcmHkdfStreamingKeyFormat) GetParams() *AesGcmHkdfStreamingParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *AesGcmHkdfStreamingKeyFormat) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.AesGcmHkdfStreamingKey
type AesGcmHkdfStreamingKey struct {
	Version              uint32                     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params               *AesGcmHkdfStreamingParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KeyValue             []byte                     `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AesGcmHkdfStreamingKey) Reset()         { *m = AesGcmHkdfStreamingKey{} }
func (m *AesGcmHkdfStreamingKey) String() string { return proto.CompactTextString(m) }
func (*AesGcmHkdfStreamingKey) ProtoMessage()    {}
func (*AesGcmHkdfStreamingKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_1dba2d882aaf5933, []int{2}
}

func (m *AesGcmHkdfStreamingKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AesGcmHkdfStreamingKey.Unmarshal(m, b)
}
func (m *AesGcmHkdfStreamingKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AesGcmHkdfStreamingKey.Marshal(b, m, deterministic)
}
func (m *AesGcmHkdfStreamingKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AesGcmHkdfStreamingKey.Merge(m, src)
}
func (m *AesGcmHkdfStreamingKey) XXX_Size() int {
	return xxx_messageInfo_AesGcmHkdfStreamingKey.Size(m)
}
func (m *AesGcmHkdfStreamingKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AesGcmHkdfStreamingKey.DiscardUnknown(m)
}

var xxx_messageInfo_AesGcmHkdfStreamingKey proto.InternalMessageInfo

func (m *AesGcmHkdfStreamingKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *AesGcmHkdfStreamingKey) GetParams() *AesGcmHkdfStreamingParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *AesGcmHkdfStreamingKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func init() {
	proto.RegisterType((*AesGcmHkdfStreamingParams)(nil), "google.crypto.tink.AesGcmHkdfStreamingParams")
	proto.RegisterType((*AesGcmHkdfStreamingKeyFormat)(nil), "google.crypto.tink.AesGcmHkdfStreamingKeyFormat")
	proto.RegisterType((*AesGcmHkdfStreamingKey)(nil), "google.crypto.tink.AesGcmHkdfStreamingKey")
}

func init() {
	proto.RegisterFile("proto/aes_gcm_hkdf_streaming.proto", fileDescriptor_1dba2d882aaf5933)
}

var fileDescriptor_1dba2d882aaf5933 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x6b, 0xe2, 0x40,
	0x18, 0x65, 0x5c, 0xd0, 0xdd, 0x59, 0x95, 0x25, 0xb0, 0xbb, 0xda, 0x7a, 0x10, 0x7b, 0xf1, 0xd2,
	0x04, 0x14, 0x7a, 0xea, 0xa5, 0x42, 0x5b, 0x8b, 0x50, 0x24, 0x4a, 0x0f, 0xbd, 0x0c, 0x63, 0xf2,
	0x39, 0x19, 0xe2, 0x64, 0xc2, 0xcc, 0x28, 0x8d, 0x3f, 0xa4, 0x87, 0x1e, 0xfb, 0x4f, 0xfa, 0xcf,
	0x4a, 0x26, 0x4a, 0x4b, 0x1b, 0xe9, 0xa1, 0xc7, 0x37, 0xef, 0x7b, 0x6f, 0xde, 0xfb, 0xf8, 0xf0,
	0xd0, 0x44, 0x5c, 0x85, 0x24, 0xa5, 0xca, 0x64, 0x9e, 0xe1, 0x49, 0xec, 0xa5, 0x4a, 0x1a, 0xe9,
	0x51, 0xd0, 0x84, 0x05, 0x82, 0x44, 0x71, 0xb8, 0x24, 0xda, 0x28, 0xa0, 0x82, 0x27, 0xcc, 0xb5,
	0xa4, 0xe3, 0x30, 0x29, 0xd9, 0x0a, 0xdc, 0x40, 0x65, 0xa9, 0x91, 0x6e, 0x2e, 0x3b, 0x3a, 0x39,
	0x60, 0x14, 0x48, 0x21, 0x64, 0x52, 0x08, 0x7b, 0x2f, 0x08, 0xb7, 0x2f, 0x40, 0x5f, 0x07, 0x62,
	0x1c, 0x87, 0xcb, 0xd9, 0xde, 0x76, 0x4a, 0x15, 0x15, 0xda, 0x39, 0xc3, 0xff, 0x03, 0x9e, 0x46,
	0xa0, 0x0c, 0x3c, 0x18, 0xa2, 0x81, 0x09, 0x48, 0x0c, 0xd1, 0x7c, 0x0b, 0x2d, 0xd4, 0x45, 0xfd,
	0x86, 0xff, 0xf7, 0x8d, 0x9e, 0x15, 0xec, 0x8c, 0x6f, 0xc1, 0xe9, 0xe3, 0x3f, 0x21, 0x28, 0xbe,
	0x81, 0x90, 0xc4, 0x90, 0x15, 0x82, 0x8a, 0x15, 0x34, 0x77, 0xef, 0x13, 0xc8, 0xec, 0xe4, 0x08,
	0x37, 0x6d, 0xa1, 0x88, 0xea, 0x88, 0x98, 0x2c, 0x85, 0xd6, 0x8f, 0x2e, 0xea, 0x37, 0x07, 0x1d,
	0xf7, 0x73, 0x23, 0x77, 0x4c, 0x75, 0x34, 0xcf, 0x52, 0xf0, 0xeb, 0xb9, 0x66, 0x8f, 0x7a, 0x4f,
	0x08, 0x77, 0x4a, 0x3a, 0x4c, 0x20, 0xbb, 0x92, 0x4a, 0x50, 0xe3, 0xb4, 0x70, 0x6d, 0x03, 0x4a,
	0x73, 0x99, 0x58, 0xf7, 0x86, 0xbf, 0x87, 0xce, 0x25, 0xae, 0xa6, 0xb6, 0xaa, 0xed, 0xf3, 0x7b,
	0x70, 0x5a, 0xf6, 0xed, 0xc1, 0xfd, 0xf8, 0x3b, 0xb1, 0xd3, 0xc6, 0x3f, 0x3f, 0xf4, 0xac, 0xc5,
	0x45, 0xc1, 0xde, 0x23, 0xc2, 0xff, 0xca, 0xc3, 0xbd, 0x8f, 0x85, 0x0e, 0xc5, 0xaa, 0x7c, 0x27,
	0xd6, 0x31, 0xfe, 0x95, 0xc7, 0xda, 0xd0, 0xd5, 0xba, 0xd8, 0x6b, 0xdd, 0xcf, 0x73, 0xde, 0xe5,
	0x78, 0xc4, 0x70, 0x27, 0x90, 0xa2, 0xcc, 0xd8, 0x5e, 0xc6, 0x14, 0xdd, 0x9f, 0x33, 0x6e, 0xa2,
	0xf5, 0xc2, 0x0d, 0xa4, 0xf0, 0x8a, 0xb1, 0xaf, 0xef, 0x91, 0x30, 0x49, 0x2c, 0xff, 0x5c, 0xa9,
	0xce, 0x6f, 0x6e, 0x27, 0xd3, 0xd1, 0xa2, 0x6a, 0xf1, 0xf0, 0x35, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0xd0, 0x92, 0x6b, 0xd9, 0x02, 0x00, 0x00,
}
