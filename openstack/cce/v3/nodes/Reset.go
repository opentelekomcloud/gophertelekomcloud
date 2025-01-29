package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ResetOpts struct {
	ClusterID string `json:"-"`
	// API type, fixed value List
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// List of nodes to reset
	NodeList []ResetNode `json:"nodeList" required:"true"`
}

type ResetNode struct {
	NodeID string            `json:"nodeID" required:"true"`
	Spec   ReinstallNodeSpec `json:"spec" required:"true"`
}

func Reset(client *golangsdk.ServiceClient, opts ResetOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes/reset
	raw, err := client.Post(client.ServiceURL("clusters", opts.ClusterID, "nodes", "reset"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res JobResult
	err = extract.Into(raw.Body, &res)
	return res.JobID, err
}
