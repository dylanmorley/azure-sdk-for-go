// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/log"
)

// NewSharedKeyCredential creates an immutable SharedKeyCredential containing the
// account's primary or secondary key.
func NewSharedKeyCredential(accountKey string) (*SharedKeyCredential, error) {
	c := SharedKeyCredential{}
	if err := c.SetAccountKey(accountKey); err != nil {
		return nil, err
	}
	return &c, nil
}

// SharedKeyCredential contains an account's name and its primary or secondary key.
// It is immutable making it shareable and goroutine-safe.
type SharedKeyCredential struct {
	// Only the NewSharedKeyCredential method should set these; all other methods should treat them as read-only
	accountKey atomic.Value // []byte
}

// SetAccountKey replaces the existing account key with the specified account key.
func (c *SharedKeyCredential) SetAccountKey(accountKey string) error {
	bytes, err := base64.StdEncoding.DecodeString(accountKey)
	if err != nil {
		return fmt.Errorf("decode account key: %w", err)
	}
	c.accountKey.Store(bytes)
	return nil
}

// computeHMACSHA256 generates a hash signature for an HTTP request
func (c *SharedKeyCredential) computeHMACSHA256(s string) (base64String string) {
	h := hmac.New(sha256.New, c.accountKey.Load().([]byte))
	_, _ = h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *SharedKeyCredential) buildCanonicalizedAuthHeaderFromRequest(req *policy.Request) (string, error) {
	var opValues cosmosOperationContext
	value := ""

	if req.OperationValue(&opValues) {
		resourceTypePath, err := getResourcePath(opValues.resourceType)

		if err != nil {
			return "", err
		}

		resourceAddress := opValues.resourceAddress
		if opValues.isRidBased {
			resourceAddress = strings.ToLower(resourceAddress)
		}

		value = c.buildCanonicalizedAuthHeader(req.Raw().Method, resourceTypePath, resourceAddress, req.Raw().Header.Get(headerXmsDate), "master", "1.0")
	}

	return value, nil
}

//where date is like time.RFC1123 but hard-codes GMT as the time zone
func (c *SharedKeyCredential) buildCanonicalizedAuthHeader(method, resourceType, resourceAddress, xmsDate, tokenType, version string) string {
	if method == "" || resourceType == "" {
		return ""
	}

	// https://docs.microsoft.com/en-us/rest/api/cosmos-db/access-control-on-cosmosdb-resources#constructkeytoken
	stringToSign := join(strings.ToLower(method), "\n", strings.ToLower(resourceType), "\n", resourceAddress, "\n", strings.ToLower(xmsDate), "\n", "", "\n")
	signature := c.computeHMACSHA256(stringToSign)

	return url.QueryEscape(join("type=" + tokenType + "&ver=" + version + "&sig=" + signature))
}

type sharedKeyCredPolicy struct {
	cred *SharedKeyCredential
}

func newSharedKeyCredPolicy(cred *SharedKeyCredential, opts runtime.AuthenticationOptions) *sharedKeyCredPolicy {
	s := &sharedKeyCredPolicy{
		cred: cred,
	}

	return s
}

func (c *SharedKeyCredential) NewAuthenticationPolicy(options runtime.AuthenticationOptions) policy.Policy {
	return newSharedKeyCredPolicy(c, options)
}

func (s *sharedKeyCredPolicy) Do(req *policy.Request) (*http.Response, error) {
	// Add a x-ms-date header if it doesn't already exist
	if d := req.Raw().Header.Get(headerXmsDate); d == "" {
		req.Raw().Header.Set(headerXmsDate, time.Now().UTC().Format(http.TimeFormat))
	}

	authHeader, err := s.cred.buildCanonicalizedAuthHeaderFromRequest(req)
	if err != nil {
		return nil, err
	}

	if authHeader != "" {
		req.Raw().Header.Set(headerAuthorization, authHeader)
	}

	response, err := req.Next()
	if err != nil && response != nil && response.StatusCode == http.StatusForbidden {
		// Service failed to authenticate request, log it
		log.Write(log.Response, "===== HTTP Forbidden status, Authorization:\n"+authHeader+"\n=====\n")
	}
	return response, err
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
