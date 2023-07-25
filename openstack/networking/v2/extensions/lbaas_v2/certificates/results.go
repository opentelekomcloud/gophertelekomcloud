package certificates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Certificate struct {
	ID           string `json:"id"`
	TenantID     string `json:"tenant_id"`
	AdminStateUp bool   `json:"admin_state_up"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Domain       string `json:"domain"`
	PrivateKey   string `json:"private_key"`
	Certificate  string `json:"certificate"`
	ExpireTime   string `json:"expire_time"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

// CertificatePage is the page returned by a pager when traversing over a
// collection of certificates.
type CertificatePage struct {
	pagination.SinglePageBase
}

// ExtractCertificates accepts a Page struct, specifically a CertificatePage struct,
// and extracts the elements into a slice of Certificate structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractCertificates(r pagination.Page) ([]Certificate, error) {
	var s struct {
		Certificates []Certificate `json:"certificates"`
	}

	err := extract.Into((r.(CertificatePage)).Body, &s)
	return s.Certificates, err
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Certificate, error) {
	s := &Certificate{}
	return s, r.ExtractInto(s)
}

type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation.
type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation.
type DeleteResult struct {
	golangsdk.ErrResult
}
