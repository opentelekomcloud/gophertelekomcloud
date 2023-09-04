package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Certificates, error) {
	query, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/waf/certificate
	url := client.ServiceURL("waf", "certificate") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Certificates
	err = extract.IntoSlicePtr(raw.Body, &res, "items")
	return res, err
}

type ListOpts struct {
	// Number of records on each page.
	// The maximum value is 100. Default value: 10
	PageSize string `q:"pageSize,omitempty"`
	// Current page number
	Page string `q:"page,omitempty"`
	// Domain name
	Name string `q:"name,omitempty"`
	// Whether to obtain the domain name associated with the certificate.
	// The value can be true or false.
	// true: When a certificate is queried, the domain name associated with the certificate is also queried.
	// The returned certificate information contains the associated domain name.
	// false: When a certificate is queried, the domain name associated with the certificate is not queried.
	// The returned certificate information does not contain the associated domain name.
	// Default value: false
	// Default: false
	Host *bool `q:"host,omitempty"`
	// Certificate status. The value can be:
	// 0: The certificate is valid.
	// 1: The certificate has expired.2: The certificate will expire within one month.
	ExpirationStatus int `q:"exp_status,omitempty"`
}

type Certificates struct {
	// Certificate ID.
	ID string `json:"id"`
	// Certificate name.
	Name string `json:"name"`
	// Timestamp when the certificate is uploaded.
	CreatedAt int64 `json:"timestamp"`
}
