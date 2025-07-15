package cert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// AttachDomainInfo represents the information for a domain to attach
type AttachDomainInfo struct {
	// Domain name
	Domain string `json:"domain" required:"true"`
	// Gateway IDs
	InstanceIDs []string `json:"instance_ids,omitempty"`
	// Whether to enable client certificate verification
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`
}

// AttachDomainOpts contains the options for binding a certificate to domain names
type AttachDomainOpts struct {
	CertificateID string `json:"-"`
	// Domain names the certificate is bound to
	Domains []AttachDomainInfo `json:"domains" required:"true"`
}

// BindCertToDomain binds an SSL certificate to domain names
func BindCertToDomain(client *golangsdk.ServiceClient, opts AttachDomainOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("apigw", "certificates", opts.CertificateID, "domains", "attach"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
