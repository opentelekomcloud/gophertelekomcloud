package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ModifyDbReadPolicyOpts struct {
	// Read weights of the primary DB instance and its read replicas (Mandatory)
	// Key: DB instance ID (string), Value: read weight parameter (integer)
	ReadWeight map[string]int `json:"read_weight" required:"true"`
}

// ModifyDbReadPolicy is used to modify the read policy of the DB instance associated with a DDM instance.
func ModifyDbReadPolicy(client *golangsdk.ServiceClient, instanceId string, opts ModifyDbReadPolicyOpts) (*ModifyDbReadPolicyResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/instances/{instance_id}/action/read-write-strategy
	raw, err := client.Put(client.ServiceURL("instances", instanceId, "action", "read-write-strategy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ModifyDbReadPolicyResponse
	return &res, extract.Into(raw.Body, &res)
}

type ModifyDbReadPolicyResponse struct {
	// Specifies whether the operation is successful
	Success bool `json:"success"`
	// DDM instance ID
	InstanceId string `json:"instance_id"`
}
