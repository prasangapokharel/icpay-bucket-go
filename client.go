package icpaybucket

import (
	"fmt"
	"net/url"

	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/identity"
	"github.com/aviate-labs/agent-go/principal"
)

// Client talks to the ICPay bucket canister.
//
// Writes (upload/delete) are authorized by a bucket API key or by an
// Internet Identity. Reads (list/download/publicUrl) work for any caller on a
// public bucket.
type Client struct {
	agent      *agent.Agent
	canisterID principal.Principal
	apiKey     string
}

// ClientOptions configures a Client.
type ClientOptions struct {
	// APIKey authorizes writes without an identity.
	APIKey string
	// CanisterID overrides the default ICPay canister.
	CanisterID string
	// Host overrides the IC host (default https://icp0.io, mainnet).
	Host string
	// Identity overrides the anonymous caller used for updates.
	Identity identity.Identity
}

// NewClient creates a Client. An anonymous caller is used unless an identity
// is supplied; the API key (when set) authorizes writes as the bucket owner.
func NewClient(options ClientOptions) (*Client, error) {
	canisterID := options.CanisterID
	if canisterID == "" {
		canisterID = DefaultCanisterID
	}
	canister, err := principal.Decode(canisterID)
	if err != nil {
		return nil, fmt.Errorf("invalid canister id: %w", err)
	}

	ident := options.Identity
	if ident == nil {
		ident = identity.AnonymousIdentity{}
	}

	host := options.Host
	if host == "" {
		host = "https://icp0.io"
	}
	hostURL, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("invalid host: %w", err)
	}

	ag, err := agent.New(agent.Config{
		Identity:     ident,
		ClientConfig: []agent.ClientOption{agent.WithHostURL(hostURL)},
	})
	if err != nil {
		return nil, fmt.Errorf("create agent: %w", err)
	}

	return &Client{agent: ag, canisterID: canister, apiKey: options.APIKey}, nil
}

func optString(value string) *string {
	return &value
}

// UploadFile stores a file in a single call (up to ~1.85 MB).
func (c *Client) UploadFile(bucketID, path string, data []byte, contentType string, apiKey string) (resultFileID, error) {
	if len(data) > SingleMaxBytes {
		return resultFileID{}, fmt.Errorf("file too large for single upload")
	}
	key := c.resolveAPIKey(apiKey)
	out := resultFileID{}
	err := c.agent.Call(
		c.canisterID,
		"uploadFile",
		[]any{bucketID, path, data, contentType, key},
		[]any{&out},
	)
	return out, err
}

// DeleteFile removes a file. Authorized by the API key.
func (c *Client) DeleteFile(bucketID, path, apiKey string) (resultUnit, error) {
	key := c.resolveAPIKey(apiKey)
	out := resultUnit{}
	err := c.agent.Call(
		c.canisterID,
		"deleteFile",
		[]any{bucketID, path, key},
		[]any{&out},
	)
	return out, err
}

// ListFiles returns a paginated file list for a bucket. Public buckets allow
// any caller.
func (c *Client) ListFiles(bucketID string, page, pageSize uint64) (resultFileList, error) {
	out := resultFileList{}
	err := c.agent.Query(
		c.canisterID,
		"listFiles",
		[]any{bucketID, idl.NewNat(page), idl.NewNat(pageSize)},
		[]any{&out},
	)
	return out, err
}

// DownloadFile returns file bytes (public + private buckets).
func (c *Client) DownloadFile(bucketID, path string) (resultBlob, error) {
	out := resultBlob{}
	err := c.agent.Query(
		c.canisterID,
		"downloadFile",
		[]any{bucketID, path},
		[]any{&out},
	)
	return out, err
}

// GetPublicFileURL returns the canister-provided CDN URL for a file.
func (c *Client) GetPublicFileURL(bucketID, path string) (resultString, error) {
	out := resultString{}
	err := c.agent.Query(
		c.canisterID,
		"getPublicFileUrl",
		[]any{bucketID, path},
		[]any{&out},
	)
	return out, err
}

// PublicURL builds the CDN URL locally — no canister call.
func (c *Client) PublicURL(bucketName, path string) string {
	return PublicFileURL(c.canisterID.String(), bucketName, path)
}

func (c *Client) resolveAPIKey(explicit string) *string {
	if explicit != "" {
		return optString(explicit)
	}
	if c.apiKey != "" {
		return optString(c.apiKey)
	}
	return nil
}
