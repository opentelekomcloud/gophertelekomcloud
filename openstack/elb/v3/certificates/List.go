package certificates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies a certificate ID.
	//
	// Multiple IDs can be queried in the format of id=xxx&id=xxx.
	ID []string `q:"id"`
	// Specifies the certificate name.
	//
	// Multiple names can be queried in the format of name=xxx&name=xxx.
	Name []string `q:"name"`
	// Provides supplementary information about the certificate.
	//
	// Multiple descriptions can be queried in the format of description=xxx&description=xxx.
	Description []string `q:"description"`
	// Specifies the certificate type.
	//
	// The value can be server or client. server indicates server certificates, and client indicates CA certificates.
	//
	// Multiple types can be queried in the format of type=xxx&type=xxx.
	Type []string `q:"type"`
	// Specifies the domain names used by the server certificate. This parameter is available only when type is set to server.
	//
	// Multiple domain names can be queried in the format of domain=xxx&domain=xxx.
	Domain []string `q:"domain"`
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	Limit int `q:"limit"`
	// Specifies the ID of the last record on the previous page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If this parameter is not specified, the first page will be queried.
	//
	// This parameter cannot be left blank or set to an invalid ID.
	Marker string `q:"marker"`
	// Specifies the page direction. The value can be true or false, and the default value is false. The last page in the list requested with page_reverse set to false will not contain the "next" link, and the last page in the list requested with page_reverse set to true will not contain the "previous" link. This parameter must be used together with limit.
	PageReverse bool `q:"page_reverse"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}
	// GET /v3/{project_id}/elb/certificates
	return pagination.NewPager(client, client.ServiceURL("certificates")+query.String(), func(r pagination.PageResult) pagination.Page {
		return CertificatePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CertificatePage is the page returned by a pager when traversing over a
// collection of certificates.
type CertificatePage struct {
	pagination.LinkedPageBase
}

// ExtractCertificates accepts a Page struct, specifically a CertificatePage struct,
// and extracts the elements into a slice of Certificate structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractCertificates(r pagination.Page) ([]Certificate, error) {
	var res []Certificate
	err := extract.IntoSlicePtr(r.(CertificatePage).BodyReader(), &res, "certificates")
	if err != nil {
		return nil, err
	}
	return res, nil
}
