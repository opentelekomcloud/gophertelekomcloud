package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetCert retrieves a particular cluster certificate based on its unique ID. (Depreciated)
func GetCert(client *golangsdk.ServiceClient, clusterId string) (*Certificate, error) {
	// GET /api/v3/projects/{project_id}/clusters/{cluster_id}/clustercert
	raw, err := client.Get(client.ServiceURL("clusters", clusterId, "clustercert"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res Certificate
	return &res, extract.Into(raw.Body, &res)
}
