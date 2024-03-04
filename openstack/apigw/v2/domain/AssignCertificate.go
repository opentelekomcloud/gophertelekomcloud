package domain

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateCertOpts struct {
	GatewayID string `json:"-"`
	GroupID   string `json:"-"`
	DomainID  string `json:"-"`
	// Certificate content.
	Content string `json:"cert_content" required:"true"`
	// Certificate name. It can contain 4 to 50 characters, starting with a letter.
	// Only letters, digits, and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Private key.
	PrivateKey string `json:"private_key" required:"true"`
}

func AssignCertificate(client *golangsdk.ServiceClient, opts CreateCertOpts) (*DomainCertResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificate
	raw, err := client.Post(client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID, "domains", opts.DomainID, "certificate"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res DomainCertResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type DomainCertResp struct {
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
	// Certificate name.
	SslName string `json:"ssl_name"`
	// Certificate ID.
	SslId string `json:"ssl_id"`
}
