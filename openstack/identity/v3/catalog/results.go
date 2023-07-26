package catalog

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/tokens"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// CatalogPage is a single page of Service results.
type CatalogPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the CatalogPage contains no results.
func (p CatalogPage) IsEmpty() (bool, error) {
	services, err := ExtractServiceCatalog(p)
	return len(services) == 0, err
}

// ExtractServiceCatalog extracts a slice of Catalog from a Collection acquired
// from List.
func ExtractServiceCatalog(r pagination.Page) ([]tokens.CatalogEntry, error) {
	var s struct {
		Entries []tokens.CatalogEntry `json:"catalog"`
	}

	err := extract.Into(bytes.NewReader((r.(CatalogPage)).Body), &s)
	return s.Entries, err
}
