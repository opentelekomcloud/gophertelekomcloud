package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateIpOpts struct {
	Action    string `json:"action" required:"true"`
	Spec      IpSpec `json:"spec,omitempty"`
	ElasticIp string `json:"elasticIp"`
}

type IpSpec struct {
	ID string `json:"id" required:"true"`
}

// Update the access information of a specified cluster.
func UpdateMasterIp(client *golangsdk.ServiceClient, clusterId string, opts UpdateIpOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// PUT /api/v3/projects/{project_id}/clusters/{cluster_id}/mastereip
	_, err = client.Put(client.ServiceURL("clusters", clusterId, "mastereip"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return err
	}

	return nil
}
