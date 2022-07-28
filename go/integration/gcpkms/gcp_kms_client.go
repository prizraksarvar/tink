// Copyright 2017 Google Inc.
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

// Package gcpkms provides integration with the GCP Cloud KMS.
// Tink APIs work with GCP and AWS KMS.
package gcpkms

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/prizraksarvar/tink/go/core/registry"
	"github.com/prizraksarvar/tink/go/tink"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudkms/v1"
)

const (
	gcpPrefix = "gcp-kms://"
)

var (
	errCred       = errors.New("invalid credential path")
	tinkUserAgent = "Tink/" + tink.Version + " Golang/" + runtime.Version()
)

// gcpClient represents a client that connects to the GCP KMS backend.
type gcpClient struct {
	keyURIPrefix string
	kms          *cloudkms.Service
}

var _ registry.KMSClient = (*gcpClient)(nil)

// NewClient returns a new GCP KMS client which will use default
// credentials to handle keys with uriPrefix prefix.
// uriPrefix must have the following format: 'gcp-kms://[:path]'.
func NewClient(uriPrefix string) (registry.KMSClient, error) {
	if !strings.HasPrefix(strings.ToLower(uriPrefix), gcpPrefix) {
		return nil, fmt.Errorf("uriPrefix must start with %s", gcpPrefix)
	}

	ctx := context.Background()
	client, err := google.DefaultClient(ctx, cloudkms.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	kmsService, err := cloudkms.New(client)
	if err != nil {
		return nil, err
	}

	return &gcpClient{
		keyURIPrefix: uriPrefix,
		kms:          kmsService,
	}, nil
}

// NewClientWithCredentials returns a new GCP KMS client which will use given
// credentials to handle keys with uriPrefix prefix.
// uriPrefix must have the following format: 'gcp-kms://[:path]'.
func NewClientWithCredentials(uriPrefix string, credentialPath string) (registry.KMSClient, error) {
	if !strings.HasPrefix(strings.ToLower(uriPrefix), gcpPrefix) {
		return nil, fmt.Errorf("uriPrefix must start with %s", gcpPrefix)
	}

	ctx := context.Background()
	if len(credentialPath) <= 0 {
		return nil, errCred
	}
	data, err := ioutil.ReadFile(credentialPath)
	if err != nil {
		return nil, errCred
	}
	creds, err := google.CredentialsFromJSON(ctx, data, "https://www.googleapis.com/auth/cloudkms")
	if err != nil {
		return nil, errCred
	}
	client := oauth2.NewClient(ctx, creds.TokenSource)
	kmsService, err := cloudkms.New(client)
	kmsService.UserAgent = tinkUserAgent
	if err != nil {
		return nil, err
	}

	return &gcpClient{
		keyURIPrefix: uriPrefix,
		kms:          kmsService,
	}, nil
}

// Supported true if this client does support keyURI
func (c *gcpClient) Supported(keyURI string) bool {
	return strings.HasPrefix(keyURI, c.keyURIPrefix)
}

// GetAEAD gets an AEAD backend by keyURI.
func (c *gcpClient) GetAEAD(keyURI string) (tink.AEAD, error) {
	if !c.Supported(keyURI) {
		return nil, errors.New("unsupported keyURI")
	}

	uri := strings.TrimPrefix(keyURI, gcpPrefix)
	return newGCPAEAD(uri, c.kms), nil
}
