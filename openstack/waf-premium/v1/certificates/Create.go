package certificates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Certificate name. The value can contain a maximum of 64 characters.
	// Only digits, letters, hyphens (-), underscores (_), and periods (.) are allowed.
	Name string `json:"name" required:"true"`
	// Certificate file. Only certificates and private key files in PEM format are supported,
	// and the newline characters in the file must be replaced with \n.
	Content string `json:"content" required:"true"`
	// Certificate private key.
	// Only certificates and private key files in PEM format are supported,
	// and the newline characters in the files must be replaced with \n.
	Key string `json:"key" required:"true"`
}

// Create will create a new Waf Certificate on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CertificateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/certificate
	raw, err := client.Post(client.ServiceURL("waf", "certificate"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res CertificateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CertificateResponse struct {
	// Certificate ID.
	ID string `json:"id"`
	// Certificate name.
	Name string `json:"name"`
	// Timestamp when the certificate expires.
	ExpireAt int64 `json:"expire_time"`
	// Timestamp when the certificate is uploaded.
	CreatedAt int64 `json:"timestamp"`
}
