package cert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains the options for updating an SSL certificate
type UpdateOpts struct {
	// Certificate name. It can contain 4 to 50 characters, starting with a letter.
	// Only letters, digits, and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Certificate content
	CertContent string `json:"cert_content" required:"true"`
	// Certificate private key
	PrivateKey string `json:"private_key" required:"true"`
	// Certificate scope
	Type string `json:"type,omitempty"`
	// Gateway ID. Required if type is set to instance
	InstanceID string `json:"instance_id,omitempty"`
	// Trusted root certificate (CA)
	TrustedRootCA string `json:"trusted_root_ca,omitempty"`
}

// Update modifies an existing SSL certificate
func Update(client *golangsdk.ServiceClient, certificateID string, opts UpdateOpts) (*CertificateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("apigw", "certificates", certificateID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CertificateResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}
