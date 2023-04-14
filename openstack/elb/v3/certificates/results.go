package certificates

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CertificatePage is the page returned by a pager when traversing over a
// collection of certificates.
type CertificatePage struct {
	pagination.SinglePageBase
}

// ExtractCertificates accepts a Page struct, specifically a CertificatePage struct,
// and extracts the elements into a slice of Certificate structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractCertificates(r pagination.Page) ([]Certificate, error) {
	var s []Certificate
	err := (r.(CertificatePage)).ExtractIntoSlicePtr(&s, "certificates")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Certificate, error) {
	s := new(Certificate)
	err := r.ExtractIntoStructPtr(s, "certificate")
	if err != nil {
		return nil, err
	}
	return s, nil
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
