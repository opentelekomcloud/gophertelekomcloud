package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOptsBuilder interface {
	ToConfigurationListQuery() (string, error)
}

type ListOpts struct {
	Name        string `q:"scaling_configuration_name"`
	ImageID     string `q:"image_id"`
	StartNumber int    `q:"start_number"`
	Limit       int    `q:"limit"`
}

func (opts ListOpts) ToConfigurationListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), err
}

func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := client.ServiceURL("scaling_configuration")
	if opts != nil {
		query, err := opts.ToConfigurationListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ConfigurationPage{pagination.SinglePageBase(r)}
	})
}

type ConfigurationPage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ListResult contains no Volumes.
func (r ConfigurationPage) IsEmpty() (bool, error) {
	configs, err := r.Extract()
	return len(configs) == 0, err
}

func (r ConfigurationPage) Extract() ([]Configuration, error) {
	var cs []Configuration
	err := r.Result.ExtractIntoSlicePtr(&cs, "scaling_groups")
	return cs, err
}

func ExtractConfigurations(r pagination.Page) ([]Configuration, error) {
	var s struct {
		Configurations []Configuration `json:"scaling_configurations"`
	}
	err := (r.(ConfigurationPage)).ExtractInto(&s)
	return s.Configurations, err
}
