package certificates

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	ID          []string `q:"id,omitempty"`
	Name        []string `q:"name,omitempty"`
	Description []string `q:"description,omitempty"`
	Type        []string `q:"type,omitempty"`
	Domain      []string `q:"domain,omitempty"`
	Limit       int      `q:"limit,omitempty"`
	Marker      string   `q:"marker,omitempty"`
	PageReverse bool     `q:"page_reverse,omitempty"`

	// AdminStateUp is documented but marked unsupported by the reviewed OTC API.
	AdminStateUp *bool `q:"admin_state_up,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Certificate, error) {
	if opts.Limit == 0 {
		opts.Limit = 2000
	}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("certificates").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return CertificatePage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractCertificates(pages)
}

type CertificatePage struct {
	pagination.NewPageResult
}

func (p CertificatePage) NewNextPageURL() (string, error) {
	var res struct {
		PageInfo PageInfo `json:"page_info"`
	}
	if err := extract.Into(bytes.NewReader(p.Body), &res); err != nil {
		return "", err
	}
	marker := res.PageInfo.NextMarker
	if p.URL.Query().Get("page_reverse") == "true" {
		marker = res.PageInfo.PreviousMarker
	}
	if marker == "" {
		return "", nil
	}
	next := p.URL
	query := next.Query()
	query.Set("marker", marker)
	next.RawQuery = query.Encode()
	return next.String(), nil
}

func (p CertificatePage) NewIsEmpty() (bool, error) {
	certificates, err := ExtractCertificates(p)
	return len(certificates) == 0, err
}

func ExtractCertificates(page pagination.NewPage) ([]Certificate, error) {
	var res struct {
		Certificates []Certificate `json:"certificates"`
	}
	err := extract.Into(bytes.NewReader(page.(CertificatePage).Body), &res)
	return res.Certificates, err
}

type PageInfo struct {
	PreviousMarker string `json:"previous_marker"`
	NextMarker     string `json:"next_marker"`
	CurrentCount   int    `json:"current_count"`
}
