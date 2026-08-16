package icpaybucket

import (
	"fmt"
	"net/url"

	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/aviate-labs/agent-go/identity"
	"github.com/aviate-labs/agent-go/principal"
)

type ClientOptions struct {
	APIKey     string
	CanisterID string
	Host       string
	Identity   identity.Identity
}

type Client struct {
	agent      *agent.Agent
	canisterID principal.Principal
	apiKey     string
}

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

func (c *Client) PublicURL(bucketName, path string) string {
	return PublicFileURL(c.canisterID.String(), bucketName, path)
}

func (c *Client) resolveAPIKey(explicit string) *string {
	if explicit != "" {
		return &explicit
	}
	if c.apiKey != "" {
		return &c.apiKey
	}
	return nil
}

func nat(v uint64) idl.Nat {
	return idl.NewNat(v)
}
