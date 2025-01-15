package cert

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the options for creating a new SSL certificate
type CreateOpts struct {
	// Certificate name. It can contain 4 to 50 characters, starting with a letter.
	// Only letters, digits, and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Certificate content
	CertContent string `json:"cert_content" required:"true"`
	// Certificate private key
	PrivateKey string `json:"private_key" required:"true"`
	// Certificate scope (instance or global)
	Type string `json:"type,omitempty"`
	// Gateway ID. Required if type is set to instance
	InstanceID string `json:"instance_id,omitempty"`
	// Trusted root certificate (CA)
	TrustedRootCA string `json:"trusted_root_ca,omitempty"`
}

// Create creates a new SSL certificate
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CertificateResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("apigw", "certificates"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CertificateResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// CertificateResp represents the response from certificate creation
type CertificateResp struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Type               string   `json:"type"`
	InstanceID         string   `json:"instance_id"`
	ProjectID          string   `json:"project_id"`
	CommonName         string   `json:"common_name"`
	San                []string `json:"san"`
	NotAfter           string   `json:"not_after"`
	NotBefore          string   `json:"not_before"`
	SignatureAlgorithm string   `json:"signature_algorithm"`
	CreateTime         string   `json:"create_time"`
	UpdateTime         string   `json:"update_time"`
	HasTrustedRootCA   bool     `json:"is_has_trusted_root_ca"`
	Version            int      `json:"version"`
	Organization       []string `json:"organization"`
	OrganizationalUnit []string `json:"organizational_unit"`
	Locality           []string `json:"locality"`
	State              []string `json:"state"`
	Country            []string `json:"country"`
	SerialNumber       string   `json:"serial_number"`
	Issuer             []string `json:"issuer"`
}
