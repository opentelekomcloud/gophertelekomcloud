package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ExpandInstanceNodeOpts struct {
	InstanceId string
	// Number of new nodes
	Num int32 `json:"num"`
}

func ExpandInstanceNode(client *golangsdk.ServiceClient, opts ExpandInstanceNodeOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/enlarge-node
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "enlarge-node"), b, nil, nil)
	return extraJob(err, raw)
}
