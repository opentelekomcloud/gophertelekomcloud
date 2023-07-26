package certificates

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Certificate struct {
	// Id of the certificate
	Id string `json:"id"`
	// Name of the certificate
	Name string `json:"name"`
	// ExpireTime - unix timestamp of certificate expiry
	ExpireTime int `json:"expireTime"`

	Timestamp int `json:"timestamp"`
}

type commonResult struct {
	golangsdk.Result
}

// Extract is a function that accepts a result and extracts a certificate.
func (r commonResult) Extract() (*Certificate, error) {
	var response Certificate
	err := r.ExtractInto(&response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Certificate.
type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of a update operation. Call its Extract
// method to interpret it as a Certificate.
type UpdateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Certificate.
type GetResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

type CertificatePage struct {
	pagination.OffsetPageBase
}

// IsEmpty returns true if this Page has no items in it.
func (p CertificatePage) IsEmpty() (bool, error) {
	body, err := p.GetBodyAsMap()
	if err != nil {
		return false, err
	}

	items, ok := body["items"].([]any)
	if !ok {
		return false, fmt.Errorf("map `items` are not a slice: %+v", body)
	}

	return len(items) == 0, nil
}

func (p CertificatePage) NextPageURL() (string, error) {
	currentURL := p.URL
	q := currentURL.Query()
	// The default value is 10. If limit is -1, one page with 65535 records is displayed.
	switch q.Get("limit") {
	case "-1":
		return "", nil // in this case is a SinglePageBase
	case "":
		q.Set("limit", "10")
		p.Limit = 10
	}
	// Its value ranges from 0 to 65535. The default value is 0.
	if q.Get("offset") == "" {
		p.Offset = 0
	}
	q.Set("offset", strconv.Itoa(p.LastElement()))
	currentURL.RawQuery = q.Encode()
	return currentURL.String(), nil
}

func ExtractCertificates(p pagination.Page) ([]Certificate, error) {
	var certs []Certificate
	err := extract.IntoSlicePtr(bytes.NewReader(p.(CertificatePage).Body), &certs, "items")
	return certs, err
}
