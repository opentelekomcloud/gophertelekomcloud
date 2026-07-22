package providers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Provider string `q:"provider"`
	Locale   string `q:"locale"`
	Limit    int    `q:"limit"`
	Offset   int    `q:"offset"`
}

// List retrieves the services (providers) that support enterprise projects.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Provider, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("enterprise-projects", "providers").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return res.Providers, err
}

type ListResponse struct {
	Providers  []Provider `json:"providers"`
	TotalCount int        `json:"total_count"`
}

type Provider struct {
	Provider            string          `json:"provider"`
	ProviderI18nDisplay string          `json:"provider_i18n_display_name"`
	ResourceTypes       []ResourceTypes `json:"resource_types"`
}

type ResourceTypes struct {
	ResourceType                string   `json:"resource_type"`
	ResourceTypeI18nDisplayName string   `json:"resource_type_i18n_display_name"`
	Regions                     []string `json:"regions"`
	Global                      bool     `json:"global"`
}
