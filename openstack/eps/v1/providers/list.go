package providers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Provider struct {
	Provider            string `json:"provider"`
	ProviderI18nDisplay string `json:"provider_i18n_display_name"`
}

type ListOpts struct {
	Locale string `q:"locale"`
	Limit  int    `q:"limit"`
	Offset int    `q:"offset"`
}

func (opts ListOpts) ToQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List retrieves the services (providers) that support enterprise projects.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Provider, error) {
	url := client.ServiceURL("enterprise-projects", "providers")
	query, err := opts.ToQuery()
	if err != nil {
		return nil, err
	}
	url += query

	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Provider
	err = extract.IntoSlicePtr(raw.Body, &res, "providers")
	return res, err
}
