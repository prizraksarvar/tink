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
// source: third_party/tink/proto/aes_ctr_hmac_streaming.proto

package aes_ctr_hmac_streaming_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common_go_proto "github.com/prizraksarvar/tink/go/proto/common_go_proto"
	hmac_go_proto "github.com/prizraksarvar/tink/go/proto/hmac_go_proto"
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

type AesCtrHmacStreamingParams struct {
	CiphertextSegmentSize uint32                    `protobuf:"varint,1,opt,name=ciphertext_segment_size,json=ciphertextSegmentSize,proto3" json:"ciphertext_segment_size,omitempty"`
	DerivedKeySize        uint32                    `protobuf:"varint,2,opt,name=derived_key_size,json=derivedKeySize,proto3" json:"derived_key_size,omitempty"`
	HkdfHashType          common_go_proto.HashType  `protobuf:"varint,3,opt,name=hkdf_hash_type,json=hkdfHashType,proto3,enum=google.crypto.tink.HashType" json:"hkdf_hash_type,omitempty"`
	HmacParams            *hmac_go_proto.HmacParams `protobuf:"bytes,4,opt,name=hmac_params,json=hmacParams,proto3" json:"hmac_params,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                  `json:"-"`
	XXX_unrecognized      []byte                    `json:"-"`
	XXX_sizecache         int32                     `json:"-"`
}

func (m *AesCtrHmacStreamingParams) Reset()         { *m = AesCtrHmacStreamingParams{} }
func (m *AesCtrHmacStreamingParams) String() string { return proto.CompactTextString(m) }
func (*AesCtrHmacStreamingParams) ProtoMessage()    {}
func (*AesCtrHmacStreamingParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa520d3eb9ab35e, []int{0}
}

func (m *AesCtrHmacStreamingParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AesCtrHmacStreamingParams.Unmarshal(m, b)
}
func (m *AesCtrHmacStreamingParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AesCtrHmacStreamingParams.Marshal(b, m, deterministic)
}
func (m *AesCtrHmacStreamingParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AesCtrHmacStreamingParams.Merge(m, src)
}
func (m *AesCtrHmacStreamingParams) XXX_Size() int {
	return xxx_messageInfo_AesCtrHmacStreamingParams.Size(m)
}
func (m *AesCtrHmacStreamingParams) XXX_DiscardUnknown() {
	xxx_messageInfo_AesCtrHmacStreamingParams.DiscardUnknown(m)
}

var xxx_messageInfo_AesCtrHmacStreamingParams proto.InternalMessageInfo

func (m *AesCtrHmacStreamingParams) GetCiphertextSegmentSize() uint32 {
	if m != nil {
		return m.CiphertextSegmentSize
	}
	return 0
}

func (m *AesCtrHmacStreamingParams) GetDerivedKeySize() uint32 {
	if m != nil {
		return m.DerivedKeySize
	}
	return 0
}

func (m *AesCtrHmacStreamingParams) GetHkdfHashType() common_go_proto.HashType {
	if m != nil {
		return m.HkdfHashType
	}
	return common_go_proto.HashType_UNKNOWN_HASH
}

func (m *AesCtrHmacStreamingParams) GetHmacParams() *hmac_go_proto.HmacParams {
	if m != nil {
		return m.HmacParams
	}
	return nil
}

type AesCtrHmacStreamingKeyFormat struct {
	Params               *AesCtrHmacStreamingParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	KeySize              uint32                     `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AesCtrHmacStreamingKeyFormat) Reset()         { *m = AesCtrHmacStreamingKeyFormat{} }
func (m *AesCtrHmacStreamingKeyFormat) String() string { return proto.CompactTextString(m) }
func (*AesCtrHmacStreamingKeyFormat) ProtoMessage()    {}
func (*AesCtrHmacStreamingKeyFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa520d3eb9ab35e, []int{1}
}

func (m *AesCtrHmacStreamingKeyFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AesCtrHmacStreamingKeyFormat.Unmarshal(m, b)
}
func (m *AesCtrHmacStreamingKeyFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AesCtrHmacStreamingKeyFormat.Marshal(b, m, deterministic)
}
func (m *AesCtrHmacStreamingKeyFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AesCtrHmacStreamingKeyFormat.Merge(m, src)
}
func (m *AesCtrHmacStreamingKeyFormat) XXX_Size() int {
	return xxx_messageInfo_AesCtrHmacStreamingKeyFormat.Size(m)
}
func (m *AesCtrHmacStreamingKeyFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_AesCtrHmacStreamingKeyFormat.DiscardUnknown(m)
}

var xxx_messageInfo_AesCtrHmacStreamingKeyFormat proto.InternalMessageInfo

func (m *AesCtrHmacStreamingKeyFormat) GetParams() *AesCtrHmacStreamingParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *AesCtrHmacStreamingKeyFormat) GetKeySize() uint32 {
	if m != nil {
		return m.KeySize
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.AesCtrHmacStreamingKey
type AesCtrHmacStreamingKey struct {
	Version              uint32                     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params               *AesCtrHmacStreamingParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KeyValue             []byte                     `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AesCtrHmacStreamingKey) Reset()         { *m = AesCtrHmacStreamingKey{} }
func (m *AesCtrHmacStreamingKey) String() string { return proto.CompactTextString(m) }
func (*AesCtrHmacStreamingKey) ProtoMessage()    {}
func (*AesCtrHmacStreamingKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa520d3eb9ab35e, []int{2}
}

func (m *AesCtrHmacStreamingKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AesCtrHmacStreamingKey.Unmarshal(m, b)
}
func (m *AesCtrHmacStreamingKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AesCtrHmacStreamingKey.Marshal(b, m, deterministic)
}
func (m *AesCtrHmacStreamingKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AesCtrHmacStreamingKey.Merge(m, src)
}
func (m *AesCtrHmacStreamingKey) XXX_Size() int {
	return xxx_messageInfo_AesCtrHmacStreamingKey.Size(m)
}
func (m *AesCtrHmacStreamingKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AesCtrHmacStreamingKey.DiscardUnknown(m)
}

var xxx_messageInfo_AesCtrHmacStreamingKey proto.InternalMessageInfo

func (m *AesCtrHmacStreamingKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *AesCtrHmacStreamingKey) GetParams() *AesCtrHmacStreamingParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *AesCtrHmacStreamingKey) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func init() {
	proto.RegisterType((*AesCtrHmacStreamingParams)(nil), "google.crypto.tink.AesCtrHmacStreamingParams")
	proto.RegisterType((*AesCtrHmacStreamingKeyFormat)(nil), "google.crypto.tink.AesCtrHmacStreamingKeyFormat")
	proto.RegisterType((*AesCtrHmacStreamingKey)(nil), "google.crypto.tink.AesCtrHmacStreamingKey")
}

func init() {
	proto.RegisterFile("proto/aes_ctr_hmac_streaming.proto", fileDescriptor_eaa520d3eb9ab35e)
}

var fileDescriptor_eaa520d3eb9ab35e = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcf, 0xcb, 0xd3, 0x30,
	0x18, 0x26, 0x53, 0xbe, 0x4f, 0xf3, 0x7d, 0x0e, 0x09, 0xa8, 0xfd, 0xe6, 0x90, 0x39, 0x2f, 0xbd,
	0xd8, 0xc2, 0x06, 0x9e, 0x04, 0x71, 0xa2, 0x0c, 0x76, 0x19, 0x9d, 0x78, 0xf0, 0x60, 0xc8, 0xd2,
	0xd7, 0x26, 0x74, 0x69, 0x4a, 0x92, 0x0d, 0xbb, 0x93, 0xf8, 0x47, 0xf8, 0xf7, 0x4a, 0x93, 0x4e,
	0x45, 0x3b, 0x3c, 0x78, 0x6a, 0xdf, 0x3c, 0x3f, 0xf2, 0xe4, 0xe1, 0xc5, 0x73, 0x27, 0xa4, 0xc9,
	0x69, 0xcd, 0x8c, 0x6b, 0x52, 0x27, 0xab, 0x32, 0xad, 0x8d, 0x76, 0x3a, 0x65, 0x60, 0x29, 0x77,
	0x86, 0x0a, 0xc5, 0x38, 0xb5, 0xce, 0x00, 0x53, 0xb2, 0x2a, 0x12, 0x0f, 0x12, 0x52, 0x68, 0x5d,
	0xec, 0x20, 0xe1, 0xa6, 0xa9, 0x9d, 0x4e, 0x5a, 0xd9, 0xe8, 0xd9, 0x19, 0x23, 0xae, 0x95, 0xd2,
	0x55, 0x10, 0x8e, 0x9e, 0x9e, 0x21, 0xb5, 0xb7, 0x04, 0xca, 0xf4, 0xdb, 0x00, 0xdf, 0xbc, 0x06,
	0xfb, 0xc6, 0x99, 0xa5, 0x62, 0x7c, 0x73, 0xba, 0x79, 0xcd, 0x0c, 0x53, 0x96, 0xbc, 0xc0, 0x8f,
	0xb8, 0xac, 0x05, 0x18, 0x07, 0x5f, 0x1c, 0xb5, 0x50, 0x28, 0xa8, 0x1c, 0xb5, 0xf2, 0x08, 0x11,
	0x9a, 0xa0, 0xf8, 0x5e, 0xf6, 0xe0, 0x17, 0xbc, 0x09, 0xe8, 0x46, 0x1e, 0x81, 0xc4, 0xf8, 0x7e,
	0x0e, 0x46, 0x1e, 0x20, 0xa7, 0x25, 0x34, 0x41, 0x30, 0xf0, 0x82, 0x61, 0x77, 0xbe, 0x82, 0xc6,
	0x33, 0x17, 0x78, 0x28, 0xca, 0xfc, 0x33, 0x15, 0xcc, 0x0a, 0xea, 0x9a, 0x1a, 0xa2, 0x5b, 0x13,
	0x14, 0x0f, 0x67, 0xe3, 0xe4, 0xef, 0x47, 0x27, 0x4b, 0x66, 0xc5, 0xfb, 0xa6, 0x86, 0xec, 0xba,
	0xd5, 0x9c, 0x26, 0xf2, 0x0a, 0x5f, 0xf9, 0xde, 0x6a, 0x1f, 0x3a, 0xba, 0x3d, 0x41, 0xf1, 0xd5,
	0xec, 0x49, 0xaf, 0x81, 0x62, 0x3c, 0x3c, 0x2d, 0xc3, 0xe2, 0xe7, 0xff, 0xf4, 0x2b, 0xc2, 0xe3,
	0x9e, 0x12, 0x56, 0xd0, 0xbc, 0xd3, 0x46, 0x31, 0x47, 0xde, 0xe2, 0x8b, 0xce, 0x1c, 0x79, 0xf3,
	0xe7, 0x7d, 0xe6, 0x67, 0x6b, 0xcc, 0x3a, 0x31, 0xb9, 0xc1, 0x77, 0xfe, 0xa8, 0xe3, 0xb2, 0x0c,
	0x3d, 0x4c, 0xbf, 0x23, 0xfc, 0xb0, 0x3f, 0x02, 0x89, 0xf0, 0xe5, 0x01, 0x8c, 0x95, 0xba, 0xea,
	0x4a, 0x3f, 0x8d, 0xbf, 0xc5, 0x1a, 0xfc, 0x4f, 0xac, 0xc7, 0xf8, 0x6e, 0x1b, 0xeb, 0xc0, 0x76,
	0xfb, 0x50, 0xff, 0x75, 0xd6, 0xe6, 0xfc, 0xd0, 0xce, 0x8b, 0x4f, 0x78, 0xcc, 0xb5, 0xea, 0x33,
	0xf6, 0x0b, 0xb4, 0x46, 0x1f, 0x5f, 0x16, 0xd2, 0x89, 0xfd, 0x36, 0xe1, 0x5a, 0xa5, 0x81, 0xf6,
	0xef, 0xcd, 0xa6, 0x85, 0xa6, 0x1e, 0xdf, 0x5e, 0xf8, 0xcf, 0xfc, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x67, 0x84, 0x6f, 0x1a, 0x03, 0x00, 0x00,
}
