package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This API is used to synchronize configuration information of all data nodes that are associated with a DDM instance.
func SyncDataNodesInfo(client *golangsdk.ServiceClient, instanceId string) (*SyncDataNodesInfoResponse, error) {

	// POST /v1/{project_id}/instances/{instance_id}/rds/sync
	raw, err := client.Post(client.ServiceURL("instances", instanceId, "rds", "sync"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res SyncDataNodesInfoResponse
	return &res, extract.Into(raw.Body, &res)
}

type SyncDataNodesInfoResponse struct {
	// DDM instance ID
	InstanceId string `json:"instanceId"`
	// Task ID
	JobId string `json:"jobId"`
}
