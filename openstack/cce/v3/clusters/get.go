package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular cluster based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Clusters, error) {
	raw, err := client.Get(client.ServiceURL("clusters", id), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})
	if err != nil {
		return nil, err
	}

	var res Clusters
	err = extract.Into(raw.Body, &res)
	return &res, err
}
