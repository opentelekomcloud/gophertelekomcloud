package nodepools

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

var RequestOpts = map[string]string{"Content-Type": "application/json"}

// Get retrieves a particular node pool based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterid, nodepoolid string) (*NodePool, error) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterid, "nodepools", nodepoolid), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})
	if err != nil {
		return nil, err
	}

	var res NodePool
	err = extract.Into(raw.Body, &res)
	return &res, err
}
