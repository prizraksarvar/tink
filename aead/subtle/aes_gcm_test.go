// Copyright 2020 Google LLC
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

package subtle_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/prizraksarvar/tink/aead/subtle"
	"github.com/prizraksarvar/tink/subtle/random"
	"github.com/prizraksarvar/tink/testutil"
)

var keySizes = []int{16, 32}

// Since the tag size depends on the Seal() function of crypto library,
// this test checks that the tag size is always 128 bit.
func TestAESGCMTagLength(t *testing.T) {
	for _, keySize := range keySizes {
		key := random.GetRandomBytes(uint32(keySize))
		a, _ := subtle.NewAESGCM(key)
		ad := random.GetRandomBytes(32)
		pt := random.GetRandomBytes(32)
		ct, _ := a.Encrypt(pt, ad)
		actualTagSize := len(ct) - subtle.AESGCMIVSize - len(pt)
		if actualTagSize != subtle.AESGCMTagSize {
			t.Errorf("tag size is not 128 bit, it is %d bit", actualTagSize*8)
		}
	}
}

func TestAESGCMKeySize(t *testing.T) {
	for _, keySize := range keySizes {
		if _, err := subtle.NewAESGCM(make([]byte, keySize)); err != nil {
			t.Errorf("unexpected error when key size is %d btyes", keySize)
		}
		if _, err := subtle.NewAESGCM(make([]byte, keySize+1)); err == nil {
			t.Errorf("expect an error when key size is not supported %d", keySize)
		}
	}
}

func TestAESGCMEncryptDecrypt(t *testing.T) {
	for _, keySize := range keySizes {
		key := random.GetRandomBytes(uint32(keySize))
		a, err := subtle.NewAESGCM(key)
		if err != nil {
			t.Errorf("unexpected error when creating new cipher: %s", err)
		}
		ad := random.GetRandomBytes(5)
		for ptSize := 0; ptSize < 75; ptSize++ {
			pt := random.GetRandomBytes(uint32(ptSize))
			ct, err := a.Encrypt(pt, ad)
			if err != nil {
				t.Errorf("unexpected error in encryption: keySize %v, ptSize %v", keySize, ptSize)
			}
			decrypted, err := a.Decrypt(ct, ad)
			if err != nil {
				t.Errorf("unexpected error in decryption: keySize %v, ptSize %v", keySize, ptSize)
			}
			if !bytes.Equal(pt, decrypted) {
				t.Errorf("decrypted text and plaintext don't match: keySize %v, ptSize %v", keySize, ptSize)
			}
		}
	}
}

func TestAESGCMLongMessages(t *testing.T) {
	ptSize := 16
	for ptSize <= 1<<24 {
		pt := random.GetRandomBytes(uint32(ptSize))
		ad := random.GetRandomBytes(uint32(ptSize / 3))
		for _, keySize := range keySizes {
			key := random.GetRandomBytes(uint32(keySize))
			a, _ := subtle.NewAESGCM(key)
			ct, _ := a.Encrypt(pt, ad)
			decrypted, _ := a.Decrypt(ct, ad)
			if !bytes.Equal(pt, decrypted) {
				t.Errorf("decrypted text and plaintext don't match: keySize %v, ptSize %v", keySize, ptSize)
			}
		}
		ptSize += 5 * ptSize / 11
	}
}

func TestAESGCMModifyCiphertext(t *testing.T) {
	ad := random.GetRandomBytes(33)
	key := random.GetRandomBytes(16)
	pt := random.GetRandomBytes(32)
	a, _ := subtle.NewAESGCM(key)
	ct, _ := a.Encrypt(pt, ad)
	// flipping bits
	for i := 0; i < len(ct); i++ {
		tmp := ct[i]
		for j := 0; j < 8; j++ {
			ct[i] ^= 1 << uint8(j)
			if _, err := a.Decrypt(ct, ad); err == nil {
				t.Errorf("expect an error when flipping bit of ciphertext: byte %d, bit %d", i, j)
			}
			ct[i] = tmp
		}
	}
	// truncated ciphertext
	for i := 1; i < len(ct); i++ {
		if _, err := a.Decrypt(ct[:i], ad); err == nil {
			t.Errorf("expect an error ciphertext is truncated until byte %d", i)
		}
	}
	// modify additinal authenticated data
	for i := 0; i < len(ad); i++ {
		tmp := ad[i]
		for j := 0; j < 8; j++ {
			ad[i] ^= 1 << uint8(j)
			if _, err := a.Decrypt(ct, ad); err == nil {
				t.Errorf("expect an error when flipping bit of ad: byte %d, bit %d", i, j)
			}
			ad[i] = tmp
		}
	}
}

/**
 * This is a very simple test for the randomness of the nonce.
 * The test simply checks that the multiple ciphertexts of the same
 * message are distinct.
 */
func TestAESGCMRandomNonce(t *testing.T) {
	nSample := 1 << 17
	key := random.GetRandomBytes(16)
	pt := []byte{}
	ad := []byte{}
	a, _ := subtle.NewAESGCM(key)
	ctSet := make(map[string]bool)
	for i := 0; i < nSample; i++ {
		ct, _ := a.Encrypt(pt, ad)
		ctHex := hex.EncodeToString(ct)
		_, existed := ctSet[ctHex]
		if existed {
			t.Errorf("nonce is repeated after %d samples", i)
		}
		ctSet[ctHex] = true
	}
}

func TestAESGCMWycheproofCases(t *testing.T) {
	testutil.SkipTestIfTestSrcDirIsNotSet(t)
	suite := new(AEADSuite)
	if err := testutil.PopulateSuite(suite, "aes_gcm_test.json"); err != nil {
		t.Fatalf("failed populating suite: %s", err)
	}
	for _, group := range suite.TestGroups {
		if err := subtle.ValidateAESKeySize(group.KeySize / 8); err != nil {
			continue
		}
		if group.IvSize != subtle.AESGCMIVSize*8 {
			continue
		}
		for _, test := range group.Tests {
			caseName := fmt.Sprintf("%s-%s(%d,%d):Case-%d",
				suite.Algorithm, group.Type, group.KeySize, group.TagSize, test.CaseID)
			t.Run(caseName, func(t *testing.T) { runAESGCMWycheproofCase(t, test) })
		}
	}
}

func runAESGCMWycheproofCase(t *testing.T, tc *AEADCase) {
	var combinedCt []byte
	combinedCt = append(combinedCt, tc.Iv...)
	combinedCt = append(combinedCt, tc.Ct...)
	combinedCt = append(combinedCt, tc.Tag...)
	// create cipher and do encryption
	cipher, err := subtle.NewAESGCM(tc.Key)
	if err != nil {
		t.Fatalf("cannot create new instance of AESGCM in test case: %s", err)
	}
	decrypted, err := cipher.Decrypt(combinedCt, tc.Aad)
	if err != nil {
		if tc.Result == "valid" {
			t.Errorf("unexpected error in test case: %s", err)
		}
	} else {
		if tc.Result == "invalid" {
			t.Error("decrypted invalid test case")
		}
		if !bytes.Equal(decrypted, tc.Msg) {
			t.Error("incorrect decryption in test case")
		}
	}
}
