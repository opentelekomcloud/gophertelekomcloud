package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Update allows clusters to update description.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Clusters, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("clusters", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Clusters
	err = extract.Into(raw.Body, &res)
	return &res, err
}
