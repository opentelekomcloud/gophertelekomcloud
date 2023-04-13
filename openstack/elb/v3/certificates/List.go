package certificates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToCertificateListQuery() (string, error)
}

type ListOpts struct {
	ID          []string `q:"id"`
	Name        []string `q:"name"`
	Description []string `q:"description"`
	Type        []string `q:"type"`
	Domain      []string `q:"domain"`

	Limit       int    `q:"limit"`
	Marker      string `q:"marker"`
	PageReverse bool   `q:"page_reverse"`
}

// ToCertificateListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToCertificateListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("certificates")
	if opts != nil {
		query, err := opts.ToCertificateListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return CertificatePage{pagination.SinglePageBase(r)}
	})
}
