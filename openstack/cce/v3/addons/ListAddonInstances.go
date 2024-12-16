package addons

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListAddonInstances(client *golangsdk.ServiceClient, clusterId string) (*AddonInstanceList, error) {
	// GET /api/v3/addons?cluster_id={cluster_id}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("addons").WithQueryParams(&ClusterIdQueryParam{ClusterId: clusterId}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(CCEServiceURL(client, clusterId, url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res AddonInstanceList
	return &res, extract.Into(raw.Body, &res)
}
