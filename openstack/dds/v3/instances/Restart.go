package instances

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestartOpts struct {
	// Specifies the instance ID, which can be obtained by calling the API for querying instances.
	InstanceId string `json:"-"`
	// Specifies the type of the object to restart.
	//
	// This parameter is mandatory when you restart one or more nodes of a cluster instance.
	// Set the value to mongos if mongos nodes are restarted.
	// Set the value to shard if shard nodes are restarted.
	// Set the value to config if config nodes are restarted.
	// This parameter is not transferred when the DB instance is restarted.
	TargetType string `json:"target_type,omitempty"`
	// Specifies the ID of the object to be restarted, which can be obtained by calling the API for querying instances. If you do not have an instance, you can call the API used for creating an instance.
	//
	// In a cluster instance, the value is the ID of the node to restart.
	// When you restart the entire DB instance, the value is the DB instance ID.
	TargetId string `json:"target_id" required:"true"`
}

func Restart(client *golangsdk.ServiceClient, opts RestartOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/restart
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "restart"), b, nil, nil)
	return extractJob(err, raw)
}
