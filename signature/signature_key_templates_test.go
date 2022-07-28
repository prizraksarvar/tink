// Copyright 2018 Google LLC
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

package signature_test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/prizraksarvar/tink/go/keyset"
	tinkpb "github.com/prizraksarvar/tink/go/proto/tink_go_proto"
	"github.com/prizraksarvar/tink/go/signature"
	"github.com/prizraksarvar/tink/go/testutil"
)

func TestKeyTemplates(t *testing.T) {
	testutil.SkipTestIfTestSrcDirIsNotSet(t)
	var testCases = []struct {
		name     string
		template *tinkpb.KeyTemplate
	}{
		{name: "ECDSA_P256",
			template: signature.ECDSAP256KeyTemplate()},
		{name: "ECDSA_P384",
			template: signature.ECDSAP384KeyTemplate()},
		{name: "ECDSA_P521",
			template: signature.ECDSAP521KeyTemplate()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want, err := testutil.KeyTemplateProto("signature", tc.name)
			if err != nil {
				t.Fatalf("testutil.KeyTemplateProto('signature', tc.name) failed: %s", err)
			}
			if !proto.Equal(want, tc.template) {
				t.Errorf("template %s is not equal to '%s'", tc.name, tc.template)
			}
			if err := testSignVerify(tc.template); err != nil {
				t.Errorf("%v", err)
			}
		})
	}
}

func TestKeyWithoutPrefixTemplates(t *testing.T) {
	testutil.SkipTestIfTestSrcDirIsNotSet(t)
	var testCases = []struct {
		name     string
		template *tinkpb.KeyTemplate
	}{
		{name: "ECDSA_P256",
			template: signature.ECDSAP256KeyWithoutPrefixTemplate()},
		{name: "ECDSA_P384",
			template: signature.ECDSAP384KeyWithoutPrefixTemplate()},
		{name: "ECDSA_P521",
			template: signature.ECDSAP521KeyWithoutPrefixTemplate()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want, err := testutil.KeyTemplateProto("signature", tc.name)
			if err != nil {
				t.Fatalf("testutil.KeyTemplateProto('signature', tc.name) failed: %s", err)
			}
			want.OutputPrefixType = tinkpb.OutputPrefixType_RAW
			if !proto.Equal(want, tc.template) {
				t.Errorf("template %s is not equal to '%s'", tc.name, tc.template)
			}
			if err := testSignVerify(tc.template); err != nil {
				t.Errorf("%v", err)
			}
		})
	}
}

func testSignVerify(template *tinkpb.KeyTemplate) error {
	privateHandle, err := keyset.NewHandle(template)
	if err != nil {
		return fmt.Errorf("keyset.NewHandle(tc.template) failed: %s", err)
	}

	signer, err := signature.NewSigner(privateHandle)
	if err != nil {
		return fmt.Errorf("signature.NewSigner(privateHandle) failed: %s", err)
	}
	publicHandle, err := privateHandle.Public()
	if err != nil {
		return fmt.Errorf("privateHandle.Public() failed: %s", err)
	}
	verifier, err := signature.NewVerifier(publicHandle)
	if err != nil {
		return fmt.Errorf("signature.NewVerifier(publicHandle) failed: %s", err)
	}

	var testInputs = []struct {
		message1 []byte
		message2 []byte
	}{
		{
			message1: []byte("this data needs to be signed"),
			message2: []byte("this data needs to be signed"),
		}, {
			message1: []byte(""),
			message2: []byte(""),
		}, {
			message1: []byte(""),
			message2: nil,
		}, {
			message1: nil,
			message2: []byte(""),
		}, {
			message1: nil,
			message2: nil,
		},
	}
	for _, ti := range testInputs {
		sig, err := signer.Sign(ti.message1)
		if err != nil {
			return fmt.Errorf("signer.Sign(ti.message1) failed: %s", err)
		}
		if err := verifier.Verify(sig, ti.message2); err != nil {
			return fmt.Errorf("verifier.Verify(sig, ti.message2) failed: %s", err)
		}
	}
	return nil
}
