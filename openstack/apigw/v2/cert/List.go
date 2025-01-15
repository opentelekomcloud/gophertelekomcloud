package cert

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	InstanceId string `q:"instance_id"`
	// Offset from which the query starts
	Offset *int64 `q:"offset"`
	// Number of items displayed on each page
	Limit *int `q:"limit"`
	// Certificate name
	Name string `q:"name"`
	// Certificate domain name
	CommonName string `q:"common_name"`
	// Certificate signature algorithm
	SignatureAlgorithm string `q:"signature_algorithm"`
	// Certificate scope (instance or global)
	Type string `q:"type"`
}

// List retrieves a list of SSL certificates
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]CertBase, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "certificates").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return CertificatePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractCertificates(pages)
}

// ExtractCertificates extracts certificates from the response
func ExtractCertificates(r pagination.NewPage) ([]CertBase, error) {
	var s struct {
		Size  int        `json:"size"`
		Total int64      `json:"total"`
		Certs []CertBase `json:"certs"`
	}
	err := extract.Into(bytes.NewReader((r.(CertificatePage)).Body), &s)
	return s.Certs, err
}

// CertBase represents the basic content of an SSL certificate
type CertBase struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	Type               string   `json:"type"`
	InstanceID         string   `json:"instance_id"`
	ProjectID          string   `json:"project_id"`
	CommonName         string   `json:"common_name"`
	San                []string `json:"san"`
	NotAfter           string   `json:"not_after"`
	SignatureAlgorithm string   `json:"signature_algorithm"`
	CreateTime         string   `json:"create_time"`
	UpdateTime         string   `json:"update_time"`
	HasTrustedRootCA   bool     `json:"is_has_trusted_root_ca"`
}

// CertificatePage represents a single page of certificates
type CertificatePage struct {
	pagination.NewSinglePageBase
}
