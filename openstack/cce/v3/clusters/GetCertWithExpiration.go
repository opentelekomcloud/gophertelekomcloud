package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ExpirationOpts struct {
	Duration int `json:"duration" required:"true"`
}

// GetCertWithExpiration retrieves a particular cluster certificate based on its unique ID.
func GetCertWithExpiration(client *golangsdk.ServiceClient, clusterId string, opts ExpirationOpts) (*Certificate, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /api/v3/projects/{project_id}/clusters/{cluster_id}/clustercert
	raw, err := client.Post(client.ServiceURL("clusters", clusterId, "clustercert"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res Certificate
	return &res, extract.Into(raw.Body, &res)
}
