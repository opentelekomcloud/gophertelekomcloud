package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type KillSessionsOpts struct {
	// Specifies the node ID. The following nodes can be queried: mongos nodes in the cluster, and all nodes in the replica set and single node instances.
	NodeId string `json:"-"`
	// Specifies the IDs of sessions to be terminated.
	Sessions []string `json:"sessions" required:"true"`
}

func KillSessions(client *golangsdk.ServiceClient, opts KillSessionsOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST https://{Endpoint}/v3/{project_id}/nodes/{node_id}/session
	_, err = client.Post(client.ServiceURL("nodes", opts.NodeId, "session"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return err
	}

	return nil
}
