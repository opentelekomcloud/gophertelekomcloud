package providers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Provider struct {
	ID          string            `json:"id"`
	Description string            `json:"description"`
	Enabled     bool              `json:"enabled"`
	RemoteIDs   []string          `json:"remote_ids"`
	Links       map[string]string `json:"links"`
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Provider, error) {
	s := new(Provider)
	return s, r.ExtractIntoStructPtr(s, "identity_provider")
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type ProviderPage struct {
	pagination.LinkedPageBase
}

func (p ProviderPage) IsEmpty() (bool, error) {
	providers, err := ExtractProviders(p)
	if err != nil {
		return false, err
	}
	return len(providers) == 0, err
}

func ExtractProviders(r pagination.Page) ([]Provider, error) {
	var providers []Provider

	err := extract.IntoSlicePtr((r.(ProviderPage)).Body, &providers, "identity_providers")
	return providers, err
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}
