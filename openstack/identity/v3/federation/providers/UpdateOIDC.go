package providers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/federation"
)

type UpdateOIDCOpts struct {
	// Identity provider name.
	IdpIp string `json:"-"`
	// Mode to access manage identity provider
	AccessMode string `json:"access_mode,omitempty"`
	// URL of the OpenID Connect identity provider.
	// This field corresponds to the iss field in the ID token.
	IdpUrl string `json:"idp_url,omitempty"`
	// ID of a client registered with the OpenID Connect identity provider.
	ClientId string `json:"client_id,omitempty"`
	// Authorization endpoint of the OpenID Connect identity provider.
	AuthEndpoint string `json:"authorization_endpoint,omitempty"`
	// Scope of authorization requests.
	Scope string `json:"scope,omitempty"`
	// Response type
	ResponseType string `json:"response_type,omitempty"`
	// Response mode
	ResponseMode string `json:"response_mode,omitempty"`
	// SigningKey
	SigningKey string `json:"signing_key,omitempty"`
}

func UpdateOIDC(c *golangsdk.ServiceClient, opts UpdateOIDCOpts) (*CreateOIDCOpts, error) {
	b, err := build.RequestBody(opts, "openid_connect_config")
	if err != nil {
		return nil, err
	}
	raw, err := c.Put(c.ServiceURL(federation.BaseURL, "identity-providers", opts.IdpIp, "openid-connect-config"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}

	var res CreateOIDCOpts
	err = extract.IntoStructPtr(raw.Body, &res, "openid_connect_config")
	return &res, err
}
