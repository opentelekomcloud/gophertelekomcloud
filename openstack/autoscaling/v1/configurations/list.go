package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Name        string `q:"scaling_configuration_name"`
	ImageID     string `q:"image_id"`
	StartNumber int    `q:"start_number"`
	Limit       int    `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Configuration, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("scaling_configuration")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Configuration
	err = extract.IntoSlicePtr(raw.Body, &res, "scaling_configurations")
	return res, err
}
