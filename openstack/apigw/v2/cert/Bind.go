package cert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// BindOpts contains the options for binding SSL certificates to a domain
type BindOpts struct {
	// Gateway ID
	InstanceID string `json:"-"`
	// API group ID
	GroupID string `json:"-"`
	// Domain ID
	DomainID string `json:"-"`
	// Certificate IDs to attach
	CertificateIDs []string `json:"certificate_ids" required:"true"`
	// Whether to enable client certificate verification
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`
}

// Bind SSL certificates to a domain
func Bind(client *golangsdk.ServiceClient, opts BindOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// Build URL: /v2/{project_id}/apigw/instances/{instance_id}/api-groups/{group_id}/domains/{domain_id}/certificates/attach
	_, err = client.Post(client.ServiceURL("apigw", "instances", opts.InstanceID,
		"api-groups", opts.GroupID, "domains", opts.DomainID, "certificates", "attach"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
