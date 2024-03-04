package domain

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	GatewayID string `json:"-"`
	GroupID   string `json:"-"`
	DomainID  string `json:"-"`
	// Minimum SSL version. TLS 1.1 and TLS 1.2 are supported.
	MinSslVersion string `json:"min_ssl_version" required:"true"`
	// Whether to enable HTTP redirection to HTTPS.
	// The value false means disable and true means enable. The default value is false.
	IsHttpRedirectToHttps *bool `json:"is_http_redirect_to_https,omitempty"`
	// Whether to enable client certificate verification.
	// This parameter is available only when a certificate is bound.
	// It is enabled by default if trusted_root_ca exists, and disabled if trusted_root_ca does not exist.
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*DomainResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}
	raw, err := client.Put(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID, "domains", opts.DomainID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res DomainResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
