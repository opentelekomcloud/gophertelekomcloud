package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func DeleteReplica(client *golangsdk.ServiceClient, instanceID string, nodeId string) (*string, error) {
	raw, err := client.Delete(client.ServiceURL("instances", instanceID, "nodes", nodeId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res jobResponse
	return &res.JobID, extract.Into(raw.Body, &res)
}
