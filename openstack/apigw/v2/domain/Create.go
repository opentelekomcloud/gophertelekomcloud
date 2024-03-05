package domain

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	GatewayID string `json:"-"`
	GroupID   string `json:"-"`
	// Minimum SSL version. TLS 1.1 and TLS 1.2 are supported.
	MinSslVersion string `json:"min_ssl_version,omitempty"`
	// Whether to enable HTTP redirection to HTTPS.
	// The value false means disable and true means enable. The default value is false.
	IsHttpRedirectToHttps *bool `json:"is_http_redirect_to_https,omitempty"`
	// Custom domain name.
	// It can contain a maximum of 255 characters and must comply with domain name specifications.
	UrlDomain string `json:"url_domain" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*DomainResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains
	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID, "domains"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res DomainResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DomainResp struct {
	// Custom domain name.
	UrlDomain string `json:"url_domain"`
	// Domain ID.
	ID string `json:"id"`
	// CNAME resolution status.
	// 1: not resolved
	// 2: resolving
	// 3: resolved
	// 4: resolution failed
	Status int `json:"status"`
	// Minimum SSL version supported.
	MinSslVersion string `json:"min_ssl_version"`
	// Whether to enable HTTP redirection to HTTPS.
	// The value false means disable and true means enable. The default value is false.
	IsHttpRedirectToHttps bool `json:"is_http_redirect_to_https"`
	// Whether to enable client certificate verification.
	// This parameter is available only when a certificate is bound.
	// It is enabled by default if trusted_root_ca exists, and disabled if trusted_root_ca does not exist.
	VerifiedClientCertificateEnabled bool `json:"verified_client_certificate_enabled"`
}
