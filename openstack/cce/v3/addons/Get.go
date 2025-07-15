package addons

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular addon based on its unique ID.
func Get(client *golangsdk.ServiceClient, addonId string, clusterId string) (*Addon, error) {
	// GET /api/v3/addons/{id}?cluster_id={cluster_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("addons", addonId).WithQueryParams(&ClusterIdQueryParam{ClusterId: clusterId}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(CCEServiceURL(client, clusterId, url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Addon
	return &res, extract.Into(raw.Body, &res)
}
