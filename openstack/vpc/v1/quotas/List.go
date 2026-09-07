package quotas

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Type string `q:"type,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Quota, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(client.ProjectID, "quotas").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Quotas struct {
			Resources []Quota `json:"resources"`
		} `json:"quotas"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Quotas.Resources, err
}

type Quota struct {
	Type  string `json:"type"`
	Used  int    `json:"used"`
	Quota int    `json:"quota"`
	Min   int    `json:"min"`
}
