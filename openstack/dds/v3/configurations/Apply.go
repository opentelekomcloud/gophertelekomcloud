package configurations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dds/v3/instances"
)

type ApplyOpts struct {
	// Instance IDs, group IDs, or node IDs. You can call the API used for querying instances and details to obtain the value.
	// If you do not have an instance, you can call the API used for creating an instance.
	// If the DB instance type is cluster and the shard or config parameter template is to be changed, the value is the group ID.
	// If the parameter template of the mongos node is changed, the value is the node ID.
	// If the DB instance to be changed is a replica set instance or a single node instance, the value is the instance ID.
	EntityIDs []string `json:"entity_ids" required:"true"`
}

func Apply(client *golangsdk.ServiceClient, configId string, opts ApplyOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT https://{Endpoint}/v3/{project_id}/configurations/{config_id}/apply
	raw, err := client.Put(client.ServiceURL("configurations", configId, "apply"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	return instances.ExtractJob(err, raw)
}
