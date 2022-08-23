package groups_hcs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Name            string `q:"scaling_group_name"`
	ConfigurationID string `q:"scaling_configuration_id"`
	Status          string `q:"scaling_group_status"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Group, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("scaling_group")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Group
	err = extract.IntoSlicePtr(raw.Body, &res, "scaling_groups")
	return res, err
}
