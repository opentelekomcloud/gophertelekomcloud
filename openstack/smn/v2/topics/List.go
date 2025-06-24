package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Offset string `q:"offset,omitempty"`
	Limit  int    `q:"limit,omitempty"`
}

// List all the topics
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Topic, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("topics").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	// GET /v2/{project_id}/notifications/topics
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []Topic
	err = extract.IntoSlicePtr(raw.Body, &res, "topics")
	return res, err
}
